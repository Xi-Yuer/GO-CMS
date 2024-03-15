package userRepositorysModules

import (
	"database/sql"
	"errors"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/responsies"
	"github.com/Xi-Yuer/cms/utils"
)

var UserRepository = &userRepository{}

type userRepository struct{}

func (r *userRepository) GetUser(id string) (*responsies.UsersSingleResponse, error) {
	rows, err := db.DB.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Warn(err)
		}
	}(rows)

	user := &responsies.UsersSingleResponse{}
	found := false
	var password string
	for rows.Next() {
		found = true
		err = rows.Scan(&user.ID, &user.Account, &password, &user.Nickname, &user.Avatar, &user.Gender, &user.Role, &user.CreateTime, &user.UpdateTime, &user.DeleteTime, &user.Status)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *responsies.CreateSingleUserRequest) error {
	id := utils.GenID()
	exec, err := db.DB.Exec("INSERT INTO users (id,account,nickname,password) VALUES (?, ?, ?, ?)", id, user.Account, user.Nickname, user.Password)
	if err != nil {
		return err
	}

	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("插入数据失败")
	}
	return nil
}
