# gcors

简单粗暴的解决GoWeb跨域

> 使用方式：只要实现了`http.handler`接口即可轻松对接

# go默认web服务使用

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/wuwenbao/gcors"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("cors"))
    })
    
    //cors := gcors.New(mux) //默认参数为 *
    cors := gcors.New(
        mux,
        gcors.WithOrigin("*"),
        gcors.WithMethods("*"),
        gcors.WithHeaders("*"),
    )
    
    log.Fatal(http.ListenAndServe(":8080", cors))
}
```

# 第三方gorilla/mux

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/gorilla/mux"
    "github.com/wuwenbao/gcors"
)

func main() {
    router := mux.NewRouter()
    
    router.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("cors"))
    })
    
    //cors := gcors.New(mux) //默认参数为 *
    cors := gcors.NewCors(
        router,
        gcors.WithOrigin("*"),
        gcors.WithMethods("*"),
        gcors.WithHeaders("*"),
    )
    
    log.Fatal(http.ListenAndServe(":8080", cors))
}
```
