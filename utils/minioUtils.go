package utils

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"strings"
)

func Upload(objectName string, file bytes.Buffer) {
	endpoint := "8.130.66.162:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio123456"
	useSSL := false

	// 初始化 MinIO 客户端
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// 上传文件
	bucketName := "douyin"
	//filePath := "C:\\Users\\kele\\Desktop\\defalut.jpg"

	// 使用字节内容上传
	_, err = client.PutObject(context.Background(), bucketName, objectName, &file, int64(file.Len()), minio.PutObjectOptions{ContentType: getFileContentType(objectName)})
	if err != nil {
		log.Fatalln(err)
	}
	// 使用文件路径上传
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File uploaded successfully")
}
func getFileContentType(fileName string) string {
	returnFileName := fileName[strings.LastIndex(fileName, "."):]
	if returnFileName != "" {
		switch returnFileName {
		case ".jpeg", ".png", ".jpg":
			return "image/jpeg"
		case ".mp4":
			return "video/mp4"
		case ".html":
			return "text/html"
		case ".css":
			return "text/css"
		case ".js":
			return "application/javascript"
		case ".pdf":
			return "application/pdf"
		default:
			return "application/octet-stream"
		}
	}
	return ""
}
