package utils

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

func Upload() {
	endpoint := "8.130.66.162:9000"
	accessKeyID := "rYoZNgAhrvV83QlLowxS"
	secretAccessKey := "BSFf09j2jCJvdp21Sxzqqsvn5wMdeD45F3lFeh9y"
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
	objectName := "defalut.jpg"
	filePath := "C:\\Users\\kele\\Desktop\\OIP-C.jpg"

	// 使用文件路径上传
	_, err = client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File uploaded successfully")
}
