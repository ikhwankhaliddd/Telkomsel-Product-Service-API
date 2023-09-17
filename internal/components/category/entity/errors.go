package entity

type categoryError string

const (
	ErrMockCategory categoryError = "error mock category"
)

func (e categoryError) Error() string {
	return string(e)
}
