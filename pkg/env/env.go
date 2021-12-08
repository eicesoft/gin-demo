package env

import (
	"flag"
	"strings"
)

const (
	DevEnv  = "dev"
	TestEnv = "test"
	StagEnv = "stag"
	ProdEnv = "prod"
)

var (
	active Environment
	dev    Environment = &environment{value: DevEnv}
	test   Environment = &environment{value: TestEnv}
	stag   Environment = &environment{value: StagEnv}
	prod   Environment = &environment{value: ProdEnv}
)
var _ Environment = (*environment)(nil)

type environment struct {
	value  string
	daemon bool
}

func (e environment) Value() string {
	return e.value
}

func (e environment) IsDaemon() bool {
	return e.daemon
}

type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsStag() bool
	IsProd() bool
	SetDaemon(b bool)
	IsDaemon() bool
	p()
}

func (e *environment) SetDaemon(b bool) {
	e.daemon = b
}

func (e *environment) IsDev() bool {
	return e.value == DevEnv
}

func (e *environment) IsTest() bool {
	return e.value == TestEnv
}

func (e *environment) IsStag() bool {
	return e.value == StagEnv
}

func (e *environment) IsProd() bool {
	return e.value == ProdEnv
}

func (e *environment) p() {}

func init() {
	env := flag.String("env", "", "请输入运行环境:\n dev:开发环境\n test:测试环境\n stag:预上线环境\n prod:正式环境\n")
	daemon := flag.Bool("d", false, "是否后台守护进程方式运行")

	//if !flag.Parsed() {
	//	flag.Parse()
	//}

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "test":
		active = test
	case "stag":
		active = stag
	case "prod":
		active = prod
	default: //默认为Dev环境
		active = dev
	}

	active.SetDaemon(*daemon)
}

// Get 当前配置的env
func Get() Environment {
	return active
}
