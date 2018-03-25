package utils

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

type myFileHandler struct {
	h http.Handler
}

func muxVariableLookup(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func IndexNameLookup(req *http.Request) string {
	fmt.Println("IndexNameLookup")
	return muxVariableLookup(req, "indexName") // 通过request中的indexName来查找
}