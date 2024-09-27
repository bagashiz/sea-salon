package postgres

import "github.com/bagashiz/sea-salon/internal/app/user"

// ToDomain converts generated Account struct to domain account struct
func (a Account) ToDomain() *user.Account {
	return &user.Account{
		ID:          a.ID,
		FullName:    a.FullName,
		Email:       a.Email,
		Password:    a.Password,
		PhoneNumber: a.PhoneNumber,
		Role:        string(a.Role),
		CreatedAt:   a.CreatedAt.Time,
		UpdatedAt:   a.UpdatedAt.Time,
	}
}
