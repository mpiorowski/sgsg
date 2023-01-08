package main

import (
	"context"
	"database/sql"
	"log"

	pb "go-svelte-grpc/proto"
)

func (s *server) GetFiles(in *pb.TargetId, stream pb.FilesService_GetFilesServer) error {

	rules := map[string]string{
		"TargetId": "required,max=100,uuid",
	}
	validate.RegisterStructValidationMapRules(rules, pb.TargetId{})
	err := validate.Struct(in)
	if err != nil {
		log.Printf("validate.Struct: %v", err)
		return err
	}

	rows, err := db.Query(`select * from files where "targetId" = $1`, in.TargetId)
	if err != nil {
		log.Printf("db.Query: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		file, err := mapFile(rows, nil)
		if err != nil {
			log.Printf("mapFile: %v", err)
			return err
		}
		if ENV == "production" {
			file.Url, err = generateV4GetObjectSignedURL(file.TargetId + "/" + file.Name)
			if err != nil {
				log.Printf("generateV4GetObjectSignedURL: %v", err)
				return err
			}
		} else {
            file.Url = ""
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

	rules := map[string]string{
		"TargetId": "required,max=100,uuid",
		"Name":     "required,max=100",
		"Type":     "required,max=100",
		"Data":     "required",
	}
	validate.RegisterStructValidationMapRules(rules, pb.File{})
	err := validate.Struct(in)
	if err != nil {
		log.Printf("validate.Struct: %v", err)
		return nil, err
	}

	if ENV == "production" {
		path := in.TargetId + "/" + in.Name
		err = uploadFile(path, in.Data)
		if err != nil {
			log.Printf("uploadFile: %v", err)
			return nil, err
		}
	}

	// check if file exists
	var exists bool
	var row *sql.Row
	err = db.QueryRow(`select exists(select 1 from files where "targetId" = $1 and "name" = $2)`, in.TargetId, in.Name).Scan(&exists)
	if err != nil {
		log.Printf("db.QueryRow: %v", err)
		return nil, err
	}

	if exists {
		row = db.QueryRow(`update files set type = $1 where "targetId" = $2 and "name" = $3 returning *`,
			in.Type,
			in.TargetId,
			in.Name,
		)
	} else if in.Id != "" {
		row = db.QueryRow(`update files set "name" = $1, "type" = $2 where "id" = $3 and "targetId" = $4 returning *`,
			in.Name,
			in.Type,
			in.Id,
			in.TargetId,
		)
	} else {
		row = db.QueryRow(`insert into files ("targetId", "name", "type") values ($1, $2, $3) returning *`,
			in.TargetId,
			in.Name,
			in.Type,
		)
	}
	file, err := mapFile(nil, row)
	if err != nil {
		log.Printf("mapFile: %v", err)
		return nil, err
	}
	return file, nil
}

// TODO - delete form bucket
func (s *server) DeleteFile(ctx context.Context, in *pb.FileId) (*pb.File, error) {

	row := db.QueryRow(`update files set "deleted" = now() where "id" = $1 and "targetId" = $2 returning *`, in.FileId, in.TargetId)

	file, err := mapFile(nil, row)
	if err != nil {
		log.Printf("mapFile: %v", err)
		return nil, err
	}

	return file, nil
}
