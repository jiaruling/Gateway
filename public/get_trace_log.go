package public

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/golang_utils/lib"
)

func GetTraceAndLogByContext(c *gin.Context) (*lib.TraceContext, *lib.Log) {
	trace, _ := c.Get("trace")
	traceContext, ok := trace.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	log := lib.GetLog()
	return traceContext, log
}

func GetTraceAndLog() (*lib.TraceContext, *lib.Log) {
	return lib.NewTrace(), lib.GetLog()
}
