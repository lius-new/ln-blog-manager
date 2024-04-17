// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var prefixTagCacheKey = "cache:tag:"

type tagModel interface {
	Insert(ctx context.Context, data *Tag) error
	FindOne(ctx context.Context, id string) (*Tag, error)
	Update(ctx context.Context, data *Tag) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultTagModel struct {
	conn *monc.Model
}

func newDefaultTagModel(conn *monc.Model) *defaultTagModel {
	return &defaultTagModel{conn: conn}
}

func (m *defaultTagModel) Insert(ctx context.Context, data *Tag) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixTagCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultTagModel) FindOne(ctx context.Context, id string) (*Tag, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Tag
	key := prefixTagCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTagModel) Update(ctx context.Context, data *Tag) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixTagCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultTagModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}
	key := prefixTagCacheKey + id
	res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return res, err
}
