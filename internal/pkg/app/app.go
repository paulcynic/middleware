package app

import (
    "fmt"
    "log"
    "github.com/paulcynic/middleware/internal/app/mw"
    "github.com/paulcynic/middleware/internal/app/endpoint"
    "github.com/paulcynic/middleware/internal/app/service"
    
    "github.com/labstack/echo/v4"
)

type App struct {
    e *endpoint.Endpoint
    s *service.Service
    echo *echo.Echo
}

func New() (*App, error) {
    a := &App{}

    a.s = service.New()

    a.e = endpoint.New(a.s)

    a.echo = echo.New()
    
    a.echo.Use(mw.RoleCheck)

    a.echo.GET("/status", a.e.Status)

    return a, nil
}

func (a *App) Run() error {
    fmt.Println("server is running...")
    
    err := a.echo.Start(":8087")
    if err != nil {
        log.Fatal(err)
    }
    
    return nil
}
