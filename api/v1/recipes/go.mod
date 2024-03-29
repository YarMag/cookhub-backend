module cookhub.com/app/api/v1/recipes

go 1.17

require github.com/labstack/echo/v4 v4.6.3

require (
	cookhub.com/app/api/entities v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/models v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/repositories v0.0.0-00010101000000-000000000000 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace cookhub.com/app/models => ./../../../models

replace cookhub.com/app/api/entities => ./../../entities

replace cookhub.com/app/repositories => ./../../../repositories

replace cookhub.com/app/cache => ./../../../cache
