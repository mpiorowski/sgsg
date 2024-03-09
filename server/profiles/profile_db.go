package profiles

import (
	pb "sgsg/proto"
	"sgsg/system"

	"github.com/google/uuid"
)

type ProfileDBProvider interface {
	SelectProfileByUserID(userID string) (*pb.Profile, error)
	InsertProfile(profile *pb.Profile) (*pb.Profile, error)
	UpdateProfile(profile *pb.Profile) (*pb.Profile, error)
}

type ProfileDBImpl struct {
	*system.Storage
}

func NewProfileDB(storage *system.Storage) ProfileDBProvider {
	return &ProfileDBImpl{storage}
}

func (db *ProfileDBImpl) SelectProfileByUserID(userID string) (*pb.Profile, error) {
	row := db.Conn.QueryRow("select * from profiles where user_id = $1", userID)
	var profile pb.Profile
	err := row.Scan(dest(&profile)...)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (db *ProfileDBImpl) InsertProfile(in *pb.Profile) (*pb.Profile, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	row := db.Conn.QueryRow(`
		insert into profiles (id, user_id, username, about, resume_id, cover_id, cover_url)
		values ($1, $2, $3, $4, $5, $6, $7)
		returning *
	`, id, in.UserId, in.Username, in.About, in.ResumeId, in.CoverId, in.CoverUrl)
	var profile pb.Profile
	err = row.Scan(dest(&profile)...)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (db *ProfileDBImpl) UpdateProfile(in *pb.Profile) (*pb.Profile, error) {
	row := db.Conn.QueryRow(`
		update profiles
		set username = $1, about = $2, resume_id = $3, cover_id = $4, cover_url = $5
		where id = $6 and user_id = $7
		returning *
	`, in.Username, in.About, in.ResumeId, in.CoverId, in.CoverUrl, in.Id, in.UserId)

	var profile pb.Profile
	err := row.Scan(dest(&profile)...)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func dest(profile *pb.Profile) []interface{} {
	return []interface{}{
		&profile.Id,
		&profile.Created,
		&profile.Updated,
		&profile.Deleted,
		&profile.UserId,
		&profile.Username,
		&profile.About,
		&profile.ResumeId,
		&profile.CoverId,
		&profile.CoverUrl,
	}
}
