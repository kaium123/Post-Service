package errors

type RestErr struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
type NewErr struct {
	*RestErr
	Errors error
}

func ErrorResp(status int, errorMessage string) *RestErr {
	return &RestErr{
		Status: status,
		Error:  errorMessage,
	}
}

func ErrorRespWithError(status int, errors error, errorMessage string) *NewErr {
	return &NewErr{
		Errors:  errors,
		RestErr: ErrorResp(status, errorMessage),
	}
}
