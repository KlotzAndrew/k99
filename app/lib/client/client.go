package client

import (
	"context"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func PingService(ctx context.Context, url, str string) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		panic("no span in PingService, gotta make a new one")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")

	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	if _, err := client.Do(req); err != nil {
		panic(err.Error())
	}
}

func ContextFromHTTP(name string, r *http.Request) (context.Context, opentracing.Span) {
	wireCtx, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header),
	)
	if err != nil {
		switch err {
		default:
			panic(err)
		case opentracing.ErrUnsupportedFormat:
			panic(err)
		case opentracing.ErrSpanContextNotFound:
			// all good
		}
	}
	serverSpan := opentracing.StartSpan(
		"pinging",
		ext.RPCServerOption(wireCtx),
	)
	ctx := opentracing.ContextWithSpan(context.Background(), serverSpan)

	return ctx, serverSpan
}
