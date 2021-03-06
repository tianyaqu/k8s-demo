package main

import (
    "github.com/gomodule/redigo/redis"
)


func DialFunc() (redis.Conn, error){
    return redis.DialURL("redis://:@ssdb-service:8888")
}

func NewPool()* redis.Pool {
    return &redis.Pool {
        MaxIdle: 10,
        Dial : DialFunc,
    }
}

func Get(key string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

func Set(key, value string) (string, error) {
	conn := pool.Get()
	defer conn.Close()

	return redis.String(conn.Do("SET", key, value))
}
