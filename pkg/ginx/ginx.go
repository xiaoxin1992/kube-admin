package ginx

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/pkg/ginx/middleware"
	"github.com/xiaoxin1992/kube-admin/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func Init(host string, port int, level string) *GinX {
	r := gin.Default()
	if level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	gx := &GinX{
		Route: r,
		Service: http.Server{
			Addr:              fmt.Sprintf("%s:%d", host, port),
			Handler:           r,
			ReadTimeout:       60 * time.Second,
			ReadHeaderTimeout: 60 * time.Second,
			WriteTimeout:      60 * time.Second,
			IdleTimeout:       60 * time.Second,
			MaxHeaderBytes:    1 << 25},
		logger: logger.GetLogger().S("http"),
	}
	return gx
}

type GinX struct {
	Route   *gin.Engine
	Service http.Server
	logger  *zap.SugaredLogger
}

func (g *GinX) register() {
	for _, r := range route {
		if !r.Auth {
			g.Route.Handle(r.Method, r.URI, r.Handlers)
		} else {
			g.Route.Handle(r.Method, r.URI, middleware.JWTAuth(g.logger), r.Handlers)
		}

	}
}

func (g *GinX) Start() (err error) {
	g.logger.Infof("start http %s", g.Service.Addr)
	g.register()
	if err = g.Service.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		} else {
			g.logger.Infof("start http server error %s", err.Error())
		}
	}
	return
}

func (g *GinX) Stop(ctx context.Context) (err error) {
	timeOut, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	err = g.Service.Shutdown(timeOut)
	return
}

func (g *GinX) WaitStop(sign chan os.Signal, ctx context.Context) {
	for {
		select {
		case sg := <-sign:
			switch v := sg.(type) {
			default:
				g.logger.Infof("receive signal %v, shutdown http\n", v.String())
				if err := g.Stop(ctx); err != nil {
					g.logger.Infof("http graceful shutdown err: %s force exit", err)
				} else {
					g.logger.Info("http service stop complete")
				}
				return
			}

		}
	}
}
