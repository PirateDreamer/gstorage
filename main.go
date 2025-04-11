package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	h := server.Default()

	h.PUT("/*key", func(c context.Context, ctx *app.RequestContext) {
		key := ctx.Param("key")
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(400, utils.H{"error": "文件获取失败", "details": err.Error()})
			return
		}
		if err = ctx.SaveUploadedFile(file, key); err != nil {
			ctx.JSON(400, utils.H{"error": "文件保存失败", "details": err.Error()})
			return
		}
	})

	h.GET("/*key", func(c context.Context, ctx *app.RequestContext) {
		key := ctx.Param("key")
		ctx.File(key)
	})

	h.Spin()
}
