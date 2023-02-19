package statusHandler

import (
	"fmt"
	custom_error "main/src/entities/custom_error"
	status "main/src/entities/status"
	statusRepository "main/src/repository/status_repository"
	"main/src/utils"
)

func UpdateStatus(s status.Status, id string) (_ ResponseStatus, err *custom_error.CustomError) {
	a := statusRepository.New()
	var val int
	val, err = utils.ParseInt(id)
	if err != nil {
		return ResponseStatus{Status: false}, err
	}
	fmt.Sscan(fmt.Sprintf("%d", val), &s.Id)
	var success bool
	success, err = a.Update(s)
	return ResponseStatus{Status: success}, err
}

func DeleteStatus(id string) (_ ResponseStatus, err *custom_error.CustomError) {
	a := statusRepository.New()
	var val int
	val, err = utils.ParseInt(id)
	if err != nil {
		return ResponseStatus{Status: false}, err
	}
	var success bool
	success, err = a.Delete(val)
	return ResponseStatus{Status: success}, err
}

func CreateStatus(s status.Status) (_ ResponseStatus, err *custom_error.CustomError) {
	a := statusRepository.New()
	var success bool
	success, err = a.Create(s)
	return ResponseStatus{Status: success}, err
}
