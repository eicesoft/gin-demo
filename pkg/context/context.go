package context

import (
	stdctx "context"
	"eicesoft/web-demo/pkg/trace"

	"go.uber.org/zap"
)

type Trace = trace.T

type StdContext struct {
	stdctx.Context
	Trace
	*zap.Logger
}
