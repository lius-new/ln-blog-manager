package models

import (
	"context"
	"errors"
	"time"

	"github.com/lius-new/liusnew-blog-backend-server/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	Id      string
	Title   string
	Content string
	Tags    []string
	Covers  []string
	Status  bool
	Time    int64
}

func (a *Article) ToBson() (d bson.D) {
	d = append(d, bson.E{Key: "title", Value: a.Title})
	d = append(d, bson.E{Key: "content", Value: a.Content})
	d = append(d, bson.E{Key: "tags", Value: a.Tags})
	d = append(d, bson.E{Key: "covers", Value: a.Covers})
	d = append(d, bson.E{Key: "status", Value: a.Status})
	d = append(d, bson.E{Key: "time", Value: a.Time})
	return
}

func CreateArticles(title, content string, tags, covers []string) (*Article, error) {
	save := func(a *Article) (string, error) {
		client := Pool.GetClient()
		Pool.ReleaseClient(client)

		coll := client.Database("liusnew-blog").Collection("articles")

		ctx := context.Background()

		if count, err := coll.CountDocuments(ctx, a.ToBson()); err != nil {
			panic(err)
		} else if count > 0 {
			return "", errors.New("article exist")
		}

		insertRes, err := coll.InsertOne(ctx, a.ToBson())

		if err != nil {
			panic(err)
		}

		return insertRes.InsertedID.(primitive.ObjectID).Hex(), nil
	}

	a := Article{
		"", title, content, tags, covers, true, time.Now().UnixNano(),
	}
	id, err := save(&a)
	a.Id = id

	return &a, err
}
func ModifyArticles(id, title, content string, tags, covers []string, status bool) (*Article, error) {
	client := Pool.GetClient()
	Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	modify := func(a *Article) error {
		objectId, _ := primitive.ObjectIDFromHex(a.Id)

		if count, err := coll.CountDocuments(ctx, bson.D{{"_id", objectId}}); err != nil {
			panic(err)
		} else if count == 0 {
			return errors.New("article not found")
		}

		_, err := coll.UpdateOne(ctx, bson.D{{"_id", objectId}}, bson.D{{"$set", a.ToBson()}})

		if err != nil {
			panic(err)
		}

		return nil
	}

	a := Article{
		id, title, content, tags, covers, status, time.Now().UnixNano(),
	}

	err := modify(&a)

	return &a, err
}

func DeleteArticles(id string) {
	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	delete := func(id string) {
		objectId, _ := primitive.ObjectIDFromHex(id)
		deleteRes, err := coll.DeleteOne(ctx, bson.D{{"_id", objectId}})
		if err != nil {
			panic(err)
		}
		logger.Debug(deleteRes)
	}
	delete(id)
}

func ViewArticles(pageSize, pageNumber int64) ([]Article, int64) {

	client := Pool.GetClient()
	defer Pool.ReleaseClient(client)
	coll := client.Database("liusnew-blog").Collection("articles")
	ctx := context.Background()

	view := func() (articles []Article) {
		findOptions := options.Find()

		if pageNumber <= 0 {
			pageNumber = 1
		}

		findOptions.SetLimit(pageSize)
		findOptions.SetSkip(pageSize * (pageNumber - 1))
		findOptions.SetSort(bson.M{"time": 1})

		cur, err := coll.Find(ctx, bson.D{{}}, findOptions)
		if err != nil {
			panic(err)
		}

		cur.All(ctx, &articles)
		return
	}

	count := func() int64 {
		c, err := coll.CountDocuments(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}
		return int64(c/pageSize) + 1
	}

	return view(), count()
}
