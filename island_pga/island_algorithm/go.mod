module island_algorithm

go 1.17

replace ga_operators => ./../ga_operators

replace cookhub.com/app/cache => ./../../cache

require (
	ga_operators v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
)

require (
	cookhub.com/app/cache v0.0.0-00010101000000-000000000000 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
)
