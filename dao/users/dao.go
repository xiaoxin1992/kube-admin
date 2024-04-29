package users

import (
	"context"
	"fmt"
	models "github.com/xiaoxin1992/kube-admin/models/users"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func NewDao() *dao {
	return &dao{
		logger: logger.GetLogger().S("MySQL").Named("users"),
	}
}

type dao struct {
	logger *zap.SugaredLogger
}

// 密码加密
func (d *dao) generatePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		d.logger.Errorf("password generate error %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (d *dao) VerifyPassword(ctx context.Context, username, password string) (*models.QueryUser, error) {
	query, err := d.QueryByPassword(ctx, username)
	if err != nil {
		d.logger.Errorf("query user info error %v", err)
		return &models.QueryUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(query.Password), []byte(password))
	if err != nil {
		d.logger.Errorf("password compare error %v", err)
		return &models.QueryUser{}, fmt.Errorf("password verify error %v", err)
	}
	return query, nil
}
