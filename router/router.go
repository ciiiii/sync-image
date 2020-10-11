package router

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/ciiiii/sync-image/config"
    "github.com/ciiiii/sync-image/sync"
)


func R() http.Handler {
    gin.SetMode(config.Parser().Server.Mode)
    r := gin.New()
    r.Use(gin.Logger())
    r.GET("/image", func(ctx *gin.Context) {
        image := ctx.Query("name")
        if err := sync.Sync(image); err != nil {
            ctx.JSON(http.StatusOK, gin.H{
                "success": false,
                "message": err.Error(),
            })
        } else {
            ctx.JSON(http.StatusOK, gin.H{
                "success": true,
            })
        }
    })
    return r
}