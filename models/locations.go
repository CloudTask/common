package models

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
	Key      string `json:"key" bson:"key"`            //主机Key
	Name     string `json:"name" bson:"name"`          //主机名称
	IPAddr   string `json:"ipaddr" bson:"ipaddr"`      //Ip地址
	APIAddr  string `json:"apiaddr" bson:"apiaddr"`    //API地址
	OS       string `json:"os"  bson:"os"`             //系统环境
	Platform string `json:"platform"  bson:"platform"` //运行平台
	Status   int    `json:"status" bson:"status"`      //状态(0:在线 -1:离线)
}

/*
集群名称定义
*/
const (
	CLUSTER_DEV      = "dev"
	CLUSTER_GDEV     = "gdev"
	CLUSTER_GQC      = "gqc"
	CLUSTER_WH7      = "wh7"
	CLUSTER_E3       = "e3"
	CLUSTER_E4       = "e4"
	CLUSTER_E11      = "e11"
	CLUSTER_CROSS_DC = "cross-dc"
)

/*
 位置信息
*/
type WorkLocation struct {
	Location string    `json:"location" bson:"location"` //位置名称
	Cluster  string    `json:"cluster" bson:"cluster"`   //隶属调度集群
	Group    []*Group  `json:"group" bson:"group"`       //分组信息
	Server   []*Server `json:"server" bson:"server"`     //主机信息
}
