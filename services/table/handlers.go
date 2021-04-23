package main

import (
	"fmt"
	"init/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// showAllUsers - handler на показ всех пользователей
func (s *server) showAllUsers(c echo.Context) error {
	uu := models.Users{}
	uu, s.err = getAllUsers(c.Request().Context(), s.db, uu)
	if s.err != nil {
		return s.err
	}
	return c.JSON(http.StatusOK, uu)
}

// createNewUser - на создание нового пользователя
func (s *server) createNewUser(c echo.Context) error {
	u := &models.User{}
	if s.err = c.Bind(u); s.err != nil {
		s.l.Errorf("bad request %v", s.err)
		return c.JSON(http.StatusBadRequest, s.err)
	}
	u, s.err = insertUser(c.Request().Context(), s.db, u)
	if s.err != nil {
		s.l.Errorf("createUser error %v", s.err)
		return c.JSON(http.StatusInternalServerError, s.err)
	}
	return c.JSON(http.StatusOK, u)

}

func (s *server) updateUser(c echo.Context) error {
	u := &models.User{}
	if s.err = c.Bind(u); s.err != nil {
		s.l.Errorf("bad request %v", s.err)
		return c.JSON(http.StatusBadRequest, s.err)
	}
	s.err = updateUser(c.Request().Context(), s.db, u)
	if s.err != nil {
		s.l.Errorf("updateUser error %v", s.err)
		return c.JSON(http.StatusInternalServerError, s.err)
	}
	return c.JSON(http.StatusOK, "OK")
}

// delete user - handler на показ одного пользователя по user_id
func (s *server) deleteUser(c echo.Context) error {
	u := &models.User{}
	if s.err = c.Bind(u); s.err != nil {
		s.l.Errorf("bad request %v", s.err)
		return c.JSON(http.StatusBadRequest, s.err)
	}
	fmt.Printf("\n%v\n", u)
	s.err = deleteUserByID(c.Request().Context(), s.db, u)
	if s.err != nil {
		s.l.Errorf("delete user error %v", s.err)
		return c.JSON(http.StatusInternalServerError, s.err)
	}
	return c.JSON(http.StatusOK, "OK")
}
