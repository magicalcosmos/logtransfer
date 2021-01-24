package main

import (
	"fmt"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	// 1.加载配置文件
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init config, err: %v\n", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)
	// 2.初始化

	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init es, err: %v\n", err)
		return
	}
	fmt.Println("init es success")
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafak, err: %v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 3.从kafka取日志数据
	// 4.发送到es
	select {}
}
