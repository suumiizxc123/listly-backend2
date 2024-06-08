package utils

type Message struct {
	English   string `json:"english"`
	Mongolian string `json:"mongolian"`
	Chinese   string `json:"chinese"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message Message     `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func Success(message []string, data interface{}) Response {

	return Response{
		Status: true,
		Message: Message{
			English:   message[0],
			Mongolian: message[1],
		},
		Errors: nil,
		Data:   data,
	}
}

func Error(message []string, errors interface{}) Response {

	return Response{
		Status: false,
		Message: Message{
			English:   message[0],
			Mongolian: message[1],
		},
		Errors: errors,
		Data:   nil,
	}

}
