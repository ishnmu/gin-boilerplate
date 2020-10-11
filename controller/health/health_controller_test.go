package healthcontroller

import (
	"net/http/httptest"
	"gin-boilerplate/testutil"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type HealthControllerTestSuite struct {
	suite.Suite

	recorder       *httptest.ResponseRecorder
	ginContext     *gin.Context
	mockController *gomock.Controller

	healthController HealthController
}

func TestHealthController(t *testing.T) {
	suite.Run(t, new(HealthControllerTestSuite))
}

func (suite *HealthControllerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.ginContext, _ = gin.CreateTestContext(suite.recorder)
	suite.mockController = gomock.NewController(suite.T())

	suite.healthController = NewHealthController()
}

func (suite *HealthControllerTestSuite) TearDownTest() {
}

func (suite *HealthControllerTestSuite) TestHealthController() {
	suite.healthController.Check(suite.ginContext)

	expectedResponse := gin.H{
		"message": "gin-boilerplate is up and running",
	}

	suite.Equal(200, suite.recorder.Code)
	suite.Equal(testutil.MustMarshal(expectedResponse), suite.recorder.Body.String())
}
