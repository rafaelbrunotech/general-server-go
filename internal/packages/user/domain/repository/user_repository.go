package repository

import (
	valueobject "github.com/rafaelbrunoss/general-server-go/internal/common/domain/value-object"
	db "github.com/rafaelbrunoss/general-server-go/internal/common/infrastructure/database"
	"github.com/rafaelbrunoss/general-server-go/internal/packages/user/domain/entity"
)

type UserRepository struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(user *entity.User) error {
	query := `
		INSERT INTO users (
			id,
			email,
			name,
			password,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6
		);
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Id.Value(),
		user.Email.Value(),
		user.Name,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(email *valueobject.Email) (*entity.User, error) {
	query := `
		SELECT
			id,
			email,
			name,
			password,
			created_at,
			updated_at
		FROM users
		WHERE email = $1;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(email.Value())

	if err != nil {
		return nil, err
	}

	var user entity.User

	for rows.Next() {
		var userInput entity.UserRestoreInput
		err := rows.Scan(
			&userInput.Id,
			&userInput.Email,
			&userInput.Name,
			&userInput.Password,
			&userInput.CreatedAt,
			&userInput.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		err = user.Restore(userInput)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) GetUserById(id *valueobject.Id) (*entity.User, error) {
	query := `
		SELECT
			id,
			email,
			name,
			password,
			created_at,
			updated_at
		FROM users
		WHERE id = $1;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id.Value())

	if err != nil {
		return nil, err
	}

	var user entity.User

	for rows.Next() {
		var userInput entity.UserRestoreInput
		err := rows.Scan(
			&userInput.Id,
			&userInput.Email,
			&userInput.Name,
			&userInput.Password,
			&userInput.CreatedAt,
			&userInput.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		err = user.Restore(userInput)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserRepository) GetUsers() ([]entity.User, error) {
	query := `
		SELECT
			id,
			email,
			name,
			password,
			created_at,
			updated_at
		FROM users;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var users []entity.User

	for rows.Next() {
		var userInput entity.UserRestoreInput
		err := rows.Scan(
			&userInput.Id,
			&userInput.Email,
			&userInput.Name,
			&userInput.Password,
			&userInput.CreatedAt,
			&userInput.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		var user entity.User
		err = user.Restore(userInput)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) UpdateUser(user *entity.User) error {
	query := `
		UPDATE users
		SET
			email = $2,
			name = $3,
			password = $4,
			updated_at = $5
		WHERE id = $1;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.Id.Value(),
		user.Email.Value(),
		user.Name,
		user.Password,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
