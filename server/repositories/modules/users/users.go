package userRepositorysModules

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
	"time"
)

var UserRepository = &userRepository{}

type userRepository struct{}

func (r *userRepository) GetUser(id string) (*dto.UsersSingleResponse, error) {
	rows, err := db.DB.Query(`
	SELECT u.id, u.account, u.nickname, GROUP_CONCAT(ur.role_id), u.avatar,u.status, u.create_time, u.update_time
	FROM users u
	LEFT JOIN users_roles ur ON u.id = ur.user_id
	WHERE u.id = ? AND u.delete_time IS NULL
	GROUP BY u.id, u.account, u.nickname, u.avatar, u.create_time, u.update_time, u.status
	`, id)

	if err != nil {
		fmt.Println('1', err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)

	user := &dto.UsersSingleResponse{}
	var rolesID []uint8

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &rolesID, &user.Avatar, &user.Status, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == "" {
		return nil, errors.New("用户不存在")
	}
	if rolesID != nil {
		user.RolesID = strings.Split(string(rolesID), ",")
	}
	return user, nil
}

func (r *userRepository) GetUsers(page dto.Page) ([]dto.UsersSingleResponse, error) {
	rows, err := db.DB.Query(`
	SELECT u.id, u.account, u.nickname, GROUP_CONCAT(ur.role_id), u.avatar,u.status, u.create_time, u.update_time
	FROM users u
	LEFT JOIN users_roles ur ON u.id = ur.user_id
	WHERE u.delete_time IS NULL
	GROUP BY u.id, u.account, u.nickname, u.avatar, u.create_time, u.update_time, u.status
	LIMIT ?, ?
`, page.Offset, page.Limit)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Warn(err)
		}
	}(rows)

	users := make([]dto.UsersSingleResponse, 0)
	for rows.Next() {
		user := &dto.UsersSingleResponse{}
		var rolesID []uint8
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &rolesID, &user.Avatar, &user.Status, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			continue
		}
		if rolesID != nil {
			user.RolesID = strings.Split(string(rolesID), ",")
		}
		users = append(users, *user)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(user *dto.CreateSingleUserRequest) error {
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
	exec, err := db.DB.Query("SELECT count(account) FROM users WHERE account = ? AND delete_time IS NULL", account)
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

func (r *userRepository) FindUserByParams(params *dto.QueryUsersParams) ([]dto.UsersSingleResponse, error) {
	query := `
		SELECT 
			id, account, nickname, status, create_time, update_time, status 
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
	var users []dto.UsersSingleResponse
	for rows.Next() {
		var user dto.UsersSingleResponse
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Status, &user.CreateTime, &user.UpdateTime, &user.Status)
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

func (r *userRepository) FindUserByAccount(account string) (*dto.SingleUserResponseHasPassword, bool) {
	if account == "" {
		return nil, false
	}
	user := &dto.SingleUserResponseHasPassword{}
	query := "SELECT id, account, nickname, password, status, create_time, update_time, status FROM users WHERE account = ? AND delete_time IS NULL"
	rows, err := db.DB.Query(query, account)
	if err != nil {
		return nil, false
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Password, &user.Status, &user.CreateTime, &user.UpdateTime, &user.Status)
		if err != nil {
			return nil, false
		}
		return user, true
	}
	return nil, false
}

func (r *userRepository) FindUserById(id string) (*dto.UsersSingleResponse, bool) {
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
	now := time.Now()
	stmt, err := db.DB.Prepare(`UPDATE users SET delete_time = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	_, err = stmt.Exec(now, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateUser(params *dto.UpdateUserRequest, id string) error {
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
