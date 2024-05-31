package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaoxin1992/kube-admin/apis/auth"
	"github.com/xiaoxin1992/kube-admin/apis/cluster"
	"github.com/xiaoxin1992/kube-admin/apis/configmap"
	"github.com/xiaoxin1992/kube-admin/apis/cronjob"
	"github.com/xiaoxin1992/kube-admin/apis/daemonset"
	"github.com/xiaoxin1992/kube-admin/apis/deployment"
	"github.com/xiaoxin1992/kube-admin/apis/namespace"
	"github.com/xiaoxin1992/kube-admin/apis/nodes"
	"github.com/xiaoxin1992/kube-admin/apis/ping"
	"github.com/xiaoxin1992/kube-admin/apis/pods"
	"github.com/xiaoxin1992/kube-admin/apis/pv"
	"github.com/xiaoxin1992/kube-admin/apis/pvc"
	"github.com/xiaoxin1992/kube-admin/apis/secret"
	"github.com/xiaoxin1992/kube-admin/apis/service"
	"github.com/xiaoxin1992/kube-admin/apis/statefulset"
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
	register(http.MethodGet, "wsPing", false, ping.Ping)
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
	register(http.MethodGet, "/pods/logs", false, pods.LogsPods)
	register(http.MethodGet, "/pods/cmd", false, pods.CmdPods)
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
	register(http.MethodPost, "/deployment/update", true, deployment.UpdateDeployment)
	register(http.MethodGet, "/deployment/list", true, deployment.ListDeployment)
	register(http.MethodGet, "/deployment/detail", true, deployment.DetailDeployment)
	/* deployment 管理接口结束*/

	/* service 管理接口开始*/
	register(http.MethodPost, "/service/create", true, service.CreateService)
	register(http.MethodPost, "/service/delete", true, service.DeleteService)
	register(http.MethodPost, "/service/update", true, service.UpdateService)
	register(http.MethodGet, "/service/list", true, service.ListService)
	register(http.MethodGet, "/service/detail", true, service.DetailService)
	/* service 管理接口结束*/

	/* secret 管理接口开始*/
	register(http.MethodPost, "/secret/create", true, secret.CreateSecret)
	register(http.MethodPost, "/secret/delete", true, secret.DeleteSecret)
	register(http.MethodPost, "/secret/update", true, secret.UpdateSecret)
	register(http.MethodGet, "/secret/list", true, secret.ListSecret)
	register(http.MethodGet, "/secret/detail", true, secret.DetailSecret)
	/* secret 管理接口结束*/

	/* daemonSet 管理接口开始*/
	register(http.MethodPost, "/ds/create", true, daemonset.CreateDaemonSet)
	register(http.MethodPost, "/ds/delete", true, daemonset.DeleteDaemonSet)
	register(http.MethodPost, "/ds/update", true, daemonset.UpdateDaemonSet)
	register(http.MethodGet, "/ds/list", true, daemonset.ListDaemonSet)
	register(http.MethodGet, "/ds/detail", true, daemonset.DetailDaemonSet)
	/* daemonSet 管理接口结束*/

	/* cronjob 管理接口开始*/
	register(http.MethodPost, "/cronjob/create", true, cronjob.CreateCronjob)
	register(http.MethodPost, "/cronjob/delete", true, cronjob.DeleteCronjob)
	register(http.MethodPost, "/cronjob/update", true, cronjob.UpdateCronjob)
	register(http.MethodGet, "/cronjob/list", true, cronjob.ListCronjob)
	register(http.MethodGet, "/cronjob/detail", true, cronjob.DetailCronjob)
	/* cronjob 管理接口结束*/

	/* StatefulSet 管理接口开始*/
	register(http.MethodPost, "/sts/create", true, statefulset.CreateStateFulSet)
	register(http.MethodPost, "/sts/delete", true, statefulset.DeleteStateFulSet)
	register(http.MethodPost, "/sts/update", true, statefulset.UpdateStateFulSet)
	register(http.MethodGet, "/sts/list", true, statefulset.ListStateFulSet)
	register(http.MethodGet, "/sts/detail", true, statefulset.DetailStateFulSet)
	/* StatefulSet 管理接口结束*/

	/* pv 管理接口开始*/
	register(http.MethodPost, "/pv/create", true, pv.CreatePV)
	register(http.MethodPost, "/pv/delete", true, pv.DeletePV)
	register(http.MethodPost, "/pv/update", true, pv.UpdatePV)
	register(http.MethodGet, "/pv/list", true, pv.ListPV)
	register(http.MethodGet, "/pv/detail", true, pv.DetailPV)
	/* pv 管理接口结束*/

	/* pvc 管理接口开始*/
	register(http.MethodPost, "/pvc/create", true, pvc.CreatePersistentVolumeClaim)
	register(http.MethodPost, "/pvc/delete", true, pvc.DeletePersistentVolumeClaim)
	register(http.MethodGet, "/pvc/list", true, pvc.ListPersistentVolumeClaim)
	register(http.MethodGet, "/pvc/detail", true, pvc.DetailPersistentVolumeClaim)
	/* pvc 管理接口结束*/
}
