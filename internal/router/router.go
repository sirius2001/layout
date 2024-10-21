package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterInner interface {
	Route(e *gin.Engine)
}

type router struct{}

func (r *router) success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "OK", "data": data})
	ctx.Abort()
}

func (r *router) faild(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "errpr": err.Error()})
	ctx.Abort()
}
