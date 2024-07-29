package config

const (
	FlagConfigPath = "config-path"

	DBDialectMysql = "mysql"

	EnvVarConfigFilePath = "CONFIG_FILE_PATH"
	EnvVarDBUserName     = "DB_USERNAME"
	EnvVarDBUserPass     = "DB_PASSWORD"
	EnvVarPrivateKey     = "PRIVATE_KEY"

	DefaultCreateBundleSlotInterval = 2500 // around 10MB

	DefaultBlockSyncThreshold = 1000

	DefaultReUploadBundleThreshold = 3600 // in second

	DefaultConcurrencyLimit = 5
)
