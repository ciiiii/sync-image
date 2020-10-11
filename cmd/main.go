package main

import (
    "fmt"
    "net/http"

    "github.com/ciiiii/sync-image/config"
    "github.com/ciiiii/sync-image/router"
)

func main()  {
    if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Parser().Server.Port), router.R()); err != nil {
        panic(err)
    }
}
