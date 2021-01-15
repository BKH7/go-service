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
	e.GET("/users/:id", h.GetByID)
	e.POST("/users", h.Store)
}

// GetByID ...
func (h *UserHandle) GetByID(c echo.Context) error {
	var user model.User
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.userUseCase.GetByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

// Store ...
func (h *UserHandle) Store(c echo.Context) error {
	var body model.User
	ctx := c.Request().Context()
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrBadParamInput)
	}
	if err := h.userUseCase.Store(ctx, &body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, domain.ErrInternalServerError)
	}
	return c.JSON(http.StatusCreated, domain.OnCreated)
}
