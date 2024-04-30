package k8s

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	/*
		token 需要在k8s创建一个新用户
		apiVersion: v1
		kind: ServiceAccount
		metadata:
		  name: admin-user
		  namespace: kubernetes-dashboard
		---
		apiVersion: rbac.authorization.k8s.io/v1
		kind: ClusterRoleBinding
		metadata:
		  name: admin-user
		roleRef:
		  apiGroup: rbac.authorization.k8s.io
		  kind: ClusterRole
		  name: cluster-admin
		subjects:
		- kind: ServiceAccount
		  name: admin-user
		  namespace: kubernetes-dashboard
		---
		apiVersion: v1
		kind: Secret
		metadata:
		  name: admin-user
		  namespace: kubernetes-dashboard
		  annotations:
			kubernetes.io/service-account.name: "admin-user"
		type: kubernetes.io/service-account-token
	*/
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IngwWFpETmxTdE1raGExazZXR3ltNkVUQlVKU01uVkNfRXI4eXhuVFJDYkEifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI4OTQwM2EyMi0wOGE2LTRiMTUtYjdkOC0wZGYwZWRlOTQ4M2IiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZXJuZXRlcy1kYXNoYm9hcmQ6YWRtaW4tdXNlciJ9.AGQPwJ-aWuJJnogR836Ha632hK1WdVtV9LD0VNUVlk8nik8zs2d4DFRrzLJeMPQ72Y7h8prfICdnsnuorJPDw2K0HQdykh982yOQjYypzkYvZjw0dzh6owOME_urgEP8p-OKbnQssnmun-giTVR17sCuwd1fuXkbkUK5EUXZZq2K3GNjGvpoOv6UGapXwKw_5bqRIhvt8acHA5QciwVLfmdkh3QI2Z6564XNIdna3Ck-3igql90PUhZHcWrr0HD_8xAdsb0-jbu_6Ws3qvQyzZxL0ZULv9hG3dRoPZoTo7G3NeT0u_0zTPX0C97hk3IYXqPyksOYCajAnyg7ibgrrw"
	k := NewClient("https://127.0.0.1:63774", token)
	clientSet, err := k.GetClientSet()
	if err != nil {
		t.Fatal(err)
	}
	v, err := clientSet.DiscoveryClient.ServerVersion()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
