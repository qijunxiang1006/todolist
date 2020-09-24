package model

type DBData struct {
	ID int `db:"id"`
	CreateTime string `db:"createtime"`
	UpdateTime string `db:"updatetime"`
}