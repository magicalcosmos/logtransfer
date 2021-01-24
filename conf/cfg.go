package conf

// LogTransferCfg log transfer
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// KafkaCfg kafka
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg es
type ESCfg struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chan_size"`
	Nums     int    `ini:"nums"`
}
