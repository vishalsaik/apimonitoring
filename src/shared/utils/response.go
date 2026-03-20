package utils

import "time"

type APIResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	StatusCode int         `json:"status_code"`
	Timestamp  string      `json:"timestamp"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type PaginatedResponse struct {
	Success    bool           `json:"success"`
	Data       interface{}    `json:"data,omitempty"`
	Pagination PaginationMeta `json:"pagination"`
	Timestamp  string         `json:"timestamp"`
}

func Success(data interface{}, message string, statusCode int) APIResponse {
	return APIResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		StatusCode: statusCode,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}

func ErrorResponse(message string, statusCode int, err interface{}) APIResponse {
	return APIResponse{
		Success:    false,
		Message:    message,
		Error:      err,
		StatusCode: statusCode,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}

func ValidationError(err interface{}) APIResponse {
	return APIResponse{
		Success:    false,
		Message:    "Validation failed",
		Error:      err,
		StatusCode: 400,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}
}

func Paginated(data interface{}, page, limit, total int) PaginatedResponse {
	totalPages := 0
	if limit > 0 {
		totalPages = (total + limit - 1) / limit
	}
	return PaginatedResponse{
		Success: true,
		Data:    data,
		Pagination: PaginationMeta{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}
