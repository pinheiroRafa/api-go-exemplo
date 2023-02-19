package user_handler

import (
	"fmt"
	custom_error "main/src/entities/custom_error"
	entities "main/src/entities/device"
	user "main/src/entities/user"
	deviceRepository "main/src/repository/device_repository"
	userRepository "main/src/repository/user_repository"
	"main/src/utils"
	"time"

	"golang.org/x/text/language"
)

type UserLoginHandler struct {
	Language language.Tag
}

func New(lang language.Tag) UserLoginHandler {
	return UserLoginHandler{Language: lang}
}

func (uh *UserLoginHandler) FindUserByAuth(u RequestLoginUser, userAgent string) (_ any, err *custom_error.CustomError) {
	a := userRepository.New(uh.Language)
	u.Password = utils.Md5(u.Password)
	var user user.User
	user, err = a.FindByEmail(u.Email)
	if err != nil {
		return user, err
	}
	if user.Password != u.Password {
		err := custom_error.New(utils.GetString("wrongPassowrd", uh.Language))
		return user, &err
	}
	if user.Status == 3 {
		err := custom_error.New(utils.GetString("blockedUser", uh.Language))
		return user, &err
	}
	dr := deviceRepository.New(uh.Language)
	var device *entities.Device
	device, err = dr.FindByUserId(user.Id, userAgent)
	if err != nil {
		return user, err
	}
	if device == nil {
		var err = custom_error.NewCode(utils.GetString("suspectedDevice", uh.Language), "SUSPECTED_DEVICE")
		return user, &err
	}
	var time = time.Now()
	device.LastUsed = &time
	_, err = dr.Update(*device)
	var token string
	token, err = utils.GenerateJWT(map[string]string{"user": user.Id, "status": fmt.Sprint(user.Status), "email": user.Email})
	return ResponseLoginUser{Data: token}, err
}
