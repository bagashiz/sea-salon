package postgres

import "github.com/bagashiz/sea-salon/internal/app/user"

// ToDomain converts generated User struct to domain user struct
func (u User) ToDomain() *user.User {
	return &user.User{
		ID:          u.ID,
		FullName:    u.FullName,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
		Role:        string(u.Role),
		CreatedAt:   u.CreatedAt.Time,
		UpdatedAt:   u.UpdatedAt.Time,
	}
}
