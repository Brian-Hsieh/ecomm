package user

import (
	"database/sql"
	"fmt"

	"github.com/Brian-Hsieh/ecomm/pkg"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*pkg.User, error) {
	user := s.db.QueryRow("SELECT * FROM ecomm WHERE email = ?", email)

	u := new(pkg.User)
	if err := user.Scan(&u.ID, &u.Name, &u.Password, &u.Email, &u.DOB, &u.CreatedAt); err != nil {
		fmt.Println("User not found")
		return nil, err
	}

	return u, nil
}
