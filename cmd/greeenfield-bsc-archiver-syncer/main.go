package main

import (
	"flag"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"greeenfield-bsc-archiver/config"
	syncerdb "greeenfield-bsc-archiver/db"
	"greeenfield-bsc-archiver/logging"
	"greeenfield-bsc-archiver/metrics"
	"greeenfield-bsc-archiver/syncer"
)

func initFlags() {
	flag.String(config.FlagConfigPath, "", "config file path")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		cfg            *config.SyncerConfig
		configFilePath string
	)
	initFlags()
	configFilePath = viper.GetString(config.FlagConfigPath)
	if configFilePath == "" {
		configFilePath = os.Getenv(config.EnvVarConfigFilePath)
	}
	cfg = config.ParseSyncerConfigFromFile(configFilePath)
	if cfg == nil {
		panic("failed to get configuration")
	}
	cfg.Validate()
	logging.InitLogger(&cfg.LogConfig)
	db := config.InitDBWithConfig(&cfg.DBConfig, true)
	blockDB := syncerdb.NewBlockSvcDB(db)
	bs := syncer.NewBlockIndexer(blockDB, cfg)
	err := os.MkdirAll(cfg.TempDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	go bs.StartLoop()

	if cfg.MetricsConfig.Enable {
		if cfg.MetricsConfig.HttpAddress == "" {
			cfg.MetricsConfig.HttpAddress = metrics.DefaultMetricsAddress
		}
		metric := metrics.NewMetrics(cfg.MetricsConfig.HttpAddress)
		go metric.Start()
	}

	select {}
}
