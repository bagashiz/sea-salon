package user

import "context"

// GetAccountByEmail validates and retrieves a user account by provided email and password.
func (s *Service) GetAccountByEmail(ctx context.Context, email, password string) (*Account, error) {
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

	account, err := s.repo.GetAccountByEmail(ctx, input.email.String())
	if err != nil {
		return nil, err
	}

	if ok := input.password.Compare(account.Password); !ok {
		return nil, ErrPasswordMismatch
	}

	return account, nil
}
