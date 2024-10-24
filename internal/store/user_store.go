package store

import (
	"app/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	sq "github.com/Masterminds/squirrel"
)

func (s *Store) UserExists(email string) (bool, error) {
	query, args, err := sq.
		Select(fmt.Sprintf("EXISTS (SELECT 1 FROM \"%s\" WHERE \"%s\" = ?)", tableNameUsers, fieldEmail)).
		From(tableNameUsers).
		Where(sq.Eq{fieldEmail: email}).
		ToSql()
	if err != nil {
		return false, err
	}

	var exists bool
	err = s.db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *Store) CreateUser(user *models.SignUpParams) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query, args, err := sq.
		Insert(tableNameUsers).
		Columns(fieldEmail, fieldPassword).
		Values(user.Email, string(hashedPassword)).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, args...)
	return err
}
