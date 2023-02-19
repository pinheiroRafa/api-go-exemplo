package repository

import (
	custom_error "main/src/entities/custom_error"
	device "main/src/entities/device"
	db "main/src/repository"
	"os/user"

	"golang.org/x/text/language"
)

type DeviceRepository struct {
	TableName string
	Language  language.Tag
}

func New(lang language.Tag) DeviceRepository {
	return DeviceRepository{TableName: "EST_DEVICES", Language: lang}
}

func (p *DeviceRepository) Create(s device.Device) (_ bool, err *custom_error.CustomError) {
	return db.Exec[user.User]("INSERT INTO "+p.TableName+" (user_agent, user_id) VALUES ($1, $2)",
		s.UserAgent,
		s.UserId,
	)
}
func (p *DeviceRepository) Update(s device.Device) (_ bool, err *custom_error.CustomError) {
	return db.Exec[user.User]("UPDATE "+p.TableName+" SET last_used = $1 WHERE id = $2",
		s.LastUsed,
		s.Id,
	)
}
func (p *DeviceRepository) FindByUserId(userId string, userAgent string) (_ *device.Device, err *custom_error.CustomError) {
	var list []device.Device
	list, err = db.Select[device.Device](
		"SELECT  id,user_agent,created_at,last_used,user_id FROM "+p.TableName+" WHERE user_id = $1 and user_agent = $2",
		userId,
		userAgent,
	)

	if list != nil && len(list) > 0 {
		return &list[0], err
	}

	return nil, err
}
