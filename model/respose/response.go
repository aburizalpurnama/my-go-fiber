package respose

type (
	Default struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		Data    any    `json:"data,omitempty"`
	}

	UserData struct {
		Id string `json:"id,omitempty"`
	}
)
