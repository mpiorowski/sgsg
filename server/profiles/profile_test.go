package profiles

import (
	"sgsg/db"
	pb "sgsg/proto"
	"strings"
	"testing"
)

var profiles = []pb.Profile{
	{
		UserId:   "test",
		Username: "test username 1",
		About:    "test about 1",
		ResumeId:   "test resume 1",
		CoverUrl:    "test cover 1",
	},
	{
		UserId:   "test",
		Username: "test username 2",
		About:    "test about 2",
		ResumeId:   "test resume 2",
		CoverUrl:    "test cover 2",
	},
	{
		UserId:   "test",
		Username: "test username 3",
		About:    "test about 3",
		ResumeId:   "test resume 3",
		CoverUrl:    "test cover 3",
	},
}

func TestMain(m *testing.M) {
	err := db.ConnectTest()
	if err != nil {
		panic(err)
	}
	err = db.Migrations()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func clearProfiles() {
    _, err := db.Db.Exec("delete from profiles")
    if err != nil {
        panic(err)
    }
}

func TestInsertProfile(t *testing.T) {
	// Test case 1: Insert a valid profile
	profile, err := insertProfile(&profiles[0])
	if err != nil {
		t.Errorf("insertProfile error: %v", err)
	}
	equal := profile.Username == profiles[0].Username &&
		profile.About == profiles[0].About &&
		profile.ResumeId == profiles[0].ResumeId &&
		profile.CoverUrl == profiles[0].CoverUrl

	if !equal {
		t.Errorf("insertProfile error: not equal")
	}

	// Test case 2: Insert a second valid profile
	profile, err = insertProfile(&profiles[1])
	if err != nil {
		t.Errorf("insertProfile error: %v", err)
	}
	if profile.Username != profiles[1].Username {
		t.Errorf("insertProfile error: title not equal")
	}
}

func TestUpdateProfile(t *testing.T) {
	// Test case 1: Update a valid profile
	profile, err := insertProfile(&profiles[0])
	if err != nil {
		t.Errorf("insertProfile error: %v", err)
	}
	newProfile := pb.Profile{
		Id:       profile.Id,
		UserId:   profile.UserId,
		Username: "New username",
		About:    "New about",
		ResumeId:   "New resume",
		CoverUrl:    "New cover",
	}
	profile, err = updateProfile(&newProfile)
	if err != nil {
		t.Errorf("updateProfileTitle error: %v", err)
	}
	if profile.Username != newProfile.Username {
		t.Errorf("updateProfileTitle error: title not equal")
	}

	// Test case 2: Update a profile that does not exist
	newProfile = pb.Profile{
		Id: "not_exist",
	}
	_, err = updateProfile(&newProfile)
	if err == nil {
		t.Errorf("updateProfileTitle error: %v", err)
	}
}

func TestDeleteProfileById(t *testing.T) {
    clearProfiles()
	// Test case 1: Delete a profile
	profile, err := insertProfile(&profiles[0])
	if err != nil {
		t.Errorf("insertProfile error: %v", err)
	}
	err = deleteProfileById(profile.Id)
	if err != nil {
		t.Errorf("deleteProfileById error: %v", err)
	}
	profile, err = selectProfileByUserId(profile.UserId)
	if profile.Id != "" || err != nil {
		t.Errorf("selectProfileByUserId error: %v", err)
	}

	// Test case 2: Delete a profile that does not exist
	err = deleteProfileById("not_exist")
	if err == nil {
		t.Errorf("deleteProfileById error: %v", err)
	}
}

func TestSelectProfileyId(t *testing.T) {
    clearProfiles()
	// Test case 1: Select a profile by id
	newProfile, _ := insertProfile(&profiles[2])
	profile, err := selectProfileByUserId(newProfile.UserId)
	if err != nil {
		t.Errorf("selectProfileByUserId error: %v", err)
	}
	if profile.Id != newProfile.Id {
		t.Errorf("selectProfileByUserId error: not equal")
	}

    // Test case 2: Select a profile by id that does not exist
    profile, err = selectProfileByUserId("not_exist")
    if err != nil && profile.Id != "" {
        t.Errorf("selectProfileByUserId error: %v", err)
    }
}

func TestProfileValidation(t *testing.T) {
	profiles[0].UserId = ""
	profiles[0].Username = ""
    profiles[0].About = ""
    err := validateProfile(&profiles[0])
	containsTitle := strings.Contains(err.Error(), "Username") && strings.Contains(err.Error(), "required")
	containsContent := strings.Contains(err.Error(), "About") && strings.Contains(err.Error(), "required")
	if !containsTitle || !containsContent {
		t.Errorf("validation error: %v", err)
	}

	// gen 101 chars
	profiles[0].Username = strings.Repeat("a", 101)
	profiles[0].About = strings.Repeat("a", 1001)
    err = validateProfile(&profiles[0])
	containsTitle = strings.Contains(err.Error(), "Username") && strings.Contains(err.Error(), "max")
	containsContent = strings.Contains(err.Error(), "About") && strings.Contains(err.Error(), "max")
	if !containsTitle || !containsContent {
		t.Errorf("validation error: %v", err)
	}
}

func TestConcurrency(t *testing.T) {
	newProfiles := make([]*pb.Profile, 0)
	// Test case 1: Insert a profile concurrently
	profilesChanel := make(chan *pb.Profile)
	gooroutines := 10
	for i := 0; i < gooroutines; i++ {
		go func() {
			newProfile, err := insertProfile(&profiles[0])
			if err != nil {
				t.Errorf("insertProfile error: %v", err)
			}
			profilesChanel <- newProfile
		}()
	}

	for i := 0; i < gooroutines; i++ {
		n := <-profilesChanel
		newProfiles = append(newProfiles, n)
	}

	// Test case 2: Update a profile concurrently
	done := make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			profile := newProfiles[i]
			profile.Username = "New username"
			_, err := updateProfile(profile)
			if err != nil {
				t.Errorf("updateProfileTitle error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}

	// Test case 4: Select profile by id concurrently
	done = make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			profile := newProfiles[i]
			_, err := selectProfileByUserId(profile.UserId)
			if err != nil {
				t.Errorf("selectProfiles error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}

	// Test case 3: Delete a profile concurrently
	done = make(chan bool)
	for i := 0; i < gooroutines; i++ {
		go func(i int) {
			profile := newProfiles[i]
			err := deleteProfileById(profile.Id)
			if err != nil {
				t.Errorf("deleteProfileById error: %v", err)
			}
			done <- true
		}(i)
	}

	for i := 0; i < gooroutines; i++ {
		<-done
	}
}
