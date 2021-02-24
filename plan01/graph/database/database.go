package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Connect(dbUser string, dbPwd string, instanceConnectionName string) (*gorm.DB, error) {
	conn, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s", dbUser, dbPwd, instanceConnectionName))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
