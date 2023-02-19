package statusHandler

import (
	custom_error "main/src/entities/custom_error"
	status "main/src/entities/status"
	statusRepository "main/src/repository/status_repository"
	"main/src/utils"
)

func ListAllStatus() (_ ResponseListStatus, err *custom_error.CustomError) {
	a := statusRepository.New()
	var list []status.Status
	list, err = a.FindAll()
	return ResponseListStatus{Data: list}, err
}

func ListStatusById(id string) (_ []status.Status, err *custom_error.CustomError) {
	a := statusRepository.New()
	var val int
	val, err = utils.ParseInt(id)
	if err != nil {
		return nil, err
	}
	return a.FindById(val)
}
