package helper

type BaseHttpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
	Error   any    `json:"error"`
}

func GenerateBaseResponse(success bool, message string, result any) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Message: message,
		Result:  result,
	}
}

func GenerateBaseResponseWithError(success bool, message string, result any, err interface{}) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success: success,
		Message: message,
		Result:  result,
		Error:   err,
	}
}
