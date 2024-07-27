package repository

import (
	"context"

	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *PostgresRepository) CreateUser(ctx context.Context, u *user.User) error {
	arg := postgres.InsertUserParams{
		ID:          u.ID,
		Email:       u.Email,
		Password:    u.Password,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
		Role:        postgres.UserRole(u.Role),
		CreatedAt:   pgtype.Timestamptz{Time: u.CreatedAt, Valid: !u.CreatedAt.IsZero()},
		UpdatedAt:   pgtype.Timestamptz{Time: u.UpdatedAt, Valid: !u.UpdatedAt.IsZero()},
	}

	if err := r.db.InsertUser(ctx, arg); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsDataException(pgErr.Code) {
				return user.ErrUserInvalid
			}
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return user.ErrUserExists
			}
		}
		return err
	}

	return nil
}

func (r *PostgresRepository) GetUserByID(ctx context.Context, id string) (*user.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, user.ErrIDInvalid
	}

	result, err := r.db.SelectUserByID(ctx, uid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}

	user := &user.User{
		ID:          result.ID,
		FullName:    result.FullName,
		Email:       result.Email,
		Password:    result.Password,
		PhoneNumber: result.PhoneNumber,
		Role:        string(result.Role),
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}

	return user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	result, err := r.db.SelectUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}

	user := &user.User{
		ID:          result.ID,
		FullName:    result.FullName,
		Email:       result.Email,
		Password:    result.Password,
		PhoneNumber: result.PhoneNumber,
		Role:        string(result.Role),
		CreatedAt:   result.CreatedAt.Time,
		UpdatedAt:   result.UpdatedAt.Time,
	}

	return user, nil
}

func (r *PostgresRepository) ListUsers(ctx context.Context, limit, offset int) ([]*user.User, error) {
	arg := postgres.SelectAllUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	result, err := r.db.SelectAllUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	var users []*user.User

	for _, u := range result {
		users = append(users, &user.User{
			ID:          u.ID,
			FullName:    u.FullName,
			Email:       u.Email,
			Password:    u.Password,
			PhoneNumber: u.PhoneNumber,
			Role:        string(u.Role),
			CreatedAt:   u.CreatedAt.Time,
			UpdatedAt:   u.UpdatedAt.Time,
		})
	}

	return users, nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, u *user.User) error {
	arg := postgres.UpdateUserParams{
		ID:          u.ID,
		Email:       pgtype.Text{String: u.Email, Valid: u.Email != ""},
		Password:    pgtype.Text{String: u.Password, Valid: u.Password != ""},
		FullName:    pgtype.Text{String: u.FullName, Valid: u.FullName != ""},
		PhoneNumber: pgtype.Text{String: u.PhoneNumber, Valid: u.PhoneNumber != ""},
		Role:        postgres.NullUserRole{UserRole: postgres.UserRole(u.Role), Valid: u.Role != ""},
	}

	result, err := r.db.UpdateUser(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsDataException(pgErr.Code) {
				return user.ErrUserInvalid
			}
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return user.ErrUserExists
			}
		}
		if err == pgx.ErrNoRows {
			return user.ErrUserNotFound
		}
		return err
	}

	u.ID = result.ID
	u.FullName = result.FullName
	u.Email = result.Email
	u.Password = result.Password
	u.PhoneNumber = result.PhoneNumber
	u.Role = string(result.Role)
	u.CreatedAt = result.CreatedAt.Time
	u.UpdatedAt = result.UpdatedAt.Time

	return nil
}

func (r *PostgresRepository) DeleteUser(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return user.ErrIDInvalid
	}

	_, err = r.db.DeleteUser(ctx, uid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user.ErrUserNotFound
		}
		return err
	}

	return nil
}
