// Code generated by goctl. DO NOT EDIT.
package types

type UserBackend struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Status   bool   `json:"status"`
}

type AuthRequest struct {
}

type AuthResponse struct {
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
}

type DeleteResponse struct {
}

type IdRequest struct {
	Id string `path:"id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ModifyPasswordRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type ModifyRefreshSecretRequest struct {
	Id   string `json:"id"`
	Hour int64  `json:"hour"` // hour可以指定token有效时间
}

type ModifyStatusRequest struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}

type ModifyUserNameAndPasswordRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ModifyUserNameRequest struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type SelectPage struct {
	PageNum  int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}

type SelectPageResponse struct {
	Data  []UserBackend `json:"data"`
	Total int64         `json:"total"`
}
