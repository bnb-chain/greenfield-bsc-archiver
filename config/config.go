package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"greeenfield-bsc-archiver/cache"
	syncerdb "greeenfield-bsc-archiver/db"
)

const (
	BSC = "BSC"
)

type SyncerConfig struct {
	Chain                            string        `json:"chain"`                                // support BSC
	BucketName                       string        `json:"bucket_name"`                          // BucketName is the identifier of bucket on Greenfield that store blob
	StartBlock                       uint64        `json:"start_block"`                          // StartBlock is used to init the syncer which block synced from, only need to provide once.
	CreateBundleBlockInterval        uint64        `json:"create_bundle_block_interval"`         // CreateBundleBlockInterval defines the number of block that syncer would assemble blobs and upload to bundle service
	BlockSyncThreshold               uint          `json:"block_sync_threshold"`                 // BlockSyncThreshold  defines the maximum number of blocks allowed in the blocks map before pausing the block fetching process.
	BundleServiceEndpoints           []string      `json:"bundle_service_endpoints"`             // BundleServiceEndpoints is a list of bundle service address
	RPCAddrs                         []string      `json:"rpc_addrs"`                            // RPCAddrs ETH or BSC RPC addr
	GnfdRpcAddr                      string        `json:"gnfd_rpc_addr"`                        // GnfdRpcAddr is the Greenfield RPC address
	TempDir                          string        `json:"temp_dir"`                             // TempDir is used to store blobs and created bundle
	PrivateKey                       string        `json:"private_key"`                          // PrivateKey is the key of bucket owner, request to bundle service will be signed by it as well.
	BundleNotSealedReuploadThreshold int64         `json:"bundle_not_sealed_reupload_threshold"` // BundleNotSealedReuploadThreshold for re-uploading a bundle if it cant be sealed within the time threshold.
	EnableIndivBlockVerification     bool          `json:"enable_indiv_block_verification"`      // EnableIndivBlobVerification is used to enable individual block verification, otherwise only bundle level verification is performed.
	DBConfig                         DBConfig      `json:"db_config"`
	MetricsConfig                    MetricsConfig `json:"metrics_config"`
	LogConfig                        LogConfig     `json:"log_config"`
	ConcurrencyLimit                 int           `json:"concurrency_limit"`
}

func (s *SyncerConfig) Validate() {
	if !strings.EqualFold(s.Chain, BSC) {
		panic("chain not support")
	}
	if len(s.BucketName) == 0 {
		panic("the Greenfield bucket name is not is not provided")
	}
	if len(s.BundleServiceEndpoints) == 0 {
		panic("BundleService endpoints should not be empty")
	}
	if len(s.RPCAddrs) == 0 {
		panic("eth rpc address should not be empty")
	}
	if len(s.TempDir) == 0 {
		panic("temp directory is not specified")
	}
	if len(s.PrivateKey) == 0 {
		panic("private key is not provided")
	}
	if s.Chain == BSC && s.CreateBundleBlockInterval > 1000 {
		panic("create_bundle_slot_interval is supposed to be less than 1000")
	}
	if s.BundleNotSealedReuploadThreshold <= 60 {
		panic("Bundle_not_sealed_reupload_threshold is supposed larger than 60 (s)")
	}

	s.DBConfig.Validate()
}

func (s *SyncerConfig) GetCreateBundleInterval() uint64 {
	if s.CreateBundleBlockInterval == 0 {
		return DefaultCreateBundleSlotInterval
	}
	return s.CreateBundleBlockInterval
}

func (s *SyncerConfig) GetBlockSyncThreshold() uint {
	if s.BlockSyncThreshold == 0 {
		return DefaultBlockSyncThreshold
	}
	return s.BlockSyncThreshold
}

func (s *SyncerConfig) GetReUploadBundleThresh() int64 {
	if s.BundleNotSealedReuploadThreshold == 0 {
		return DefaultReUploadBundleThreshold
	}
	return s.BundleNotSealedReuploadThreshold
}

