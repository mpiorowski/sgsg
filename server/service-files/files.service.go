package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"cloud.google.com/go/storage"
	pb "go-svelte-grpc/server/grpc"
)

func mapFile(rows *sql.Rows) (*pb.File, error) {
	var file pb.File = pb.File{}
	err := rows.Scan(
		&file.Id,
		&file.Created,
		&file.Updated,
		&file.Deleted,
		&file.TargetId,
		&file.Name,
		&file.Type,
	)
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

func (s *server) GetFiles(in *pb.TargetId, stream pb.FilesService_GetFilesServer) error {
	rows, err := db.Query(`select * from files where "targetId" = $1`, in.TargetId)
	if err != nil {
		log.Printf("db.Query: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		file, err := mapFile(rows)
		if err != nil {
			log.Printf("mapFile: %v", err)
			return err
		}
		file.Url, err = generateV4GetObjectSignedURL(file.TargetId + "/" + file.Name)
		if err != nil {
			log.Printf("generateV4GetObjectSignedURL: %v", err)
			return err
		}
		err = stream.Send(file)
		if err != nil {
			log.Printf("stream.Send: %v", err)
			return err
		}
	}
	if rows.Err() != nil {
		log.Printf("rows.Err: %v", err)
		return rows.Err()
	}
	return nil
}

func (s *server) CreateFile(ctx context.Context, in *pb.File) (*pb.File, error) {
	var err error
	var rows *sql.Rows

	path := in.TargetId + "/" + in.Name
	err = uploadFile(path, in.Data)
	if err != nil {
		log.Printf("uploadFile: %v", err)
		return nil, err
	}

	if in.Id == "" {
		rows, err = db.Query(`insert into files ("targetId", "name", "type") values ($1, $2, $3) returning *`,
			in.TargetId,
			in.Name,
			in.Type,
		)
	} else {
		rows, err = db.Query(`update files set "name" = $1, "type" = $2 where "id" = $3 and "targetId" = $4 returning *`,
			in.Name,
			in.Type,
			in.Id,
			in.TargetId,
		)
	}

	if err != nil {
		log.Printf("db.Query: %v", err)
		return nil, err
	}
	defer rows.Close()

	rows.Next()
	file, err := mapFile(rows)
	if err != nil {
		log.Printf("mapFile: %v", err)
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		log.Printf("rows.Err: %v", err)
        return nil, err
	}

	return file, nil
}

func (s *server) DeleteFile(ctx context.Context, in *pb.FileId) (*pb.File, error) {
	var err error
	var rows *sql.Rows

	rows, err = db.Query(`delete from files where "id" = $1 and "targetId" = $2 returning *`,
		in.FileId,
		in.TargetId,
	)

	if err != nil {
		log.Printf("db.Query: %v", err)
		return nil, err
	}
	defer rows.Close()

	rows.Next()
	file, err := mapFile(rows)
	if err != nil {
		log.Printf("mapFile: %v", err)
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		log.Printf("rows.Err: %v", err)
        return nil, err
	}

	return file, nil
}
