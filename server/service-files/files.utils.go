package main

import (
	"context"
	"database/sql"
	pb "go-svelte-grpc/proto"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

func mapFile(rows *sql.Rows, row *sql.Row) (*pb.File, error) {
	var file pb.File = pb.File{}
	var err error
	if rows != nil {
		err = rows.Scan(
			&file.Id,
			&file.Created,
			&file.Updated,
			&file.Deleted,
			&file.TargetId,
			&file.Name,
			&file.Type,
		)
	} else if row != nil {
		err = row.Scan(
			&file.Id,
			&file.Created,
			&file.Updated,
			&file.Deleted,
			&file.TargetId,
			&file.Name,
			&file.Type,
		)
	}
	if err != nil {
		return &file, err
	}
	return &file, nil
}

// uploadFile uploads an object.
func uploadFile(object string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("storage.NewClient: %v", err)
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	o := client.Bucket(BUCKET).Object(object)
	wc := o.NewWriter(ctx)
	_, err = wc.Write(data)
	if err != nil {
		log.Printf("wc.Write: %v", err)
		return err
	}
	err = wc.Close()
	if err != nil {
		log.Printf("wc.Close: %v", err)
		return err
	}

	log.Printf("File %v uploaded to bucket %v.", object, BUCKET)
	return nil
}

func generateV4GetObjectSignedURL(object string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("storage.NewClient: %v", err)
		return "", err
	}
	defer client.Close()

	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	u, err := client.Bucket(BUCKET).SignedURL(object, opts)
	if err != nil {
		log.Printf("storage.SignedURL: %v", err)
		return "", err
	}
	log.Printf("Generated GET signed URL: %v", u)
	return u, nil
}
