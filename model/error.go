package model

type GetOurValuesContextError struct {
	error
}

func NewGetOurVaulesContextError(err error) error {
	return GetOurValuesContextError{err}
}

type GetError struct {
	error
}

func NewGetError(err error) error {
	return GetError{err}
}

type PutError struct {
	error
}

func NewPutError(err error) error {
	return PutError{err}
}

type DeleteError struct {
	error
}

func NewDeleteError(err error) error {
	return DeleteError{err}
}
