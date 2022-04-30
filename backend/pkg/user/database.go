package user

import (
	"github.com/htchan/UserService/backend/internal/utils"
)

func execute(operation, model string, sql string, params ...interface{}) error {
	tx, err := utils.GetDB().Begin()
	if err != nil { return DatabaseError{operation, model, err} }
	commited := false
	defer func() { if !commited { tx.Rollback() } }()

	_, err = tx.Exec(sql, params...)
	if err != nil {
		return DatabaseError{operation, model, err}
	}
	err = tx.Commit()
	if err != nil { return DatabaseError{operation, model, err} }
	commited = true
	return nil
}

func (user User) create() error {
	return execute("create", "user", "insert into users (uuid, username, password) values (?, ?, ?)",
		user.UUID, user.Username, user.EncryptedPassword)
}

func (user User) delete() error {
	return execute("delete", "user", "delete from users where username=?", user.Username)
}

func (user User) update() error {
	return execute("update", "user", "update users set password=? where uuid=?",
		user.Username, user.EncryptedPassword)
}

func FindUserByName(username string) (User, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return User{}, DatabaseError{"select", "user", err}
	}
	defer tx.Rollback()
	rows, err := tx.Query("select uuid, password from users where username=?",
		username)
	if err != nil {
		return User{}, DatabaseError{"select", "user", err}
	}
	if rows.Next() {
		user := User{ Username: username }
		rows.Scan(&user.UUID, &user.EncryptedPassword)
		return user, nil
	}
	return User{}, DatabaseError{"select", "user", IncorrectParamsError("username or password")}
}

func FindUserByUUID(uuid string) (User, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return User{}, DatabaseError{"select", "user", err}
	}
	defer tx.Rollback()
	rows, err := tx.Query("select username, password from users where uuid=?",
		uuid)
	if err != nil {
		return User{}, DatabaseError{"select", "user", err}
	}
	if rows.Next() {
		user := User{ UUID: uuid }
		rows.Scan(&user.Username, &user.EncryptedPassword)
		return user, nil
	}
	return User{}, DatabaseError{"select", "user", IncorrectParamsError("username or password")}
}