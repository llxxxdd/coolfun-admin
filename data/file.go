/*
 * Copyright 2023 CoolFunAdmin Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package data

import (
	"coolfun-admin/configs"
	"coolfun-admin/data/s3"
)

// FileStore offers file store interface (大写以导出变量)
var FileStore *FileAdapter

func initFileStore() {
	FileStore = NewFileAdapter(configs.Data())
}

// GetFileStore 返回一个默认的 file store 实例
func GetFileStore() *FileAdapter {
	if FileStore == nil {
		initFileStore()
	}
	return FileStore
}

type FileAdapter struct {
	Adapter s3.Adapter
}

func NewFileAdapter(c configs.Config) *FileAdapter {
	if c.S3.Type == "AliyunOSS" {
		return &FileAdapter{
			Adapter: s3.NewUploadAliyunOSSAdapter(c),
		}
	} else if c.S3.Type == "Local" {
		return &FileAdapter{
			Adapter: s3.NewLocalFileAdapter(c.S3.Local.BasePath),
		}
	}
	return nil
}
