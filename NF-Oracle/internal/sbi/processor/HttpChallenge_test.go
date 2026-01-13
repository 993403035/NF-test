package processor_test

import (
	"net/http/httptest"
	"testing"

	"github.com/andy89923/nf-example/internal/sbi/processor"
	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

func Test_Httpchallenge(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCtrl := gomock.NewController(t)
	processorNf := processor.NewMockProcessorNf(mockCtrl)
	processor, err := processor.NewProcessor(processorNf)
	if err != nil {
		t.Errorf("Failed to create processor: %s", err)
		return
	}

	t.Run("Find Character That Exists", func(t *testing.T) {
		const INPUT_NAME = "www.oracletest.com"
		const EXPECTED_STATUS = 200
		const EXPECTED_BODY = "request oracle to challenge: " + INPUT_NAME

		//processorNf.EXPECT().Context().Return(&nf_context.NFContext{
		//	SpyFamilyData: map[string]string{
		//		"Anya": "Forger",
		//	},
		//})

		httpRecorder := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(httpRecorder)
		processor.ChallengeHttp(ginCtx, INPUT_NAME)

		if httpRecorder.Code != EXPECTED_STATUS {
			t.Errorf("Expected status code %d, got %d", EXPECTED_STATUS, httpRecorder.Code)
		}

		if httpRecorder.Body.String() != EXPECTED_BODY {
			t.Errorf("Expected body %s, got %s", EXPECTED_BODY, httpRecorder.Body.String())
		}
	})

}
