module cookhub.com/app

go 1.17

require (
	cookhub.com/app/api/v1/onboarding v0.0.0-00010101000000-000000000000
	cookhub.com/app/api/v1/recipes v0.0.0-00010101000000-000000000000
	cookhub.com/app/middleware/auth v0.0.0-00010101000000-000000000000
	cookhub.com/app/models v0.0.0-00010101000000-000000000000
	cookhub.com/app/third_party/gofirebase v0.0.0-00010101000000-000000000000
	github.com/cenkalti/backoff/v4 v4.1.1
	github.com/labstack/echo/v4 v4.6.3
	github.com/lib/pq v1.10.0
)

require (
	cloud.google.com/go v0.99.0 // indirect
	cloud.google.com/go/firestore v1.6.1 // indirect
	cloud.google.com/go/storage v1.10.0 // indirect
	cookhub.com/app/api/entities v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/repositories v0.0.0-00010101000000-000000000000 // indirect
	firebase.google.com/go v3.13.0+incompatible // indirect
	github.com/census-instrumentation/opencensus-proto v0.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cncf/udpa/go v0.0.0-20201120205902-5459f2c99403 // indirect
	github.com/cncf/xds/go v0.0.0-20210805033703-aa0b78936158 // indirect
	github.com/envoyproxy/go-control-plane v0.9.10-0.20210907150352-cf90f659a021 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211013171255-e13a2654a71e // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.61.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace cookhub.com/app/third_party/gofirebase => ./third_party/firebase

replace cookhub.com/app/middleware/auth => ./middleware/auth

replace cookhub.com/app/api/v1/onboarding => ./api/v1/onboarding

replace cookhub.com/app/api/v1/recipes => ./api/v1/recipes

replace cookhub.com/app/api/entities => ./api/entities

replace cookhub.com/app/models => ./models

replace cookhub.com/app/repositories => ./repositories

replace cookhub.com/app/cache => ./cache
