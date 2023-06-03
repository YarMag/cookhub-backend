module island_pga

go 1.17

require (
	cookhub.com/app/db v0.0.0-00010101000000-000000000000
	ga_operators v0.0.0-00010101000000-000000000000
	github.com/labstack/echo/v4 v4.6.3
	island_algorithm v0.0.0-00010101000000-000000000000
)

require (
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)

replace cookhub.com/app/db => ./db

replace cookhub.com/app/cache => ./cache

replace island_algorithm => ./island_algorithm

replace ga_operators => ./ga_operators
