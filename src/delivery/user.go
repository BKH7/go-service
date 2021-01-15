package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/bunlert274/go-service/src/domain"
	"gitlab.com/bunlert274/go-service/src/model"
)

// UserHandle ...
type UserHandle struct {
	userUseCase domain.UserUseCase
}

// NewUserHandle ...
func NewUserHandle(userUseCase domain.UserUseCase) *UserHandle {
	return &UserHandle{userUseCase}
}

// GetUser ...
func (h *UserHandle) GetUser(c echo.Context) error {
	var user model.User
	id := c.Param("id")

	err := h.userUseCase.GetUser(&user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	return c.JSON(http.StatusOK, user)

}
