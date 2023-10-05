package errors

type WebErrorResponse struct {
	Description string `json:"description"`
}

var WebErrorResponseInternalServerError = WebErrorResponse{Description: "Internal Server Error"}
