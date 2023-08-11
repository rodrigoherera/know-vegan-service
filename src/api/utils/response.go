package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePagination struct {
	Success  bool        `json:"success"`
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}

type Metadata struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}
