package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lfh1999/seeme/model/hello"
)

type Hello struct {
	Base
}

func init() {
	h := &Hello{}
	h.LoadPostMethod("/api/seeme/hello-seeme", h.HelloSeeme)
	h.LoadController(h)
}

func (h *Hello) HelloSeeme(ctx *gin.Context) {
	obj := &hello.Hello{}
	obj.GetHelloSeeme()
	h.Success(ctx, "success", obj)
}
