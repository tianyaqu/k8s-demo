package main

import (
	"net/http"
	"encoding/json"
)

type Result struct {
	Retcode int  `json:"retcode"`
	Msg     string `json:"msg"`
	Data    interface{}  `json:"data"`
}

type Context struct {
	w http.ResponseWriter
	r *http.Request

	result *Result
}

func NewContext (w http.ResponseWriter, r *http.Request) (*Context) {
	return &Context{
		w : w,
		r : r,
		result : &Result{
			Retcode: 0,
			Msg : "ok",
		},
	}
}

func (ctx *Context) SetResult(ret int, msg string ) {
	ctx.result.Retcode = ret
	ctx.result.Msg = msg
}


func (ctx *Context) SetData(v interface{}) {
	ctx.w.Header().Set("Content-Type", "application/json")

	ctx.result.Data = v
	json.NewEncoder(ctx.w).Encode(ctx.result)
}

