package response

type Response struct {
	Data   interface{}    `json:"data"`
	Status ResponseStatus `json:"status"`
}

type ResponseStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
