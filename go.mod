module github.com/atomix/go-client

require (
	github.com/atomix/api v0.0.0-20200202100958-13b24edbe32d
	github.com/atomix/go-framework v0.0.0-20200202102454-440bc2678f1c
	github.com/atomix/go-local v0.0.0-20200202105028-743d224c66eb
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/stretchr/testify v1.4.0
	google.golang.org/grpc v1.27.0
)

replace github.com/atomix/go-framework => ../go-framework

replace github.com/atomix/go-client => ../go-client

replace github.com/atomix/api => ../api
