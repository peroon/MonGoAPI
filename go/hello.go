package main

import (
    "fmt"
    //"os/exec"
    
    "net/http"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func main() {
    var url = "http://localhost:1323/\n"
    fmt.Printf(url)
    //exec.Command("open " + url)

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.GET("/user", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, user")
    })
    
    e.Run(standard.New(":1323"))
}