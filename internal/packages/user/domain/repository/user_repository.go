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
			name,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4
		);
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id.GetValue(), user.Name, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserById(id *valueobject.Id) (*entity.User, error) {
	query := `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id.GetValue())

	if err != nil {
		return nil, err
	}

	var user entity.User

	for rows.Next() {
		var userInput entity.UserRestoreInput
		err := rows.Scan(
			&userInput.Id,
			&userInput.Name,
			&userInput.CreatedAt,
			&userInput.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		user.Restore(userInput)
	}

	return &user, nil
}

func (u *UserRepository) GetUsers() ([]entity.User, error) {
	query := `
		SELECT *
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
			&userInput.Name,
			&userInput.CreatedAt,
			&userInput.UpdatedAt,
		)

		var user entity.User
		user.Restore(userInput)

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
			name = $2,
			updated_at = $3
		WHERE id = $1;
	`

	stmt, err := u.db.Client.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id.GetValue(), user.Name, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
