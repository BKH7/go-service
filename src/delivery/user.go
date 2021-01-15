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
func NewUserHandle(e *echo.Echo, userUseCase domain.UserUseCase) {
	h := &UserHandle{userUseCase}
	e.GET("/users", h.GetByID)
}

// GetByID ...
func (h *UserHandle) GetByID(c echo.Context) error {
	var user *model.User
	ctx := c.Request().Context()
	id := c.Param("id")

	if err := h.userUseCase.GetByID(ctx, user, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}
