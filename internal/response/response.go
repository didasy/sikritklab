package response

type Response struct {
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
}

type M map[string]interface{}
