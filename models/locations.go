package models

import "github.com/cloudtask/libtools/gzkwrapper"

/*
 分组信息
*/
type Group struct {
	Id     string   `json:"id" bson:"id"`         //组编号
	Name   string   `json:"name" bson:"name"`     //组名称
	Owners []string `json:"owners" bson:"owners"` //组管理员
}

/*
 服务器信息
*/
type Server struct {
	Key        string `json:"key" bson:"key"`               //主机Key
	Name       string `json:"name" bson:"name"`             //主机名称
	IPAddr     string `json:"ipaddr" bson:"ipaddr"`         //Ip地址
	APIAddr    string `json:"apiaddr" bson:"apiaddr"`       //API地址
	OS         string `json:"os"  bson:"os"`                //系统环境
	Platform   string `json:"platform"  bson:"platform"`    //运行平台
	Status     int    `json:"status" bson:"status"`         //状态(0:离线 1:在线)
	Alivestamp int64  `json:"alivestamp" bson:"alivestamp"` //存活时间戳
}

func CreateServer(key string, node *gzkwrapper.NodeData, status int) *Server {

	if node == nil {
		return nil
	}

	return &Server{
		Key:        key,
		Name:       node.HostName,
		IPAddr:     node.IpAddr,
		APIAddr:    node.APIAddr,
		OS:         node.OS,
		Platform:   node.Platform,
		Status:     status,
		Alivestamp: node.Alivestamp,
	}
}

/*
 位置信息
*/
type WorkLocation struct {
	Location string    `json:"location" bson:"location"` //位置名称
	Owners   []string  `json:"owners" bson:"owners"`     //管理员
	Group    []*Group  `json:"group" bson:"group"`       //分组信息
	Server   []*Server `json:"server" bson:"server"`     //主机信息
}
