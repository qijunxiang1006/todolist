package model

import (
	uuid "github.com/iris-contrib/go.uuid"
)

type DBData struct {
	UUID uuid.UUID     `db:"uuid" json:"uuid"`
	List []*DBListItem `db:"list" json:"list"`
}

type DBListItem struct {
	ID     int    `db:"id" json:"id"`
	Value  string `db:"value" json:"value"`
	Status bool   `db:"status" json:"status"`
}

type UUIDJson struct {
	UUID []byte `json:"uuid"`
}

type response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

func NewSuccessResponse(data interface{}) interface{} {
	return response{
		Data:    data,
		Code:    0,
		Message: "succeed",
	}
}
func NewFailResponse(err error) interface{} {
	return response{
		Data:    nil,
		Code:    1,
		Message: err.Error(),
	}
}
