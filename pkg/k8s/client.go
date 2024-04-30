package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func NewClient(host, token string) *Client {
	return &Client{
		Host:  host,
		Token: token,
	}
}

type Client struct {
	Host  string
	Token string
}

func (c *Client) config() *rest.Config {
	return &rest.Config{
		Host:        c.Host,
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
}

func (c *Client) GetClientSet() (clientSet *kubernetes.Clientset, err error) {
	clientSet, err = kubernetes.NewForConfig(c.config())
	return
}
