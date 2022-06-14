package sqlstore

import (
	"database/sql"
	"go-rest-api/internal/model"
	"go-rest-api/internal/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	stmt, err := r.store.db.Prepare("insert into users(email, encrypted_password) values(?, ?);")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, u.EncryptedPassword)
	if err != nil {
		return err
	}

	u.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Find(id int64) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where id = ?",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where email = ?",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
