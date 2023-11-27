package brewery

type CreateBreweryInputDto struct {
	Name string `json:"name"`
	Email string `json:"email"`
	UserId string `json:"user_id"`
}

type CreateBreweryOutputDto struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type GetUserBreweriesOutputDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}

type InviteUserInputDTO struct {
	InvitingUserId string `json:"inviting_user_id"`
	InvitedUserEmail string `json:"invited_user_email"`
	BreweryName string `json:"brewery_name"`
	Role string `json:"role"`
}

type InviteUserRepoInputDto struct {
	InvitingUserId string
	InvitedUserId string
	BreweryId string
	Role string
}

type InviteUserOutputDTO struct {
	InvitingUserId string `json:"inviting_user_id"`
	InvitedUserId string `json:"invited_user_id"`
	BreweryId string `json:"brewery_id"`
	Role string `json:"role"`
	ErrorMessage string `json:"error_message,omitempty"`
	StatusCode int `json:"statusCode,omitempty"`
}