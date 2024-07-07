package exception

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}