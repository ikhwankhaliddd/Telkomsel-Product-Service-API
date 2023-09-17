package entity

type varietyError string

const (
	ErrMockVariety varietyError = "error mock variety"
)

func (e varietyError) Error() string {
	return string(e)
}
