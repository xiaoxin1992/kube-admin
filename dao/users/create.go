package users

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

func (d *dao) CreateUser(ctx context.Context, user *models.User) (err error) {
	db := database.GetPool().GetSqlDB()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		d.logger.Errorf("start transaction error %v", err)
		return
	}
	defer func() {
		if err != nil {
			d.logger.Errorf("transaction error %v", err)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	sql := "insert into users(username, display_name, password, role, email, phone) value (?,?,?,?,?,?)"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("prepare statement error %v", err)
		return
	}
	defer stmt.Close()
	hexPassword, err := d.generatePassword(user.Password)
	if err != nil {
		err = fmt.Errorf("password generate error %v", err)
		return
	}
	_, err = stmt.ExecContext(ctx, user.Username, user.DisplayName, hexPassword, user.Role, user.Email, user.Phone)
	return
}
