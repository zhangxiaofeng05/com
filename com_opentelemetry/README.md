## reference
https://dev.to/aurelievache/learning-go-by-examples-part-10-instrument-your-go-app-with-opentelemetry-and-send-traces-to-jaeger-distributed-tracing-1p4a

## example
https://github.com/zhangxiaofeng05/go-example/blob/main/opentelemetry/opentelemetry.go

## jaeger
http://127.0.0.1:16686  
https://www.jaegertracing.io/docs/1.46/getting-started/  
https://github.com/open-telemetry/opentelemetry-go

## opentelemetry-go-contrib
OpenTelemetry-Go的扩展集合：https://github.com/open-telemetry/opentelemetry-go-contrib  
 - [gin](https://go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin)
 - [grpc](https://go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc)
 - ...

### 其他扩展
uptrace的OpenTelemetry扩展：https://github.com/uptrace/opentelemetry-go-extra  
 - [gorm](https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelgorm)
 - ...