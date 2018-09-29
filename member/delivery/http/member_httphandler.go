package http

import (
	"fmt"
	"net/http"

	"github.com/ibadsatria/ucantplayalone/member"
	_ucaseMember "github.com/ibadsatria/ucantplayalone/member/usecase"
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type memberHTTPHandler struct {
	ucaseMember _ucaseMember.MemberUsecase
}

// NewMemberHTTPHandler init new instance
func NewMemberHTTPHandler(e *echo.Echo, u _ucaseMember.MemberUsecase) {
	handler := memberHTTPHandler{ucaseMember: u}

	fmt.Println("member HTTP handler set up")

	e.POST("/member", handler.Store)
}

func (h *memberHTTPHandler) Store(c echo.Context) error {
	var member member.Member
	err := c.Bind(&member)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if _, err := isRequestValid(&member); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newMember, err := h.ucaseMember.Store(&member)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newMember)
}

func isRequestValid(m *member.Member) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
