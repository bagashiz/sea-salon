package user

import "context"

type Service struct {
	repo ReadWriter
}

func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAccount(ctx context.Context, fullName, phoneNumber, email, password, confirmPassword string) (*User, error) {
	var input struct {
		fullName    string
		email       Email
		password    Password
		phoneNumber PhoneNumber
	}
	{
		var err error

		input.fullName = fullName
		if input.fullName == "" {
			return nil, ErrFullNameEmpty
		}
		input.phoneNumber, err = NewPhone(phoneNumber)
		if err != nil {
			return nil, err
		}
		input.email, err = NewEmail(email)
		if err != nil {
			return nil, err
		}
		input.password, err = NewPassword(password)
		if err != nil {
			return nil, err
		}
		if password != confirmPassword {
			return nil, ErrPasswordMismatch
		}
	}

	user, err := NewUser(
		input.fullName,
		input.phoneNumber,
		input.email,
		input.password,
		"customer",
	)
	if err != nil {
		return nil, err
	}
	user.Create()

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
