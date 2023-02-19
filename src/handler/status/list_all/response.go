package statusHandler

import status "main/src/entities/status"

type ResponseListStatus struct {
	Data []status.Status `json:"data"`
}
