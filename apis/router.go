package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/apis/auth"
	"github.com/xiaoxin1992/kube-admin/apis/cluster"
	"github.com/xiaoxin1992/kube-admin/apis/configmap"
	"github.com/xiaoxin1992/kube-admin/apis/deployment"
	"github.com/xiaoxin1992/kube-admin/apis/namespace"
	"github.com/xiaoxin1992/kube-admin/apis/nodes"
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

	/* namespace 管理接口开始*/
	register(http.MethodGet, "/namespace/list", true, namespace.ListNamespace)
	register(http.MethodPost, "/namespace/create", true, namespace.CreateNamespace)
	register(http.MethodPost, "/namespace/delete", true, namespace.DeleteNamespace)
	/* namespace 管理接口结束*/

	/* node 管理接口开始*/
	register(http.MethodGet, "/nodes/list", true, nodes.ListNode)
	/* node 管理接口结束*/

	/* pods 管理接口开始*/
	register(http.MethodGet, "/pods/list", true, pods.ListPods)
	register(http.MethodPost, "/pods/create", true, pods.CreatePod)
	register(http.MethodPost, "/pods/delete", true, pods.DeletePod)
	/* pods 管理接口结束*/

	/* configmap 管理接口开始*/
	register(http.MethodPost, "/configmap/create", true, configmap.CreateConfigmap)
	register(http.MethodPost, "/configmap/delete", true, configmap.DeleteConfigmap)
	register(http.MethodPost, "/configmap/update", true, configmap.UpdateConfigmap)
	register(http.MethodGet, "/configmap/list", true, configmap.ListConfigmap)
	register(http.MethodGet, "/configmap/detail", true, configmap.DetailConfigmap)
	/* configmap 管理接口结束*/

	/* deployment 管理接口开始*/
	register(http.MethodPost, "/deployment/create", true, deployment.CreateDeployment)
	register(http.MethodPost, "/deployment/delete", true, deployment.DeleteDeployment)
	register(http.MethodPost, "/deployment/update", true, deployment.Updatedeployment)
	register(http.MethodGet, "/deployment/list", true, deployment.ListDeployment)
	register(http.MethodGet, "/deployment/detail", true, deployment.DetailDeployment)
	/* deployment 管理接口结束*/

}
