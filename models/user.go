package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64     `bun:"id,pk,autoincrement"`
	GUID     uuid.UUID `bun:"guid,pk,default:uuid_generate_v4()"`
	Name     string    `bun:"name,notnull"`
	Email    string    `bun:"email,notnull"`
	Password string    `bun:"password,notnull"`
}

func NewUserModel() *User {
	return &User{}
}

func (m *User) GetUsers(ctx context.Context, db *bun.DB) ([]User, error) {
	var users []User

	err := db.NewSelect().
		Model(&users).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *User) CreateUser(ctx context.Context, db *bun.DB) error {
	_, err := db.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *User) GetUserByGUID(ctx context.Context, db *bun.DB, guid uuid.UUID) (*User, error) {
	var user User

	err := db.NewSelect().
		Model(&user).
		Where("guid = ?", guid).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *User) GetUserByEmail(ctx context.Context, db *bun.DB, email string) (*User, error) {
	var user User

	err := db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
