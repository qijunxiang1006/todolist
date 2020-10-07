package controller

import (
	"strconv"

	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	
	`todolist/model`
)

type Controller struct {
	Ser Service
}

func (c *Controller) Get(ctx iris.Context) {
	param := ctx.URLParam("uuid")
	uid, err := uuid.FromString(param)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	data, err := c.Ser.GetList(uid)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	data.UUID, _ = uuid.FromString(param)
	_, _ = ctx.JSON(model.NewSuccessResponse(data))

}
func (c *Controller) Post(ctx iris.Context) {
	param := ctx.URLParam("uuid")
	uid, err := uuid.FromString(param)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	body := new(model.DBListItem)
	err = ctx.ReadJSON(body)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	id, err := c.Ser.PostList(uid, body)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	_, _ = ctx.JSON(model.NewSuccessResponse(id))

}
func (c *Controller) Put(ctx iris.Context) {
	param := ctx.URLParam("uuid")
	uid, err := uuid.FromString(param)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	body := new(model.DBListItem)
	err = ctx.ReadJSON(body)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	err = c.Ser.PutList(uid, body)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	_, _ = ctx.JSON(model.NewSuccessResponse(nil))

}
func (c *Controller) Delete(ctx iris.Context) {
	param := ctx.URLParam("uuid")
	uid, err := uuid.FromString(param)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	param = ctx.URLParam("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	err = c.Ser.DeleteList(uid, id)
	if err != nil {
		_, _ = ctx.JSON(model.NewFailResponse(err))
		return
	}
	_, _ = ctx.JSON(model.NewSuccessResponse(nil))
}
