package config

import "time"

const (
	JWTRefreshExpireTime = 7 * 24 * time.Hour
	JWTAccessExpireTime  = 10 * time.Minute
)
