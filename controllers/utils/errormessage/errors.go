package errormessage

// MessageError struct to mock error message response
type MessageError struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
