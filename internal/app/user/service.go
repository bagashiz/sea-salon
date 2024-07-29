package user

import "context"

// Service provides user business logic and operations.
type Service struct {
	repo ReadWriter
}

// NewService creates a new user service instance.
func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}

// CreateAccount validates, creates, and stores a new user account.
func (s *Service) CreateAccount(
	ctx context.Context,
	fullName, phoneNumber, email, password, confirmPassword string,
) (*User, error) {
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

	hashedPassword, err := input.password.Hash()
	if err != nil {
		return nil, err
	}

	user, err := NewUser(
		input.fullName,
		input.phoneNumber,
		input.email,
		hashedPassword,
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

func (s *Service) GetUserByEmail(ctx context.Context, email, password string) (*User, error) {
	var input struct {
		email    Email
		password Password
	}
	{
		var err error

		input.email, err = NewEmail(email)
		if err != nil {
			return nil, err
		}
		input.password, err = NewPassword(password)
		if err != nil {
			return nil, err
		}
	}

	user, err := s.repo.GetUserByEmail(ctx, input.email.String())
	if err != nil {
		return nil, err
	}

	if ok := input.password.Compare(user.Password); !ok {
		return nil, ErrPasswordMismatch
	}

	return user, nil
}
