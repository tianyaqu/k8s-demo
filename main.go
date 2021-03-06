package main

import (
	"github.com/julienschmidt/httprouter"
    "github.com/gomodule/redigo/redis"
	"net/http"
	"log"
)

func HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")

	ctx := NewContext(w, r)

	value, err := Get(key)
	if err != nil {
        log.Println(value, ' ', err)
		ctx.SetResult(1000, "network error")
	}
	ctx.SetData(value)
}

func HandlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    key := "test"
	value := ps.ByName("key")
	ctx := NewContext(w, r)

	value, err := Set(key, value)
	if err != nil {
        log.Println(value, ' ', err)
		ctx.SetResult(1000, "network error")
	}
	ctx.SetData(value)
}

var pool *redis.Pool

func main() {
    pool = NewPool()
	router := httprouter.New()
	router.GET("/home/:key", HandleGet)
	router.POST("/home/:key", HandlePost)
	log.Fatal(http.ListenAndServe(":8080", router))
}
