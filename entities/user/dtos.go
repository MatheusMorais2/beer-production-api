package user

type CreateUserInputDto struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type CreateUserOutputDto struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type LoginUserInputDto struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUserOutputDto struct {
	Email string `json:"email"`
	Token string `json:"token"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}