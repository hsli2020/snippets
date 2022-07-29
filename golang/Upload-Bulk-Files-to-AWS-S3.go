package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type Session struct {
    S3Session *session.Session
}

func main() {
    paths := []string{"", ""}
    credential := credentials.NewStaticCredentials(
        os.Getenv("SECRET_ID"),
        os.Getenv("SECRET_KEY"),
        "",
    )
    awsConfig := aws.Config{
        Region:      aws.String(os.Getenv("REGION")),
        Credentials: credential,
    }

    s, err := session.NewSession(&awsConfig)
    if err != nil {
        log.Println("failed to create S3 session:", err.Error())
    }

    se := Session{s}
    err = se.upload(paths)
    if err != nil {
        log.Println(err.Error())
        return
    }
}

func (s Session) upload(paths []string) error {
    for _, path := range paths {
        upFile, err := os.Open(path)
        if err != nil {
            log.Printf("failed %s, error: %v", path, err.Error())
            continue
        }
        defer upFile.Close()

        upFileInfo, err := upFile.Stat()
        if err != nil {
            log.Printf("failed to get stat %s, error: %v", path, err.Error())
            continue
        }

        var fileSize int64 = upFileInfo.Size()
        fileBuffer := make([]byte, fileSize)
        upFile.Read(fileBuffer)

        // uploading
        _, err = s3.New(s.S3Session).PutObject(&s3.PutObjectInput{
            Bucket:               aws.String(os.Getenv("BUCKET_NAME")),
            Key:                  aws.String(path),
            ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
            Body:                 bytes.NewReader(fileBuffer),
            ContentLength:        aws.Int64(int64(fileSize)),
            ContentType:          aws.String(http.DetectContentType(fileBuffer)),
            ContentDisposition:   aws.String("attachment"),
            ServerSideEncryption: aws.String("AES256"),
            StorageClass:         aws.String("INTELLIGENT_TIERING"),
        })
        if err != nil {
            log.Printf("failed to upload %s, error: %v", path, err.Error())
            continue
        }
        url := "https://%s.s3-%s.amazonaws.com/%s"
        url = fmt.Sprintf(url, os.Getenv("BUCKET_NAME"), os.Getenv("REGION"), path)
        fmt.Printf("Uploaded File Url %s\n", url)
    }

    return nil
}