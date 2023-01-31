package env

import (
	"flag"
	"os"
	"sync"
)

// Flag environment flag
type Flag string

const (
	FlagProduct Flag = "product"
	FlagStaging Flag = "staging"
	FlagDev     Flag = "dev"
	FlagTesting Flag = "testing"
)

// DefaultEnvKey is the default environment variable
const DefaultEnvKey = "ENV"

var flagInitFunc = InitFromEnv(DefaultEnvKey)

// InitFromEnv 从环境变量读取环境配置
func InitFromEnv(envKey string) func() (Flag, bool) {
	return func() (Flag, bool) {
		flag, ok := os.LookupEnv(envKey)
		return Flag(flag), ok
	}
}

// SetInitFunc 设置环境初始化函数
func SetInitFunc(f func() (Flag, bool)) {
	flagInitFunc = f
}

// Current 返回当前环境名
func Current() Flag {
	return getFlag()
}

// IsTesting 判断是否是测试环境
func IsTesting() bool {
	return Current() == FlagTesting
}

// IsDev 判断是否是本地开发环境
func IsDev() bool {
	return Current() == FlagDev
}

// IsStaging 判断是否是预览环境
func IsStaging() bool {
	return Current() == FlagStaging
}

// IsProduct 判断是否是线上环境
func IsProduct() bool {
	return Current() == FlagProduct
}

var (
	envFlag Flag
	once    sync.Once
)

func getFlag() Flag {
	once.Do(func() {
		if flag.Lookup("test.v") != nil {
			envFlag = FlagTesting
		} else if f, ok := flagInitFunc(); ok {
			envFlag = f
		}
		if envFlag == "" {
			panic("init environment flag failed!")
		}
	})
	return envFlag
}
