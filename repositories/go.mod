module cookhub.com/app/repositories

go 1.17

replace cookhub.com/app/cache => ./../cache

replace cookhub.com/app/models => ./../models

require (
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	cookhub.com/app/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
)
