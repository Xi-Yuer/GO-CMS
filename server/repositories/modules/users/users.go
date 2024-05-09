package userRepositorysModules

import (
	"database/sql"
	"errors"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
)

var UserRepository = &userRepository{}

type userRepository struct{}

func (u *userRepository) GetUser(id string) (*dto.UsersSingleResponse, error) {
	rows, err := db.DB.Query(`
	SELECT u.id, u.account, u.nickname, GROUP_CONCAT(ur.role_id),u.is_admin, u.avatar,u.status,u.department_id, u.create_time, u.update_time
	FROM users u
	LEFT JOIN users_roles ur ON u.id = ur.user_id
	WHERE u.id = ? AND u.delete_time IS NULL
	GROUP BY u.id
	`, id)

	if err != nil {
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
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &rolesID, &user.IsAdmin, &user.Avatar, &user.Status, &user.DepartmentID, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == "" {
		return nil, errors.New("资源不存在")
	}
	if rolesID != nil {
		user.RolesID = strings.Split(string(rolesID), ",")
	}
	return user, nil
}

func (u *userRepository) GetUsers(page dto.Page) ([]dto.UsersSingleResponse, error) {
	rows, err := db.DB.Query(`
	SELECT u.id, u.account, u.nickname, GROUP_CONCAT(ur.role_id), u.is_admin, u.avatar,u.status,u.department_id, u.create_time, u.update_time
	FROM users u
	LEFT JOIN users_roles ur ON u.id = ur.user_id
	WHERE u.delete_time IS NULL
	GROUP BY u.id, u.account, u.nickname, u.is_admin, u.avatar,u.department_id, u.create_time, u.update_time, u.status
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
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &rolesID, &user.IsAdmin, &user.Avatar, &user.Status, &user.DepartmentID, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
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

func (u *userRepository) CreateUser(user *dto.CreateSingleUserRequest) int64 {
	id := utils.GenID()
	exec, err := db.DB.Exec("INSERT INTO users (id,account,nickname,password,department_id) VALUES (?, ?, ?, ?,?)", id, user.Account, user.Nickname, user.Password, user.DepartmentID)
	if err != nil {
		return 0
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return 0
	}
	if rowsAffected == 0 {
		return 0
	}
	return id
}

func (u *userRepository) SelectUsersByAccount(account string) bool {
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

func (u *userRepository) FindUserByParams(params *dto.QueryUsersParams) (*dto.HasTotalResponseData, error) {
	count := `SELECT count(*) FROM users WHERE delete_time IS NULL`

	var total int
	rows, err := db.DB.Query(count)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()

	}(rows)
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	query := `
		SELECT 
			id, account,nickname, status,GROUP_CONCAT(ur.role_id), is_admin, department_id, status, create_time, update_time 
		FROM 
			users 
		LEFT JOIN users_roles ur ON users.id = ur.user_id
		WHERE users.delete_time IS NULL
		`
	var queryParams []interface{}
	if params.ID != "" {
		query += " AND id = ?"
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
		query += " AND update_time <= ?"
		queryParams = append(queryParams, params.EndTime)
	}

	if params.DepartmentID != "" {
		query += " AND department_id = ?"
		queryParams = append(queryParams, params.DepartmentID)
	}

	if params.Account != "" {
		query += " AND account LIKE ?"
		queryParams = append(queryParams, "%"+params.Account+"%")
	}

	query += " GROUP BY users.id"
	query += ` LIMIT ?, ?`
	if params.Offset != 0 {
		queryParams = append(queryParams, params.Offset)
	} else {
		queryParams = append(queryParams, 0)
	}
	if params.Limit != 0 {
		queryParams = append(queryParams, params.Limit)
	} else {
		queryParams = append(queryParams, 10)
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
	rows, err = stmt.Query(queryParams...)
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
		var rolesID []uint8
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Status, &rolesID, &user.IsAdmin, &user.DepartmentID, &user.Status, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		if rolesID != nil {
			user.RolesID = strings.Split(string(rolesID), ",")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &dto.HasTotalResponseData{
		Total: total,
		List:  users,
	}, nil
}

func (u *userRepository) FindUserByParamsAndOutRoleID(roleId string, params *dto.QueryUsersParams) (*dto.HasTotalResponseData, error) {
	count := `
			SELECT count(*) FROM users WHERE delete_time IS NULL AND users.id NOT IN (
				SELECT
					user_id
				FROM
					users_roles
				WHERE
					role_id = ?)
					`

	var total int
	rows, err := db.DB.Query(count, roleId)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()

	}(rows)
	if rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	query := `
		SELECT
			id, account,nickname, status,GROUP_CONCAT(ur.role_id), is_admin, department_id, status, create_time, update_time
		FROM
			users
				LEFT JOIN users_roles ur ON users.id = ur.user_id
		WHERE
			users.id NOT IN (
				SELECT
					user_id
				FROM
					users_roles
				WHERE
					role_id = ?
    )
		`
	var queryParams []interface{}
	queryParams = append(queryParams, roleId)
	if params.ID != "" {
		query += " AND id = ?"
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
		query += " AND update_time <= ?"
		queryParams = append(queryParams, params.EndTime)
	}

	if params.DepartmentID != "" {
		query += " AND department_id = ?"
		queryParams = append(queryParams, params.DepartmentID)
	}

	if params.Account != "" {
		query += " AND account LIKE ?"
		queryParams = append(queryParams, "%"+params.Account+"%")
	}

	query += " GROUP BY users.id"
	query += ` LIMIT ?, ?`
	if params.Offset != 0 {
		queryParams = append(queryParams, params.Offset)
	} else {
		queryParams = append(queryParams, 0)
	}
	if params.Limit != 0 {
		queryParams = append(queryParams, params.Limit)
	} else {
		queryParams = append(queryParams, 10)
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
	rows, err = stmt.Query(queryParams...)
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
		var rolesID []uint8
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Status, &rolesID, &user.IsAdmin, &user.DepartmentID, &user.Status, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		if rolesID != nil {
			user.RolesID = strings.Split(string(rolesID), ",")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &dto.HasTotalResponseData{
		Total: total,
		List:  users,
	}, nil
}

func (u *userRepository) FindUserByAccount(account string) (*dto.SingleUserResponseHasPassword, bool) {
	if account == "" {
		return nil, false
	}
	user := &dto.SingleUserResponseHasPassword{}
	query := "SELECT id, account, nickname, password, is_admin, status,department_id, create_time, update_time, status FROM users WHERE account = ? AND delete_time IS NULL"
	rows, err := db.DB.Query(query, account)
	if err != nil {
		return nil, false
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Password, &user.IsAdmin, &user.Status, &user.DepartmentID, &user.CreateTime, &user.UpdateTime, &user.Status)
		if err != nil {
			return nil, false
		}
		return user, true
	}
	return nil, false
}

func (u *userRepository) FindUserById(id string) (*dto.UsersSingleResponse, bool) {
	user, err := u.GetUser(id)
	if err != nil {
		return nil, false
	}
	// 资源不存在
	if user == nil {
		return nil, false
	}
	return user, true
}

func (u *userRepository) DeleteUser(id string) error {
	stmt, err := db.DB.Prepare(`DELETE  FROM users WHERE id = ?`)
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

func (u *userRepository) UpdateUser(params *dto.UpdateUserRequest, id string) error {
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

	if params.DepartmentID != "" {
		queryParams = append(queryParams, params.DepartmentID)
		query += "department_id = ?, "
		hasSet = true
	}

	if params.Status != "" {
		queryParams = append(queryParams, params.Status)
		query += "status = ?, "
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

func (u *userRepository) GetUserByRoleID(roleID string, params dto.Page) (*dto.HasTotalResponseData, error) {
	var users []*dto.SingleUserByRoleIDResponse
	countQuery := "SELECT COUNT(*) FROM users_roles WHERE role_id = ? "
	count, err := db.DB.Query(countQuery, roleID)
	if err != nil {
		return nil, err
	}
	defer func(count *sql.Rows) {
		_ = count.Close()
	}(count)

	var total int
	if count.Next() {
		if err := count.Scan(&total); err != nil {
			return nil, err
		}
	}

	query := `
    SELECT u.id,
       u.account,
       u.nickname,
       u.avatar,
       u.is_admin,
       u.create_time,
       u.update_time,
       u.status,
       u.department_id,
       GROUP_CONCAT(users_roles.role_id) AS role_ids
	FROM users_roles
			 CROSS JOIN cms.users u ON u.id = users_roles.user_id
	WHERE user_id IN (SELECT user_id FROM users_roles WHERE role_id = ?)
	GROUP BY u.id, u.account, u.nickname, u.avatar, u.create_time, u.update_time, u.status, u.department_id
	LIMIT ?
	OFFSET ?
    `
	rows, err := db.DB.Query(query, roleID, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)
	for rows.Next() {
		var user dto.SingleUserByRoleIDResponse
		var rolesID []uint8
		err := rows.Scan(
			&user.ID,
			&user.Account,
			&user.Nickname,
			&user.Avatar,
			&user.IsAdmin,
			&user.CreateTime,
			&user.UpdateTime,
			&user.Status,
			&user.DepartmentID,
			&rolesID,
		)
		if err != nil {
			return nil, err
		}

		if rolesID != nil {
			user.RolesID = strings.Split(string(rolesID), ",")
		}
		users = append(users, &user)
	}

	return &dto.HasTotalResponseData{
		Total: total,
		List:  users,
	}, nil
}

func (u *userRepository) ExportExcel(params *dto.ExportExcelResponse) ([]*dto.UsersSingleResponse, error) {
	query := "SELECT id, account, nickname, avatar , department_id ,status, create_time, update_time  FROM users WHERE id IN "
	value := ""
	hasSet := false
	for _, id := range params.IDs {
		value += id + ","
		hasSet = true
	}
	if hasSet {
		value = value[:len(value)-1]
	}
	query += "(" + value + ")"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var users []*dto.UsersSingleResponse
	for rows.Next() {
		var user dto.UsersSingleResponse
		err := rows.Scan(&user.ID, &user.Account, &user.Nickname, &user.Avatar, &user.DepartmentID, &user.Status, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
