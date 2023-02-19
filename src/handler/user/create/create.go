package user_handler

import (
	custom_error "main/src/entities/custom_error"
	entities "main/src/entities/device"
	user "main/src/entities/user"
	deviceRepository "main/src/repository/device_repository"
	userRepository "main/src/repository/user_repository"
	utils "main/src/utils"

	"golang.org/x/text/language"
)

type UserCreateHandler struct {
	Language language.Tag
}

func New(lang language.Tag) UserCreateHandler {
	return UserCreateHandler{Language: lang}
}

func (uh *UserCreateHandler) CreateUser(u RequestRegisterUser, userAgent string) (_ ResponseCreateUser, err *custom_error.CustomError) {
	a := userRepository.New(uh.Language)
	dr := deviceRepository.New(uh.Language)
	u.Password = utils.Md5(u.Password)
	var s bool
	s, err = a.Create(user.User{Name: u.Name, Email: u.Email, Password: u.Password})
	if err != nil {
		return ResponseCreateUser{Status: false}, err
	}
	var userByEmail user.User
	userByEmail, err = a.FindByEmail(u.Email)
	if err != nil {
		return ResponseCreateUser{Status: false}, err
	}

	s, err = dr.Create(entities.Device{UserAgent: userAgent, UserId: userByEmail.Id})

	return ResponseCreateUser{Status: s}, err
}
