package main

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	mux *echo.Echo
	db  *pgxpool.Pool
	l   echo.Logger
	err error
}

func main() {
	s := &server{}
	s.registerRoutes()
	ctx := context.Background()
	s.db, s.err = newPGconn(ctx)
	if s.err != nil {
		s.l.Fatal("Unable to connection to database or: ", s.err)
	}
	s.l.Printf("DB Connected!")

	// Start server
	s.l.Fatal(s.mux.Start(":1323"))
}

func (s *server) registerRoutes() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.Static("/", "web")
	e.GET("/users", s.showAllUsers)
	e.POST("/users", s.createNewUser)
	e.PUT("/users", s.updateUser)
	e.DELETE("/users", s.deleteUser)
	s.mux = e
	s.l = e.Logger

}
