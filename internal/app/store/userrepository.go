package store

import "github.com/WorkLevchenko/willknow/internal/app/model"

type UserRepository struct {
	store *Store
}

// Принимает на входе модель, которую мы хотим создать в БД и возваращает её же
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	/*
		Метод .Scan для того, чтобы после того, как зпрос вернул строку, он эти значения
		смапил в переданные аргументы
	*/
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// Ищет пользователя по email, нужен при регистрации и составлении почты и пароля пользователя
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
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
