package validator

import (
	utils "trading-data-harvester/utils/env"
)

var (
	pageSizeStr = utils.GetEnvOrDefault("PAGE_SIZE", "10000")
)
