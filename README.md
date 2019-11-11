# go-cors
简单粗暴的解决GoWeb跨域

> 如此简单的跨域被网上的一群人介绍的那么麻烦，误导新手可耻

### go默认web服务使用

```go
package main

import (
	"log"
	"net/http"

	"github.com/wuwenbao/gocors"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cors"))
	})

	log.Fatal(http.ListenAndServe(":8080", gocors.New(mux)))

}
```

### 第三方gorilla/mux

```go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wuwenbao/gocors"
)

func main() {

	router := mux.NewRouter()

	router.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("cors"))
	})

	log.Fatal(http.ListenAndServe(":8080", gocors.New(router)))
}
```
