package user

import "context"

// CreateAccount validates, creates, and stores a new user account.
func (s *Service) CreateAccount(
	ctx context.Context,
	fullName, phoneNumber, email, password, confirmPassword string,
) (*Account, error) {
	var input struct {
		fullName    FullName
		email       Email
		password    Password
		phoneNumber PhoneNumber
	}
	{
		var err error

		input.fullName, err = NewFullName(fullName)
		if err != nil {
			return nil, err
		}
		input.phoneNumber, err = NewPhoneNumber(phoneNumber)
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

	account := NewAccount(
		input.fullName,
		input.phoneNumber,
		input.email,
		hashedPassword,
		"customer",
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.AddAccount(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}
