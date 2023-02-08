package handler

import (
	"go-goal/platform/user"
	"testing"
)

func TestFoundUsername(t *testing.T) { // Matching Username
	accounts := user.New()
	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})
	if !FoundUser(user.User{Username: "thescar101", FirstName: "xxxxx", LastName: "xxxxx", Email: "xxxxxxxx@xxxx.xxx", Password: "xxx"}, accounts.Users) {
		t.Errorf("Duplicate User was not found")
	}
}

func TestFoundEmail(t *testing.T) { // Matching Email
	accounts := user.New()
	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})
	if !FoundUser(user.User{Username: "xxxxx", FirstName: "xxxxx", LastName: "xxxxx", Email: "oscarrodas@ufl.edu", Password: "xxx"}, accounts.Users) {
		t.Errorf("Duplicate User was not found")
	}
}

func TestFoundNoOne(t *testing.T) { // Not Matching
	accounts := user.New()
	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})
	if FoundUser(user.User{Username: "xxxxx", FirstName: "xxxxx", LastName: "xxxxx", Email: "xxxxxxxx@xxxx.xxx", Password: "xxx"}, accounts.Users) {
		t.Errorf("Duplicate User was not found")
	}
}
