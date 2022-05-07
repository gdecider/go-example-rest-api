package store

import "go-rest-api/internal/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	stmt, err := r.store.db.Prepare("insert into users(email, encrypted_password) values(?, ?);")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(u.Email, u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	u.ID, err = res.LastInsertId()
	if err != nil {
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
		return nil, err
	}

	return u, nil
}
