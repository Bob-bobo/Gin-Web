package models

type UserRole struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

func QueryUserRole(userid, roleid int) int {
	row := db.Exec("select user_id,role_id from user_role_tab where user_id = ? and role_id = ?", userid, roleid)
	id := 0
	row.Scan(&id)
	return id
}
