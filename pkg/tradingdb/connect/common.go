package connect

import utils "trading-data-harvester/utils/env"

var (
	ClickhouseHost     = utils.GetEnvOrDefault("CLICKHOUSE_HOST", "127.0.0.1")
	ClickhousePort     = utils.GetEnvOrDefault("CLICKHOUSE_PORT", "9000")
	ClickhouseDatabase = utils.GetEnvOrDefault("CLICKHOUSE_DATABASE", "crypto")
	ClickhouseUsername = utils.GetEnvOrDefault("CLICKHOUSE_USERNAME", "admin")
	ClickhousePassword = utils.GetEnvOrDefault("CLICKHOUSE_PASSWORD", "ponk0132")
)
