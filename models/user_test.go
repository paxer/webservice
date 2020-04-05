package models

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	users = []*User{}
	result := GetUsers()
	if result == nil {
		t.Error("Did not get a users collection")
	}
}

func TestAddUserWithID(t *testing.T) {
	u := User{ID: 1, FirstName: "Darth", LastName: "Vader"}
	addedUser, err := AddUser(u)
	if err.Error() != "new User must not include id" {
		t.Error("Expected to return an error object")
	}
	emptyUser := User{}
	if addedUser != emptyUser {
		t.Error("Expected to return an empty user struct")
	}
}

func TestAddUserWithoutID(t *testing.T) {
	u := User{FirstName: "Darth", LastName: "Vader"}
	result, err := AddUser(u)
	if err != nil {
		t.Error("Did not expect and error")
	}
	newAddedUser := User{ID: 1, FirstName: "Darth", LastName: "Vader"}
	if result != newAddedUser {
		t.Error("Expected to return a created user")
	}
}

func TestGetUserByIDFound(t *testing.T) {
	users = nil
	existedUser := &User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	users = append(users, existedUser)
	result, err := GetUserByID(1)
	if result != *existedUser {
		t.Error("Expected to find a user")
	}
	if err != nil {
		t.Error("Expected err to be nil")
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	users = nil
	existedUser := &User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	users = append(users, existedUser)
	result, err := GetUserByID(2)
	emptyUser := User{}
	if result != emptyUser {
		t.Error("Expected result to be empty user")
	}
	if err.Error() != "user with ID '2' not found" {
		t.Error("Expected err to be returned")
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	users = nil
	existedUser := &User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	updatedUser := User{ID: 1, FirstName: "Darth", LastName: "Skywalker"}
	users = append(users, existedUser)
	result, err := UpdateUser(updatedUser)
	if *users[0] != result {
		t.Error("Expected user to be updated in users collection")
	}
	if result != updatedUser {
		t.Error("Expected result to be updated user")
	}
	if err != nil {
		t.Error("Expected err to be nil")
	}
}

func TestUpdateUserError(t *testing.T) {
	users = nil
	userToUpdate := User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	result, err := UpdateUser(userToUpdate)
	emptyUser := User{}
	if result != emptyUser {
		t.Error("Expected result to be empty user")
	}
	if err.Error() != "user with ID '1' not found" {
		t.Error("Expected err to be returned")
	}
}

func TestRemoveUserByIDSuccess(t *testing.T) {
	users = nil
	existedUser := &User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	users = append(users, existedUser)
	result := RemoveUserByID(1)
	if users[0] != nil {
		t.Error("Expected user to be removed form users collection")
	}
	if result != nil {
		t.Error("Expected result to be nil")
	}
}

func TestRemoveUserByIDError(t *testing.T) {
	users = nil
	existedUser := &User{ID: 1, FirstName: "Darth", LastName: "Wader"}
	users = append(users, existedUser)
	err := RemoveUserByID(2)

	if users[0] != existedUser {
		t.Error("Expected users collection to stay the same")
	}

	if err.Error() != "user with ID '2' not found" {
		t.Error("Expected result to be an error")
	}
}
