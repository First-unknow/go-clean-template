package http

import (
	"net/http"

	"innovasive/go-clean-template/middleware"
	"innovasive/go-clean-template/service/user"

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
