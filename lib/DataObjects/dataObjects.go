package DataObjects

import (
	"database/sql"
	"sync"
)



type DataNode struct{
	ActionType int
	RetryUp int
	RetryDown int
	Comment string
	Compression int
	Connection *sql.DB
	ConnUsed int
	Debug bool
	Dns string
	Gtid_port int
	HostgroupId int
	Hostgroups []Hostgroup
	Ip string
	MaxConnection int
	MaxLatency int
	MaxReplication_lag int
	Name string
	NodeTCPDown bool
	Password string
	Port int
	Processed bool
	ProcessStatus int
	ProxyStatus string
	ReadOnly bool
	Ssl *SslCertificates
	Status map[string]string
	UseSsl bool
	User string
	Variables map[string]string
	Weight int
}


type DataCluster struct{
	ActiveFailover int
	FailBack		bool
	ActionNodes       map[string]DataNodePxc
	BackupReaders     map[string]DataNodePxc
	BackupWriters     map[string]DataNodePxc
	BackupHgReaderId  int
	BakcupHgWriterId  int
	CheckTimeout      int
	ClusterIdentifier int //cluster_id
	ClusterSize       int
	HasPrimary        bool
	ClusterName       string
	Comment           string
	Debug             bool
	FailOverNode	  DataNodePxc
	HasFailoverNode   bool
	Haswriter         bool
	HgReaderId        int
	HgWriterId        int
	Hostgroups        map[int]Hostgroup
	//	Hosts map[string] DataNode
	MainSegment int
	MonitorPassword string
	MonitorUser string
	Name string
	NodesPxc *SyncMap //[string] DataNodePxc // <ip:port,datanode>
	NodesPxcMaint []DataNodePxc
	MaxNumWriters int
	OffLineReaders map[string]DataNodePxc
	OffLineWriters map[string]DataNodePxc
	OffLineHgReaderID int
	OffLineHgWriterId int
	ReaderNodes map[string]DataNodePxc
	RequireFailover bool
	RetryDown int
	RetryUp int
	Singlenode bool
	SinglePrimary bool
	Size int
	Ssl *SslCertificates
	Status  int
	WriterIsReader int
	WriterNodes map[string]DataNodePxc

}


type SyncMap struct {
	sync.RWMutex
	internal map[string]DataNodePxc
}

type SslCertificates struct {
	sslClient string
	sslKey string
	sslCa string
	sslCertificatePath string
}

type PxcClusterView struct {
	//'HOST_NAME', 'UUID','STATUS','LOCAL_INDEX','SEGMENT'
	HostName string
	Uuid string
	Status string
	LocalIndex int
	Segment int
}


type VariableStatus struct{
	VarName string `db:"VARIABLE_NAME"`
	VarValue string `db:"VARIABLE_VALUE"`
}