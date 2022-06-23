package dto

import "orangezoom.cn/ginessential/model"

type UserDTO struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDTO(user model.User) UserDTO {
	dto := UserDTO{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
	return dto
}
