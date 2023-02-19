package repository

import (
	custom_error "main/src/entities/custom_error"
	user "main/src/entities/user"
	db "main/src/repository"
	"main/src/utils"
	"strings"

	"golang.org/x/text/language"
)

type UserRepository struct {
	TableName string
	Language  language.Tag
}

func New(lang language.Tag) UserRepository {
	return UserRepository{TableName: "EST_USERS", Language: lang}
}

func (p *UserRepository) FindByEmail(email string) (_ user.User, err *custom_error.CustomError) {
	var success, er = db.Select[user.User]("SELECT id, name,email, status, password, created_at, updated_at FROM "+p.TableName+" WHERE email = $1", email)
	if er != nil {
		return user.User{}, er
	}
	return success[0], nil
}

func (p *UserRepository) Create(s user.User) (_ bool, err *custom_error.CustomError) {
	var success, er = db.Exec[user.User]("INSERT INTO "+p.TableName+" (name, email, password, status) VALUES ($1, $2, $3, $4)",
		s.Name,
		s.Email,
		s.Password,
		2,
	)
	if er != nil {
		e := er.Message
		if strings.Index(e, "est_users_email_key") >= 0 {
			err := custom_error.New(utils.GetString("multiplesEmails", p.Language))
			return false, &err
		}
		return false, er
	}
	return success, nil
}
