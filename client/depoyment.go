package client

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	Name     string
	Status   string
	Labels   string
	Num      string
	CreateAt string
}

func GetDeploymentList(name, namespace string) interface{} {
	clientSet, _ := InitClient(name)
	deploymentList, _ := clientSet.AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	deploymentLists := make([]*Deployment, 0)
	for _, item := range deploymentList.Items {
		d := &Deployment{
			Name:     item.Name,
			Num:      fmt.Sprintf("%s/%s", item.Status.AvailableReplicas, item.Spec.Replicas),
			CreateAt: item.Status.Conditions[0].LastUpdateTime.String(),
		}
		if item.Status.AvailableReplicas == int32(0) {
			d.Status = "未运行"
		}
		d.Status = "运行中"

		l := item.Labels
		fmt.Println(l)

		deploymentLists = append(deploymentLists, d)
	}
	return deploymentLists
}
