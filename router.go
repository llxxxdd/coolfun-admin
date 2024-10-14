/*
 * Copyright 2023 CoolFunAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package main

import (
	"context"
	"coolfun-admin/biz/handler/middleware"
	"coolfun-admin/configs"
	"coolfun-admin/data"
	"coolfun-admin/data/s3"
	"io"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	// your code ...
	// login and refresh_token power by jwt Auth middleware
	r.POST("/api/login", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).LoginHandler)
	r.POST("/api/logout", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).LogoutHandler)
	r.POST("/api/refresh_token", middleware.GetJWTMiddleware(configs.Data(), data.Default(), data.CasbinEnforcer()).RefreshHandler)
	r.POST("/api/upload", uploadFileHandler)
}

func uploadFileHandler(c context.Context, ctx *app.RequestContext) {
	// 使用 Hertz 的 FormFile 方法从请求中获取文件
	formFile, err := ctx.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		ctx.JSON(400, map[string]string{
			"error": "Failed to retrieve file",
		})
		return
	}

	if formFile == nil {
		log.Println("No file received")
		ctx.JSON(400, map[string]string{
			"error": "No file received",
		})
		return
	}

	log.Printf("File received: %s", formFile.Filename)

	// 打开文件流
	srcFile, err := formFile.Open()
	if err != nil {
		ctx.JSON(500, map[string]string{
			"error": "Failed to open file",
		})
		return
	}
	defer srcFile.Close()

	// 将文件流转换为 io.Reader
	fileReader := io.Reader(srcFile)

	// 创建 FileInfo 对象
	fileInfo := &s3.FileInfo{
		Name:   formFile.Filename,
		Size:   formFile.Size,
		Path:   "",     // 如果需要指定路径，可以在这里设定
		Type:   "file", // 可以根据需要修改文件类型
		Reader: &fileReader,
	}

	// 使用 fileStore 来存储文件
	savedFile, err := data.FileStore.Adapter.UploadFile(c, fileInfo)
	if err != nil {
		println("上传文件异常" + err.Error())
		ctx.JSON(500, map[string]string{
			"error": "Failed to upload file",
		})
		return
	}

	// 返回保存的文件信息
	ctx.JSON(200, map[string]string{
		"filePath": savedFile.Path,
		"fileName": savedFile.Name,
	})
}
