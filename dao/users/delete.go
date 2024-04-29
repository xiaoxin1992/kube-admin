package users

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

func (d *dao) DeleteUser(ctx context.Context, user *models.DeleteUser) (err error) {
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
	sql := "delete from users where username=?"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("prepare statement error %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.Username)
	return
}
