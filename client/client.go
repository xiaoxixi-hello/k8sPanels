package client

import (
	"encoding/json"
	"github.com/gookit/slog"
	"github.com/ylinyang/k8sPanels/models"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientCmdApi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/tools/clientcmd/api/latest"
	clientCmdApiV1 "k8s.io/client-go/tools/clientcmd/api/v1"
)

func InitClient(name string) (*kubernetes.Clientset, error) {
	// 从数据库中获取集群信息
	clusterInfo := models.GetClusterInfo(name)
	kubeConfig := &clientCmdApiV1.Config{}
	if err := json.Unmarshal([]byte(clusterInfo.Config), kubeConfig); err != nil {
		slog.Warn("clusterInfo转换kubeConfig失败", err)
		return nil, err
	}
	// 切换匹配的版本
	configObject, err := latest.Scheme.ConvertToVersion(kubeConfig, clientCmdApi.SchemeGroupVersion)
	if err != nil {
		slog.Error("ConvertToVersion error. %v ", err)
		return nil, err
	}
	configInternal := configObject.(*clientCmdApi.Config)

	// 实例化配置信息
	clientConfig, err := clientcmd.NewDefaultClientConfig(*configInternal, &clientcmd.ConfigOverrides{
		ClusterDefaults: clientCmdApi.Cluster{Server: clusterInfo.MasterUrl},
	}).ClientConfig()

	if err != nil {
		slog.Error("build client config error. %v ", err)
		return nil, err
	}
	clientConfig.QPS = 60
	clientConfig.Burst = 60
	// 实例化客户端
	clientSet, err := kubernetes.NewForConfig(clientConfig)

	if err != nil {
		slog.Error("(%s) kubernetes.NewForConfig(%v) error.%v", clusterInfo.MasterUrl, err, clientConfig)
		return nil, err
	}
	return clientSet, nil
}
