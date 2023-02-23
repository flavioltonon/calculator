package router

import "net/http"

var NopHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
