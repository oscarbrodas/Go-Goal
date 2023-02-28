package handler

// func TestLoginUsernameFail(t *testing.T) { // Incorrect Username
// 	loginInfo := loginInput{
// 		"thescar1o1",
// 		"123456",
// 	}

// 	accounts := user.New()
// 	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})

// 	if LoginCheck(loginInfo, accounts.Users) != nil {
// 		t.Errorf("Duplicate User was not found")
// 	}
// }

// func TestLoginPasswordFail(t *testing.T) { // Incorrect Password

// 	loginInfo := loginInput{
// 		"thescar101",
// 		"12345",
// 	}

// 	accounts := user.New()
// 	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})

// 	if LoginCheck(loginInfo, accounts.Users) != nil {
// 		t.Errorf("Duplicate User was not found")
// 	}
// }

// func TestLoginSuccess(t *testing.T) { // Successful Login
// 	loginInfo := loginInput{
// 		"thescar101",
// 		"123456",
// 	}

// 	accounts := user.New()
// 	accounts.Add(user.User{Username: "thescar101", FirstName: "Oscar", LastName: "Rodas", Email: "oscarrodas@ufl.edu", Password: "123456"})

// 	if LoginCheck(loginInfo, accounts.Users) == nil {
// 		t.Errorf("Duplicate User was not found")
// 	}
// }
