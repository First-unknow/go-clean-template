package http

import (
	"net/http"

	"innovasive/go-clean-template/middleware"
	models "innovasive/go-clean-template/models"
	"innovasive/go-clean-template/service/user"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUs user.UserUsecaseInf
}

func NewUserHandler(e *echo.Echo, middL *middleware.GoMiddleware, us user.UserUsecaseInf) {
	handler := &userHandler{
		userUs: us,
	}
	e.GET("/users", handler.FetchAll)
	e.GET("/users/:email", handler.FindByMail)
	e.POST("/users", handler.CreateUser)
}

func (u *userHandler) FetchAll(c echo.Context) error {
	users, err := u.userUs.FetchAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	responseData := map[string]interface{}{
		"users": users,
	}
	return c.JSON(http.StatusOK, responseData)
}

func (u *userHandler) FindByMail(c echo.Context) error {
	email := c.Param("email")
	users, err := u.userUs.FindByMail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	responseData := map[string]interface{}{
		"users": users,
	}
	return c.JSON(http.StatusOK, responseData)
}

func (u *userHandler) CreateUser(c echo.Context) error {
	user := new(models.User)
	resJSON := make(map[string]uuid.UUID)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusNotAcceptable, "Not Acceptable")
	}

	userID, err := u.userUs.CreateUser(user)
	if err != nil || userID == uuid.Nil {
		return c.JSON(http.StatusNotAcceptable, "Not Acceptable")
	}
	resJSON["userID"] = userID
	return c.JSON(http.StatusCreated, resJSON)
}
