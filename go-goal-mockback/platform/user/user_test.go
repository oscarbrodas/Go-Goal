package user

import "testing"

func TestAdd(t *testing.T) {
	accounts := New()
	accounts.Add(User{})
	if len(accounts.Users) != 1 {
		t.Errorf("User was not added")
	}
}

func TestGetAll(t *testing.T) {
	accounts := New()
	accounts.Add(User{})
	results := accounts.GetAll()
	if len(results) != 1 {
		t.Errorf("User was not added")
	}
}
