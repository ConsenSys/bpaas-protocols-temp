package user

//CreateUserRequest parameters
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type CreateUserMessage struct {
	Uuid string `json:uuid`
	CreateUserRequest
}
