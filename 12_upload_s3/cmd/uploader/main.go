package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

// iniciada antes da função main
func init() {
	cfg := &aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	}

	// cria uma sessão para se comunicar com a aws
	sess, err := session.NewSession(cfg)
	if err != nil {
		panic(err)
	}

	// criando o client para o s3
	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	// controlador de uploads simultaneos
	uploadControl := make(chan struct{}, 5)

	// controla a retentativa de arquivos com erro
	errorFileUpload := make(chan string, 3)

	go retryUpload(uploadControl, errorFileUpload)

	for {
		// lendo o conteudo de um diretório, nesse cado pegando apenas 1 por vez
		files, err := dir.ReadDir(1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error reading directory: %s\n", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func retryUpload(uploadControl chan struct{}, errorFileUpload chan string) {
	for filename := range errorFileUpload {
		uploadControl <- struct{}{}
		wg.Add(1)
		go uploadFile(filename, uploadControl, errorFileUpload)
	}
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("error opening file %s\n", completeFileName)
		<-uploadControl             // esvaziando o channel
		errorFileUpload <- filename // adiciona no channel para retentativa
		return
	}
	defer f.Close()

	// enviando arquivo para o s3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("error uploading file %s\n", completeFileName)
		<-uploadControl             // esvaziando o channel
		errorFileUpload <- filename // adiciona no channel para retentativa
		return
	}
	fmt.Printf("file %s uploaded successffuly\n", completeFileName)
	<-uploadControl // esvaziando o channel
}
