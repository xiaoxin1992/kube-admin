package users

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

func (d *dao) UpdateUser(ctx context.Context, updateUser models.UpdateUser) (err error) {
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
	sql := "update users set display_name=?, email=?, phone=? where username=?"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("prepare statement error %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, updateUser.DisplayName, updateUser.Email, updateUser.Phone, updateUser.Username)
	return
}

func (d *dao) ResetPassword(ctx context.Context, resetPassword models.ResetPassword) (err error) {
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
	hexPassword, err := d.generatePassword(resetPassword.Password)
	if err != nil {
		d.logger.Errorf("generate password error %v", err)
		return
	}
	sql := "update users set password=? where username=?"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("prepare statement error %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, hexPassword, resetPassword.Username)
	return
}
