package tracing

import (
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func New(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
