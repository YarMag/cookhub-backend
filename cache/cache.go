package cache

import (
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func init() {
	redisPool = &redis.Pool {
		MaxIdle: 10,
		MaxActive: 20,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "cookhub-redis:6379")

			if err != nil {
				panic(err)
			}

			return conn, err
		},
	}
}

func GetString(obj string, key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("HGET", obj, key))

	return value, err
}

func GetInt64(obj string, key string) (int64, error) {
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.Int64(conn.Do("HGET", obj, key))

	return value, err
}

func GetBytes(obj string, key string) ([]byte, error) {
	conn := redisPool.Get()
	defer conn.Close()
	
	value, err := redis.Bytes(conn.Do("HGET", obj, key))

	return value, err
}

func SetString(obj string, key string, value string) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", obj, key, value)

	return err
}

func SetInt64(obj string, key string, value int64) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", obj, key, value)

	return err
}

func SetBytes(obj string, key string, value []byte) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", obj, key, value)

	return err
}

func Delete(obj string, key string) error {
	conn := redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", obj, key)

	return err
}
