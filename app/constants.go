package app

import "github.com/Metchain/Metblock/utils/limits"

const (
	leveldbCacheSizeMiB = 256

	currentDatabaseVersion = 1
)

var desiredLimits = &limits.DesiredLimits{
	FileLimitWant: 2048,
	FileLimitMin:  1024,
}
