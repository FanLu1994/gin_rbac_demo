package main

import (
	"fmt"
	"gin_rbac/router"
	"net/http"
	"time"
)

func main() {
	myrouter := router.InitRouters()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",8000),
		Handler: myrouter,
		ReadTimeout: 60*time.Second,
		WriteTimeout: 60*time.Second,
		MaxHeaderBytes: 1<<20,
	}

	//myrouter.Run("0.0.0.0:8080")
	_ = s.ListenAndServe()


}
