package entity

type productsError string

const (
	ErrMockProducts productsError = "error mock products"
)

func (e productsError) Error() string {
	return string(e)
}