type ServerConfig struct {
	Chain                  string        `json:"chain"`
	BucketName             string        `json:"bucket_name"`
	BundleServiceEndpoints []string      `json:"bundle_service_endpoints"` // BundleServiceEndpoints is a list of bundle service address
	CacheConfig            CacheConfig   `json:"cache_config"`
	DBConfig               DBConfig      `json:"db_config"`
	MetricsConfig          MetricsConfig `json:"metrics_config"`
}

func (s *ServerConfig) Validate() {
	if !strings.EqualFold(s.Chain, BSC) {
		panic("chain not support")
	}
	if len(s.BucketName) == 0 {
		panic("the Greenfield bucket name is not is not provided")
	}
	if len(s.BundleServiceEndpoints) == 0 {
		panic("BundleService endpoints should not be empty")
	}
	s.DBConfig.Validate()
}

type CacheConfig struct {
	CacheType string `json:"cache_type"`
	URL       string `json:"url"`
	CacheSize uint64 `json:"cache_size"`
}

func (c *CacheConfig) GetCacheSize() uint64 {
	if c.CacheSize != 0 {
		return c.CacheSize
	}
	return cache.DefaultCacheSize
}

type DBConfig struct {
	Dialect      string `json:"dialect"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Url          string `json:"url"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != DBDialectMysql {
		panic(fmt.Sprintf("only %s supported", DBDialectMysql))
	}
	if cfg.Dialect == DBDialectMysql && (cfg.Username == "" || cfg.Url == "") {
		panic("db config is not correct, missing username and/or url")
	}
	if cfg.MaxIdleConns == 0 || cfg.MaxOpenConns == 0 {
		panic("db connections is not correct")
	}
}

type MetricsConfig struct {
	Enable      bool   `json:"enable"`
	HttpAddress string `json:"http_address"`
	SPEndpoint  string `json:"sp_endpoint"`
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

func ParseSyncerConfigFromFile(filePath string) *SyncerConfig {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config SyncerConfig
	if err = json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	if config.DBConfig.Username == "" || config.DBConfig.Password == "" { // read password from ENV
		config.DBConfig.Username, config.DBConfig.Password = GetDBUsernamePasswordFromEnv()
	}
	if config.PrivateKey == "" { // read private key from ENV
		config.PrivateKey = os.Getenv(EnvVarPrivateKey)
	}
	if config.ConcurrencyLimit == 0 {
		config.ConcurrencyLimit = DefaultConcurrencyLimit
	}
	return &config
}

func ParseServerConfigFromFile(filePath string) *ServerConfig {
	bz, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config ServerConfig
	if err = json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}
	if config.DBConfig.Username == "" || config.DBConfig.Password == "" { // read password from ENV
		config.DBConfig.Username, config.DBConfig.Password = GetDBUsernamePasswordFromEnv()
	}
	return &config
}

func GetDBUsernamePasswordFromEnv() (string, string) {
	username := os.Getenv(EnvVarDBUserName)
	password := os.Getenv(EnvVarDBUserPass)
	return username, password
}

func InitDBWithConfig(cfg *DBConfig, writeAccess bool) *gorm.DB {
	var db *gorm.DB
	var err error
	var dialector gorm.Dialector

	if cfg.Dialect == DBDialectMysql {
		url := cfg.Url
		dbPath := fmt.Sprintf("%s:%s@%s", cfg.Username, cfg.Password, url)
		dialector = mysql.Open(dbPath)
	} else {
		panic(fmt.Sprintf("unexpected DB dialect %s", cfg.Dialect))
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             10 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			Colorful:                  true, // Disable color
		},
	)
	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	dbConfig, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbConfig.SetMaxIdleConns(cfg.MaxIdleConns)
	dbConfig.SetMaxOpenConns(cfg.MaxOpenConns)
	if writeAccess {
		syncerdb.AutoMigrateDB(db)
	}
	return db
}
