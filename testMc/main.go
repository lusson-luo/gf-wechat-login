package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	s3client := S3Client{
		Endpoint:        "s3.ap-southeast-1.qiniucs.com",
		AccessKeyID:     "xxx",
		SecretAccessKey: "xxx-WeCGHWX8D0b",
		UseSSL:          true,
	}
	url, err := s3client.UploadFile(context.Background(), "tuboshu-static3", "ap-southeast-1", "20230902001449-201934.png", "../tempFile/20230902001449-201934.png")
	fmt.Println(url, err)
}

type S3Client struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func New(endpoint, accessKeyID, secretAccessKey string, useSSL bool) *S3Client {
	return &S3Client{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
	}
}

func (client *S3Client) UploadFile(ctx context.Context, bucketName, location, objectName, filePath string) (s3PublicURL string, err error) {
	glog.Info(ctx, "创建 s3client", client)
	// Initialize minio client object.
	minioClient, err := minio.New(client.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(client.AccessKeyID, client.SecretAccessKey, ""),
		Secure: client.UseSSL,
	})
	if err != nil {
		glog.Fatal(ctx, err)
		return "", err
	}
	glog.Info(ctx, "创建 s3 连接成功", client.Endpoint)

	// Make a new bucket.
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})

	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		glog.Infof(ctx, "桶是否存在 %v，err=%v\n", exists, errBucketExists)
		if exists {
			glog.Infof(ctx, "We already own %s\n", bucketName)
		}
	} else {
		glog.Infof(ctx, "Successfully created %s\n", bucketName)
	}

	// Upload the file with FPutObject
	glog.Printf(ctx, "putfile bucketName  %s, location:%s,  objectName  %s, filePath %s\n", bucketName, location, objectName, filePath)
	info, err := minioClient.FPutObject(ctx, bucketName, "images/"+objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		glog.Info(ctx, err)
		return "", err
	}

	glog.Printf(ctx, "Successfully uploaded %s of size %d\n", objectName, info.Size)
	return GetPublicURL(ctx, objectName), nil
}

// 七牛云获得访问地址
func GetPublicURL(ctx context.Context, filename string) string {
	config, err := gcfg.New()
	if err != nil {
		glog.Error(ctx, "创建 config 对象失败", err)
	}
	localUrl := config.MustGet(ctx, "storage.local.domain", false).String()
	return localUrl + "/" + filename
}
