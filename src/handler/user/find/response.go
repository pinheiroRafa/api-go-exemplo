package user_handler

import userToken "main/src/entities/user"

type ResponseFindUser struct {
	Data userToken.User `json:"data"`
}
