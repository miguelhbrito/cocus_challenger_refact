package integrationtests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cocus_challenger_refact/business/data/triangle"
	"github.com/cocus_challenger_refact/business/integration_tests/schemas"
	"github.com/cocus_challenger_refact/business/integration_tests/utilities"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/suite"
)

type TriangleSuite struct {
	suite.Suite
}

func TestTriangleSuite(t *testing.T) {
	suite.Run(t, new(TriangleSuite))
}

func (suite *TriangleSuite) SetupSuite() {
	utilities.Setup()
	if utilities.Token == "" {
		utilities.CreateUser()
		utilities.GenerateToken()
	}
}

func (s *TriangleSuite) TestTriangleCreateSuccess() {
	triangleStructs := triangle.NewTriangle{
		Side1: 10,
		Side2: 10,
		Side3: 10,
	}
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/triangles").
		WithHeader("Content-Type", "application/json").
		WithHeader(utilities.AuhtorizationHeader, utilities.Token).
		WithJSON(triangleStructs).
		Expect().
		Status(http.StatusCreated).JSON().Schema(schemas.TrianglesSchema())
}

func (s *TriangleSuite) TestTriangleCreateFail() {
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/triangles").
		WithHeader("Content-Type", "application/json").
		WithHeader(utilities.AuhtorizationHeader, utilities.Token).
		WithJSON(nil).
		Expect().
		Status(http.StatusBadRequest)
}

func (s *TriangleSuite) TestTriangleCreateFailAuth() {
	triangleStructs := triangle.NewTriangle{
		Side1: 10,
		Side2: 10,
		Side3: 10,
	}
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.POST("/triangles").
		WithHeader("Content-Type", "application/json").
		WithHeader(utilities.AuhtorizationHeader, "error to login with invalid token").
		WithJSON(triangleStructs).
		Expect().
		Status(http.StatusInternalServerError)
}

func (s *TriangleSuite) TestTriangleListAllTriangles() {
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.GET("/triangles").
		WithHeader(utilities.AuhtorizationHeader, utilities.Token).
		Expect().
		Status(http.StatusOK).JSON().Schema(schemas.ListTrianglesSchema())
}

func (s *TriangleSuite) TestTriangleListFailAuth() {
	e := httpexpect.New(s.T(), fmt.Sprintf("http://"+utilities.BaseURL))

	e.GET("/triangles").
		Expect().
		Status(http.StatusUnauthorized)
}
