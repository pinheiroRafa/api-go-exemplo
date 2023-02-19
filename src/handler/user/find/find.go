package user_handler

import (
	custom_error "main/src/entities/custom_error"
	user "main/src/entities/user"
	userToken "main/src/entities/user"
	userRepository "main/src/repository/user_repository"

	"golang.org/x/text/language"
)

type UserFindHandler struct {
	Language language.Tag
}

func New(lang language.Tag) UserFindHandler {
	return UserFindHandler{Language: lang}
}

func (uh *UserFindHandler) FindUser(u *userToken.UserToken) (_ any, err *custom_error.CustomError) {
	a := userRepository.New(uh.Language)
	var user user.User
	user, err = a.FindByEmail(u.Email)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return ResponseFindUser{Data: user}, nil
}
