package config

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func init() {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Println("动态配置客户端创建失败", err.Error())
	}
	config, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user",
		Group:  "dev",
	})
	if err != nil {
		log.Println("配置文件读取失败", err.Error())
	}
	log.Println(config)
	err = json.Unmarshal([]byte(config), &ServerConfigs)
	if err != nil {
		log.Println("配置文件转码失败", err.Error())
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "dataId",
		Group:  "group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("---服务配置发生变化---")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		log.Println("动态监听失败", err.Error())
	}
}

var ServerConfigs ServerConfig

type ServerConfig struct {
	Mysql Mysql `json:"Mysql"`
	Redis Redis `json:"Redis"`
	Token Token `json:"Token"`
}
type Mysql struct {
	DriverName           string `json:"DriverName"`
	Username             string `json:"Username"`
	Password             string `json:"Password"`
	ConnectionType       string `json:"ConnectionType"`
	Host                 string `json:"Host"`
	Port                 string `json:"Port"`
	DatabaseName         string `json:"DatabaseName"`
	ConnectionParameters string `json:"ConnectionParameters"`
}
type Redis struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
}
type Token struct {
	SecretKey string `json:"SecretKey"`
}
