package valuetype

import "mime/multipart"

type CreateVarietyIn struct {
	Name  string
	Price float64
	Stock int
	Image string
}

type UploadImageIn struct {
	ID   int
	File *multipart.FileHeader
}

type UpdateVarietyIn struct {
	ID    uint64
	Name  string
	Price float64
	Stock int
}
