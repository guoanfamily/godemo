package common

import (
	"github.com/guoanfamily/sqlx"
	"github.com/go-redis/redis"
)

var Db *sqlx.DB
var Rds *redis.Client
