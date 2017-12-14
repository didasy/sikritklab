package response

type Response struct {
	Error   string      `json:"error,omitempty"`
	Message interface{} `json:"message"`
}

type M map[string]interface{}
