package users

import (
	"errors"
	"github.com/htchan/UserService/internal/utils"
)

func (user User) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into users (username, password) values (?, ?)",
		user.Username, user.encryptedPassword)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (user User) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from users where username=?", user.Username)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (user User) update() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("update users set password=? where username=?",
		user.Username, user.encryptedPassword)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindUserByName(username string) (*User, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select password from users where username=?",
		username)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		user := new(User)
		user.Username = username
		rows.Scan(&user.encryptedPassword)
		return user, nil
	}
	return nil, errors.New("invalid username / password")
}