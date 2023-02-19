package repository

import (
	custom_error "main/src/entities/custom_error"
	"main/src/entities/status"
	db "main/src/repository"
)

type StatusRepository struct {
	TableName string
}

func New() StatusRepository {
	return StatusRepository{TableName: "EST_STATUS"}
}

func (p *StatusRepository) FindAll() (_ []status.Status, err *custom_error.CustomError) {
	return db.Select[status.Status]("SELECT label, id FROM " + p.TableName)
}
func (p *StatusRepository) Update(s status.Status) (_ bool, err *custom_error.CustomError) {
	return db.Exec[status.Status]("UPDATE "+p.TableName+" SET label = $1 WHERE id = $2", s.Label, s.Id)
}
func (p *StatusRepository) Delete(id int) (_ bool, err *custom_error.CustomError) {
	return db.Exec[status.Status]("DELETE FROM "+p.TableName+" WHERE id = $1", id)
}
func (p *StatusRepository) Create(s status.Status) (_ bool, err *custom_error.CustomError) {
	return db.Exec[status.Status]("INSERT INTO "+p.TableName+" (label, id) VALUES ($1, $2)", s.Label, s.Id)
}
func (p *StatusRepository) FindById(id int) (_ []status.Status, err *custom_error.CustomError) {
	return db.Select[status.Status]("SELECT label, id FROM "+p.TableName+" WHERE id = $1", id)
}
