package config

import (
	dbtypes "github.com/Station-Manager/types/database"
)

type Config struct {
	DatastoreConfigs []dbtypes.Config
}
