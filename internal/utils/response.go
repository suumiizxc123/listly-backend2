package utils

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) Response {

	return Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func Error(message string, errors interface{}) Response {

	return Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	}

}
