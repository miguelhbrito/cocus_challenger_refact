package integrationtests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cocus_challenger_refact/business/data/login"
	"github.com/cocus_challenger_refact/business/integration_tests/schemas"
	"github.com/cocus_challenger_refact/business/integration_tests/utilities"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/suite"
)

type LoginSuite struct {
	suite.Suite
}

func TestLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginSuite))
}

func (suite *LoginSuite) SetupSuite() {
	utilities.Setup()
}

func (s *LoginSuite) TestLoginSuccess() {
	login := login.NewLogin{Username: utilities.User, Password: utilities.Password}
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/login/create").
		WithHeader("Content-Type", "application/json").
		WithJSON(login).
		Expect().
		Status(http.StatusCreated)
}

func (s *LoginSuite) TestLoginFail() {
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/login/create").WithJSON(nil).
		Expect().
		Status(http.StatusBadRequest).JSON()
}

func (s *LoginSuite) TestLoginTokenSuccess() {
	utilities.CreateUser()
	loginReq := login.NewLogin{Username: utilities.User, Password: utilities.Password}
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/login").
		WithHeader("Content-Type", "application/json").
		WithJSON(loginReq).
		Expect().
		Status(http.StatusOK).JSON().Schema(schemas.LoginSuccessSchema())
}

func (s *LoginSuite) TestLoginTokenFail() {
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/login").
		WithHeader("Content-Type", "application/json").
		WithJSON(nil).
		Expect().
		Status(http.StatusBadRequest)
}
