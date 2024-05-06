package cluster

import (
	"context"
	"database/sql"
	"errors"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
)

// QueryByZone  根据用token获取信息
func (d *Dao) QueryByZone(ctx context.Context, zone string) (*models.QueryCluster, error) {
	db := database.GetPool().GetSqlDB()
	query := &models.QueryCluster{}
	querySQL := "select id, zone, host, token, remark, version, create_time, update_time from cluster where zone = ?"
	row := db.QueryRowContext(ctx, querySQL, zone)
	err := row.Scan(&query.Id, &query.Zone, &query.Host, &query.Token, &query.Remark, &query.Version, &query.CreateTime, &query.UpdateTime)
	return query, err
}

// ExistsZone 判断用户是否存在，如果存在返回true 否则返回false
func (d *Dao) ExistsZone(ctx context.Context, zone string) (bool, error) {
	_, err := d.QueryByZone(ctx, zone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
