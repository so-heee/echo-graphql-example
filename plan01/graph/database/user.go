package database

import (
	"context"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID   string `gorm:"column:travel_user_id;primary_key"`
	Name string `gorm:"column:user_name"`
}

func (u *User) TableName() string {
	return "travel_user"
}

type UserDao interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindOne(ctx context.Context, id string) (*User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) FindAll(ctx context.Context) ([]*User, error) {
	var users []*User
	res := d.db.Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *userDao) FindOne(ctx context.Context, id string) (*User, error) {
	var users []*User
	res := d.db.Where("travel_user_id = ?", id).Find(&users)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(users) < 1 {
		return nil, nil
	}
	return users[0], nil
}
