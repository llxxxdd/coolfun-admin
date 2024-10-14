package s3

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
)

// LocalFileAdapter 本地文件存储适配器
type LocalFileAdapter struct {
	BasePath string
}

// NewLocalFileAdapter 创建新的本地文件适配器
func NewLocalFileAdapter(basePath string) *LocalFileAdapter {
	return &LocalFileAdapter{
		BasePath: basePath,
	}
}

func (l *LocalFileAdapter) UploadFile(ctx context.Context, file *FileInfo) (fileInfo *FileInfo, err error) {
	// 构建本地存储路径
	filePath := filepath.Join(l.BasePath, file.Path, file.Name)

	// 创建目标文件夹（如果不存在）
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create directories for file")
	}

	// 打开或创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create local file")
	}
	defer f.Close()

	// 将数据写入本地文件
	if file.Reader != nil {
		_, err = io.Copy(f, *file.Reader)
		if err != nil {
			return nil, errors.Wrap(err, "failed to write file data")
		}
	}

	file.Path = filePath
	return file, nil
}

// SignedUrl 本地文件不需要签名 URL，直接返回文件路径
func (l *LocalFileAdapter) SignedUrl(ctx context.Context, filePath string) (signedURL string, err error) {
	// 本地存储直接返回文件的本地路径或 URL
	return filepath.Join(l.BasePath, filePath), nil
}

// FileReader 读取本地文件
func (l *LocalFileAdapter) FileReader(ctx context.Context, filePath string) (fileReader io.ReadCloser, err error) {
	// 打开本地文件进行读取
	file, err := os.Open(filepath.Join(l.BasePath, filePath))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open local file")
	}
	return file, nil
}

// DeleteFile 删除本地文件
func (l *LocalFileAdapter) DeleteFile(ctx context.Context, filePaths []string) (deletedObjects []string, err error) {
	for _, filePath := range filePaths {
		fullPath := filepath.Join(l.BasePath, filePath)
		err = os.Remove(fullPath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete local file")
		}
		deletedObjects = append(deletedObjects, fullPath)
	}
	return deletedObjects, nil
}
