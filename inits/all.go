package inits

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"goWebDemo/conf"

	"github.com/fsnotify/fsnotify"
	"github.com/pmars/beego/logs"
	"github.com/pmars/gotools"
	"github.com/spf13/viper"
)

var (
	LogoPath *string
)

// 第一步先初始相关数据
func InitAll() {
	configFile := flag.String("c", "/Users/xiaoh/go/src/goWebDemo/conf/config_dev.json", "config json file path")
	hostsFile := flag.String("h", "/Users/xiaoh/go/src/goWebDemo/conf/hosts_dev.json", "source host json file path")
	LogoPath = flag.String("l", "/Users/xiaoh/go/src/goWebDemo/conf/logo.ico", "logo path")
	flag.Parse()

	logs.Debug("Start InitAll, configFile:%v, hostsFile:%v logoPath:%v", *configFile, *hostsFile, *LogoPath)

	LoadConfig(*configFile, conf.Config)
	LoadConfig(*hostsFile, conf.Hosts)

	viper.SetConfigFile(*configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}

	viper.Set("Verbose", true)
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		LoadConfig(*configFile, conf.Config)
		LoadConfig(*hostsFile, conf.Hosts)
		configureWorld()
		logs.Info("LoadConfig File:%v Changed, Message:%v\n", configFile, e.Name)
	})
	initLog()
	configureWorld()
	logs.Debug("Config:%v", gotools.Data2Str(conf.Config))
	logs.Debug("Hosts:%v", gotools.Data2Str(conf.Hosts))
}

func LoadConfig(path string, v interface{}) (err error) {
	var j []byte
	if j, err = ioutil.ReadFile(path); err != nil {
		return
	}
	if err = json.Unmarshal(j, v); err != nil {
		return
	}
	return
}

func configureWorld() {
	logs.Info("Init All Configure Parameters Now ...")

	initMysqlEngine()
	// initRedisEngine()
	// initTask()

	logs.Info("Init Models Configure Parameters Done!")
}
