package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/apis/auth"
	"github.com/xiaoxin1992/kube-admin/apis/cluster"
	"github.com/xiaoxin1992/kube-admin/apis/ping"
	"github.com/xiaoxin1992/kube-admin/apis/pods"
	"github.com/xiaoxin1992/kube-admin/apis/users"
	"github.com/xiaoxin1992/kube-admin/pkg/ginx"
	"net/http"
)

func register(method, uri string, auth bool, handlers gin.HandlerFunc) {
	ginx.Register(method, fmt.Sprintf("/api/%s", uri), auth, handlers)
}

func init() {
	/* ping 接口开始 */
	register(http.MethodGet, "ping", false, ping.List)
	/* ping 接口结束 */

	/* 登陆接口开始*/
	register(http.MethodPost, "login", false, auth.Login)
	/* 登陆接口结束*/

	/* 用户接口开始 */
	register(http.MethodPost, "/user/create", true, users.CreateUser)
	register(http.MethodPost, "/user/update", true, users.UpdateUser)
	register(http.MethodPost, "/user/reset_password", true, users.ResetPassword)
	register(http.MethodPost, "/user/delete", true, users.DeleteUser)
	register(http.MethodGet, "/user/list", true, users.ListUser)
	/* 用户接口结束 */

	/*cluster 集群管理接口开始*/
	register(http.MethodPost, "/cluster/create", true, cluster.CrateCluster)
	register(http.MethodPost, "/cluster/update", true, cluster.UpdateCluster)
	register(http.MethodPost, "/cluster/delete", true, cluster.DeleteCluster)
	register(http.MethodGet, "/cluster/list", true, cluster.ListCluster)
	/*cluster 集群管理接口结束*/

	/* node 管理接口开始*/
	/* node 管理接口结束*/

	/* pods 管理接口开始*/
	register(http.MethodGet, "/pods/list", true, pods.ListPods)
	register(http.MethodPost, "/pods/create", true, pods.CreatePod)
	register(http.MethodPost, "/pods/delete", true, pods.DeletePod)
	/* pods 管理接口结束*/
}
