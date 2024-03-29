package userRepositorysModules

import (
	"database/sql"
	"errors"
	"fmt"
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
	var password string
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Account, &password, &user.Nickname, &user.Avatar, &user.CreateTime, &user.UpdateTime, &user.DeleteTime, &user.Status)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

func (r *userRepository) GetUsers(page responsies.Page) ([]responsies.UsersSingleResponse, error) {
	rows, err := db.DB.Query("SELECT * FROM users LIMIT ?, ?", page.Offset, page.Limit)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Warn(err)
		}
	}(rows)

	users := make([]responsies.UsersSingleResponse, 10)
	var password string
	for rows.Next() {
		user := &responsies.UsersSingleResponse{}
		err = rows.Scan(&user.ID, &user.Account, &password, &user.Nickname, &user.Avatar, &user.CreateTime, &user.UpdateTime, &user.DeleteTime, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
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

func (r *userRepository) SelectUsersByAccount(account string) bool {
	exec, err := db.DB.Query("SELECT count(account) FROM users WHERE account = ?", account)
	if err != nil {
		return false
	}

	defer func(exec *sql.Rows) {
		err := exec.Close()
		if err != nil {
			utils.Log.Warn(err)
		}
	}(exec)

	var count int

	for exec.Next() {
		err = exec.Scan(&count)
		if err != nil {
			return false
		}
	}
	return count > 0
}

func (r *userRepository) FindUserByParams(params *responsies.QueryUsersParams) ([]responsies.UsersSingleResponse, error) {
	query := `
		SELECT 
			id, account, nickname, status, create_time, update_time, delete_time, status 
		FROM 
			users 
		WHERE 
			true
			# 过滤掉已被逻辑删除的用户
			AND delete_time IS NULL
		`

	var queryParams []interface{}
	if params.ID != "" {
		query += " AND id LIKE ?"
		queryParams = append(queryParams, params.ID)
	}
	if params.Nickname != "" {
		query += " AND nickname LIKE ?"
		queryParams = append(queryParams, "%"+params.Nickname+"%")
	}
	if params.Status != "" {
		query += " AND status = ?"
		queryParams = append(queryParams, params.Status)
	}
	if params.StartTime != "" {
		query += " AND create_time >= ?"
		queryParams = append(queryParams, params.StartTime)
	}
	if params.EndTime != "" {
		query += " AND create_time <= ?"
		queryParams = append(queryParams, params.EndTime)
	}

	if params.Account != "" {
		query += " AND account LIKE ?"
		queryParams = append(queryParams, "%"+params.Account+"%")
	}

	if params.Limit != nil {
		query += " LIMIT ?"
		queryParams = append(queryParams, params.Limit)
	}

	if params.Offset != nil {
		query += " OFFSET ?"
		queryParams = append(queryParams, params.Offset)
	}
	// 准备查询语句
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	rows, err := stmt.Query(queryParams...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// 解析查询结果
	var users []responsies.UsersSingleResponse
	for rows.Next() {
		var user responsies.UsersSingleResponse
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Status, &user.CreateTime, &user.UpdateTime, &user.DeleteTime, &user.Status)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindUserByAccount(account string) (*responsies.SingleUserResponseHasPassword, bool) {
	if account == "" {
		return nil, false
	}
	user := &responsies.SingleUserResponseHasPassword{}
	query := "SELECT id, account, nickname, password, status, create_time, update_time, delete_time, status FROM users WHERE account = ?"
	rows, err := db.DB.Query(query, account)
	if err != nil {
		return nil, false
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Password, &user.Status, &user.CreateTime, &user.UpdateTime, &user.DeleteTime, &user.Status)
		if err != nil {
			return nil, false
		}
		return user, true
	}
	return nil, false
}
func (r *userRepository) FindUserById(id string) (*responsies.UsersSingleResponse, bool) {
	user, err := r.GetUser(id)
	if err != nil {
		return nil, false
	}
	// 用户不存在
	if user == nil {
		return nil, false
	}
	return user, true
}

func (r *userRepository) DeleteUser(id string) error {
	stmt, err := db.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUser(params *responsies.UpdateUserRequest, id string) error {

	query := "UPDATE cms.users SET "
	var (
		queryParams []any
		hasSet      bool
	)

	if params.Nickname != "" {
		queryParams = append(queryParams, params.Nickname)
		query += "nickname = ?, "
		hasSet = true
	}

	if params.Password != "" {
		// 密码加密
		password, err := utils.Bcrypt.HashPassword(params.Password)
		if err != nil {
			return err
		}
		queryParams = append(queryParams, password)
		query += "password = ?, "
		hasSet = true
	}

	if !hasSet {
		// 如果没有任何条件需要更新
		return nil
	}

	// 去除最后一个逗号和空格
	query = query[:len(query)-2]

	query += ` WHERE id = ?`
	queryParams = append(queryParams, id)
	fmt.Println(query)

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(queryParams...)
	if err != nil {
		return err
	}

	return nil

}
