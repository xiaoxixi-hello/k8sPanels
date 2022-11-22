package client

import (
	"context"
	"github.com/gookit/slog"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodeInfo struct {
	NodeName      string
	Status        string
	DockerVersion string
	IpAddr        string
	PodCidr       string
}

func GetNode(name string) (*[]NodeInfo, error) {
	clientSet, err := InitClient(name)
	if err != nil {
		slog.Errorf("获取集群信息异常")
		return nil, err
	}
	nodeLists, _ := clientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	nodeList := make([]NodeInfo, 0)
	for _, n := range nodeLists.Items {
		slog.Info()
		node := NodeInfo{
			NodeName:      n.Name,
			Status:        string(n.Status.Conditions[0].Type),
			DockerVersion: n.Status.NodeInfo.ContainerRuntimeVersion,
			IpAddr:        n.Status.Addresses[0].Address,
			PodCidr:       n.Spec.PodCIDR,
		}
		nodeList = append(nodeList, node)
	}
	return &nodeList, nil
}
