package pkg

import "time"

type UserPayload struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	DOB       string    `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
}
