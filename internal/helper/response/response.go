package response

type Response struct {
	StatusMessage string         `json:"status_message"`
	Code          int            `json:"code"`
	Data          interface{}    `json:"data,omitempty"`
	Pagination    *PaginationObj `json:"pagination,omitempty"`
}

type PaginationObj struct {
	CurrentPage   uint64 `json:"current_page"`
	LastPage      uint64 `json:"last_page"`
	Count         uint64 `json:"count"`
	RecordPerPage uint64 `json:"record_per_page"`
}

func CreateResponse(statusMessage string, code int, data interface{}, pagination *PaginationObj) Response {
	return Response{
		StatusMessage: statusMessage,
		Code:          code,
		Data:          data,
		Pagination:    pagination,
	}
}
