package exception

type BadRequestError struct {
	Message string
}

func (b *BadRequestError) Error() string {
	return b.Message
}

type RecordNotFoundError struct {
	Message string
}

func (b *RecordNotFoundError) Error() string {
	return b.Message
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
