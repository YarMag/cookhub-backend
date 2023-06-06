module cookhub.com/app/api/v1/search

go 1.17

require github.com/labstack/echo/v4 v4.10.2

require (
	cookhub.com/app/api/entities v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/models v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)

replace cookhub.com/app/models => ./../../../models

replace cookhub.com/app/api/entities => ./../../entities

replace cookhub.com/app/repositories => ./../../../repositories

replace cookhub.com/app/cache => ./../../../cache
