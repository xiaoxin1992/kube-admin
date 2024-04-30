package cluster

import (
	"context"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

func (d *Dao) DeleteCluster(ctx context.Context, cluster *models.DeleteCluster) (err error) {
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
	sql := "delete from cluster where zone=?"
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		d.logger.Errorf("prepare statement error %v", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, cluster.Zone)
	return
}
