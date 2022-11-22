package models

import (
	"github.com/gookit/slog"
	"gorm.io/gorm"
)

type Cluster struct {
	gorm.Model
	Name      string `gorm:"column:name;type:varchar(100);" json:"name"`             // 集群的名字
	MasterUrl string `gorm:"column:master_url;type:varchar(255);" json:"master_url"` // master url
	Config    string `gorm:"column:config;type:text" json:"config"`                  // kubeConfig文件
}

func (table *Cluster) TableName() string {
	return "tb_k8s"
}

func GetClusterInfo(name string) *Cluster {
	data := &Cluster{
		Name: name,
	}
	err := DB.Where("name = ? ", name).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Info("没有查到集群信息")
			return &Cluster{}
		}
		slog.Warn(err)
	}
	return data
}
