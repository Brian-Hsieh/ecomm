package pkg

type UserPayload struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
