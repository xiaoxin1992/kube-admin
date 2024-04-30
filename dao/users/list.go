package users

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
	"strings"
)

func (d *dao) ListUserCount(ctx context.Context, request models.QueryList) (total int64, err error) {
	db := database.GetPool().GetSqlDB()
	sql := "select count(id) from users"
	var args []interface{}
	if strings.TrimSpace(request.Query) != "" {
		sql += " where username like ?"
		args = append(args, fmt.Sprintf("%%%s%%", request.Query))
	}
	err = db.QueryRowContext(ctx, sql, args...).Scan(&total)
	return
}

func (d *dao) ListUsers(ctx context.Context, request models.QueryList) (response []models.QueryUser, err error) {
	db := database.GetPool().GetSqlDB()
	response = make([]models.QueryUser, 0)
	sql := "select id, role, username, display_name, email, phone, create_time, update_time from users"
	var args []interface{}
	if strings.TrimSpace(request.Query) != "" {
		sql += " where username like ?"
		args = append(args, fmt.Sprintf("%%%s%%", request.Query))
	}
	sql += " order by create_time desc"
	sql += " LIMIT ? OFFSET ?"
	offset := (request.Page - 1) * request.Size
	args = append(args, request.Size, offset)
	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("fail to prepare statement, err:%v", err)
		return
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		d.logger.Errorf("fail to execute query, err:%v", err)
	}
	for rows.Next() {
		var user models.QueryUser
		err = rows.Scan(&user.Id, &user.Role, &user.Username, &user.DisplayName, &user.Email, &user.Phone, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			d.logger.Errorf("fail to scan row, err:%v", err)
			continue
		}
		response = append(response, user)
	}
	return
}
