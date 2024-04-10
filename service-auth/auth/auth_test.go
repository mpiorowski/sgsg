package auth

import (
	pb "service-auth/proto"
	"service-auth/system"
	"testing"
)

var users = []*pb.User{
	{
		Email:  "john@gmail.com",
		Role:   pb.Role_ROLE_USER,
		Sub:    "123",
		Avatar: "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
	},
	{
		Email:  "jane@gmail.com",
		Role:   pb.Role_ROLE_USER,
		Sub:    "456",
		Avatar: "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50",
	},
}

func TestMain(m *testing.M) {
	storage := system.NewMemoryStorage()
	err := storage.Migrations()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func setup() AuthDB {
	storage := system.NewMemoryStorage()
	_, err := storage.Conn.Exec("delete from users")
	if err != nil {
		panic(err)
	}
	var db = NewAuthDB(&storage)
	return db
}

func TestInsertUsers(t *testing.T) {
	db := setup()

	// Test case 1: Insert users
	for _, user := range users {
		newUser, err := db.insertUser(user.Email, user.Sub, user.Avatar)
		if err != nil {
			t.Error(err)
		}
		if newUser.Id == "" {
			t.Error("User id is empty")
		}
		equal := newUser.Email == user.Email && newUser.Sub == user.Sub && newUser.Avatar == user.Avatar
		if !equal {
			t.Error("User is not equal")
		}
	}

	// Test case 2: Insert duplicate user
	_, err := db.insertUser(users[0].Email, users[0].Sub, users[0].Avatar)
	if err == nil {
		t.Error("Duplicate user is inserted")
	}

	// Test case 3: Insert user with empty email
	_, err = db.insertUser("", users[0].Sub, users[0].Avatar)
	if err == nil {
		t.Error("User with empty email is inserted")
	}

	// Test case 4: Insert user with empty sub
	_, err = db.insertUser(users[0].Email, "", users[0].Avatar)
	if err == nil {
		t.Error("User with empty sub is inserted")
	}

	// Test case 5: Insert user with empty avatar
	_, err = db.insertUser("mat@gmail.com", "789", "")
	if err != nil {
		t.Error("User with empty avatar is not inserted")
	}
}

func TestSelectUsers(t *testing.T) {
	db := setup()

	// Test case 1: Select users
	for _, user := range users {
		u, err := db.insertUser(user.Email, user.Sub, user.Avatar)
		if err != nil {
			t.Error(err)
		}
		newUser, err := db.selectUserById(u.Id)
		if err != nil {
			t.Error(err)
		}
		if newUser.Id == "" {
			t.Error("User id is empty")
		}
		equal := newUser.Email == user.Email && newUser.Sub == user.Sub && newUser.Avatar == user.Avatar
		if !equal {
			t.Error("User is not equal")
		}
	}

	// Test case 2: Select non-existing user
	_, err := db.selectUserById("789")
	if err == nil {
		t.Error("Non-existing user is selected")
	}

	// Test case 3: Select user by email and sub
	newUser, err := db.selectUserByEmailAndSub(users[0].Email, users[0].Sub)
	if err != nil {
		t.Error(err)
	}
	if newUser.Id == "" {
		t.Error("User id is empty")
	}
	equal := newUser.Email == users[0].Email && newUser.Sub == users[0].Sub && newUser.Avatar == users[0].Avatar
	if !equal {
		t.Error("User is not equal")
	}

	// Test case 4: Select non-existing user by email and sub
	_, err = db.selectUserByEmailAndSub("", users[0].Sub)
	if err == nil {
		t.Error("Non-existing user is selected")
	}
}
