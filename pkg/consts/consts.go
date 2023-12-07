package consts

import "time"

var (
	JwtSecret = []byte("Campus-Hub-Secret")
)

const (
	// Database Address
	MySQLDefaultDSN = "campus:campushub@tcp(127.0.0.1:3306)/campus?charset=utf8mb4&parseTime=True&loc=Local"
	VercelCloudDSN  = ""

	// Logger
	LogFilePath = "./log"

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond

	// Server Name
	GatewayServiceName = "campus-gateway"
	UserServiceName    = "campus-user"
	DeviceServiceName  = "campus-device"

	// Server Port
	GatewayListenAddress       = "127.0.0.1:8080"
	UserServiceListenAddress   = "127.0.0.1:8081"
	DeviceServiceListenAddress = "127.0.0.1:8082"

	// Model Table Name
	UserTableName      = "user"
	AuthorityTableName = "authority"
	DeviceTableName    = "device"

	// Snowflake Algorithm
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// Etcd
	EtcdEndpoints = "127.0.0.1:2379"

	// Limit
	MaxConnections  = 1000
	MaxQPS          = 100
	MaxVideoSize    = 300000
	MaxListLength   = 100
	MaxIdleConns    = 10
	MaxOpenConns    = 100
	ConnMaxLifetime = 10 * time.Second

	// Aliyun SDK
	// AccessKey Expires at 2023-03-03-16-00-00
	//OSSEndpoint        = "files.ozline.icu"               // 默认启用域名绑定
	//OSSAccessKeyID     = "LTAI5t6gqQgzCzVgdUWw6uip"       // AccessKeyID
	//OSSAccessKeySecret = "diAQZbzqI6JJm53IoKSdGAudNZIbol" // AccessKeySecret
	//OSSBucketName      = "ozliinex"                       // 桶名
	//UplaodRoutines     = 3                                // 并发数量
	//MainDirectory      = "tiktok"                         // 主目录
	//PartSize           = 100 * 1024

	// Page
	PageNum  = 1
	PageSize = 10
	//
	//// ffmpeg
	//FrameNum = 1

	// Security
	PassWordCost = 12 // 密码加密强度
	Salt         = "Campus-Hub"
)
