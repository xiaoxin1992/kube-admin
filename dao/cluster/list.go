package cluster

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/cluster"
	"github.com/xiaoxin1992/kube-admin/pkg/database"
	"strings"
)

func (d *Dao) ListClusterCount(ctx context.Context, request models.QueryList) (total int64, err error) {
	db := database.GetPool().GetSqlDB()
	sql := "select count(id) from cluster"
	var args []interface{}
	if strings.TrimSpace(request.Query) != "" {
		sql += " where zone like ?"
		args = append(args, fmt.Sprintf("%%%s%%", request.Query))
	}
	err = db.QueryRowContext(ctx, sql, args...).Scan(&total)
	return
}

func (d *Dao) ListCluster(ctx context.Context, request models.QueryList) (response []models.QueryCluster, err error) {
	db := database.GetPool().GetSqlDB()
	response = make([]models.QueryCluster, 0)
	sql := "select id, zone, host, token, remark, version, create_time, update_time from cluster"
	var args []interface{}
	if strings.TrimSpace(request.Query) != "" {
		sql += " where zone like ?"
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
		var cluster models.QueryCluster
		err = rows.Scan(&cluster.Id, &cluster.Zone, &cluster.Host, &cluster.Token, &cluster.Remark, &cluster.Version, &cluster.CreateTime, &cluster.UpdateTime)
		if err != nil {
			d.logger.Errorf("fail to scan row, err:%v", err)
			continue
		}
		response = append(response, cluster)
	}
	return
}
