package status

type Status struct {
	Label string `json:"label"`
	Id    int8   `json:"id"`
}

func New(name string) Status {
	return Status{name, 0} // enforce the default value here
}
