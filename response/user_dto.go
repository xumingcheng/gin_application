package response

import "github.com/xumingcheng/gin_application/model"

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}
func TouserDto(user model.User)UserDto{
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}