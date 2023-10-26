package request

type (
	CreateUser struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}
)
