package user

//CreateUserRequest parameters
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type CreateUserDTO struct {
	Uuid string `json:uuid`
	User
}
