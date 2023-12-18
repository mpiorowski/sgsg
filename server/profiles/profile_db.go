package profiles

import (
	"database/sql"
	"fmt"
	"sgsg/db"
	pb "sgsg/proto"

	"github.com/google/uuid"
)

func scanProfile(rows *sql.Rows, row *sql.Row) (*pb.Profile, error) {
	profile := pb.Profile{}
	if rows != nil {
		err := rows.Scan(&profile.Id, &profile.Created, &profile.Updated, &profile.Deleted, &profile.UserId, &profile.Username, &profile.About, &profile.ResumeId, &profile.CoverId, &profile.CoverUrl)
		if err != nil {
			return nil, err
		}
	} else {
		err := row.Scan(&profile.Id, &profile.Created, &profile.Updated, &profile.Deleted, &profile.UserId, &profile.Username, &profile.About, &profile.ResumeId, &profile.CoverId, &profile.CoverUrl)
		if err != nil {
			return nil, err
		}
	}
	return &profile, nil
}

func selectProfileByUserId(userId string) (*pb.Profile, error) {
	row := db.Db.QueryRow("select * from profiles where user_id = $1", userId)
	profile, err := scanProfile(nil, row)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("scanProfile: %w", err)
	}
	if err == sql.ErrNoRows {
		return &pb.Profile{}, nil
	}
	return profile, nil
}

func insertProfile(in *pb.Profile) (*pb.Profile, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("uuid.NewRandom: %w", err)
	}
	row := db.Db.QueryRow(
		"insert into profiles (id, user_id, username, about, resume_id, cover_id, cover_url) values ($1, $2, $3, $4, $5, $6, $7) returning *",
		id,
		in.UserId,
		in.Username,
		in.About,
		in.ResumeId,
        in.CoverId,
		in.CoverUrl,
	)
	profile, err := scanProfile(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanProfile: %w", err)
	}
	return profile, nil
}

func updateProfile(in *pb.Profile) (*pb.Profile, error) {
	row := db.Db.QueryRow(
		"update profiles set username = $1, about = $2, resume_id = $3, cover_id = $4, cover_url = $5 where id = $6 and user_id = $7 returning *",
		in.Username,
		in.About,
		in.ResumeId,
        in.CoverId,
		in.CoverUrl,
		in.Id,
		in.UserId,
	)
	profile, err := scanProfile(nil, row)
	if err != nil {
		return nil, fmt.Errorf("scanProfile: %w", err)
	}
	return profile, nil
}
