package envs

import (
	"flag"
	"os"
	"sync"
)

func IsProd() bool {
	return Is(envFlagProd)
}

func IsDev() bool {
	return Is(envFlagDev)
}

func IsTesting() bool {
	return Is(envFlagTesting)
}

func Is(flag string) bool {
	return getFlag() == flag
}

func Name() string {
	return getFlag()
}

var (
	once    sync.Once
	envFlag string
)

const (
	envFlagProd    = "prod"
	envFlagDev     = "dev"
	envFlagTesting = "testing"
)

func getFlag() string {
	once.Do(func() {
		if flag.Lookup("test.v") != nil {
			envFlag = envFlagTesting
		} else if f, ok := os.LookupEnv("ENV"); ok {
			envFlag = f
		} else {
			envFlag = envFlagProd
		}
	})
	return envFlag
}
