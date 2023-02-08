package user

type Getter interface {
	GetAll() []User
}

type Adder interface {
	GetAll() []User
	Add(user User)
}

// User data
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// User slice
type Repo struct {
	Users []User
}

// Returns a pointer to an empty User slice
func New() *Repo {
	return &Repo{
		Users: []User{},
	}
}

func (r *Repo) Add(user User) {
	r.Users = append(r.Users, user)
}

func (r *Repo) GetAll() []User {
	return r.Users
}
