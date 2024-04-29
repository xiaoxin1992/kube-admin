package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/apis/auth"
	"github.com/xiaoxin1992/kube-admin/apis/ping"
	"github.com/xiaoxin1992/kube-admin/pkg/ginx"
	"net/http"
)

func register(method, uri string, auth bool, handlers gin.HandlerFunc) {
	ginx.Register(method, fmt.Sprintf("/api/%s", uri), auth, handlers)
}

func init() {
	/* ping 接口开始 */
	register(http.MethodGet, "ping", true, ping.List)
	register(http.MethodGet, "ping1", false, ping.List)
	/* ping 接口结束 */
	/* 登陆接口开始*/
	register(http.MethodPost, "login", false, auth.Login)
	/* 登陆接口结束*/

	/*cluster 集群管理接口开始*/
	/*cluster 集群管理接口结束*/
}
