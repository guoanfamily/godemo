package common

import (
	"github.com/jmoiron/sqlx"
	"github.com/go-redis/redis"
)

var Db *sqlx.DB
var Rds *redis.Client
