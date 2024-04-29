package users

import (
	"context"
	"database/sql"
	"errors"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

// QueryByUsername 根据用户名获取当前用信息
func (d *dao) QueryByUsername(ctx context.Context, username string) (*models.QueryUser, error) {
	db := database.GetPool().GetSqlDB()
	query := &models.QueryUser{}
	sql := "select id, role, username, display_name, email, phone, create_time, update_time from users where username = ?"
	row := db.QueryRowContext(ctx, sql, username)
	err := row.Scan(&query.Id, &query.Role, &query.Username, &query.DisplayName, &query.Email, &query.Phone, &query.CreateTime, &query.UpdateTime)
	return query, err
}

// QueryByPassword 根据用户名获取用户认证需要的信息
func (d *dao) QueryByPassword(ctx context.Context, username string) (*models.QueryUser, error) {
	db := database.GetPool().GetSqlDB()
	query := &models.QueryUser{}
	sql := "select id, role, username, password, display_name, email, phone, create_time, update_time from users where username = ?"
	row := db.QueryRowContext(ctx, sql, username)
	err := row.Scan(&query.Id, &query.Role, &query.Username, &query.Password, &query.DisplayName, &query.Email, &query.Phone, &query.CreateTime, &query.UpdateTime)
	return query, err
}

// 判断用户是否存在，如果存在返回true 否则返回false

func (d *dao) ExistsUsername(ctx context.Context, username string) (bool, error) {
	_, err := d.QueryByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
