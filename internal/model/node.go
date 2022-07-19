package model

import (
	"github.com/hhandhuan/ku-bbs/pkg/db"
	"gorm.io/gorm"
)

type Nodes struct {
	Model
	Name  string `gorm:"column:name" db:"name" json:"name" form:"name"`     //节点名称
	Alias string `gorm:"column:alias" db:"alias" json:"alias" form:"alias"` //节点别名
	Desc  string `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`     //节点介绍
	Count int64  `gorm:"column:count" db:"count" json:"count" form:"count"` //资源统计
	Pid   int64  `gorm:"column:pid" db:"pid" json:"pid" form:"pid"`         //节点父级
	State int8   `gorm:"column:state" db:"state" json:"state" form:"state"` //节点状态:0-关闭/1-开启
}

type nodeModel struct {
	M     *gorm.DB
	table string
}

func Node() *nodeModel {
	return &nodeModel{M: db.DB.Model(&Nodes{}), table: "nodes"}
}