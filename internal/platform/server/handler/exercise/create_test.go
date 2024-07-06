package exercise

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/huandu/go-assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/wodm8/wodm8-core/internal/application"
	"github.com/wodm8/wodm8-core/internal/platform/storage/storagemocks"
)

func TestHandler_create(t *testing.T) {
	exerciseRepository := new(storagemocks.ExerciseRepository)
	exerciseRepository.On("Save", mock.Anything, mock.Anything).Return(nil)

	exerciseServer := application.NewExerciseService(exerciseRepository)

	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.POST("/api/v1/exercises", CreateHandler(exerciseServer))

	t.Run("given an invalid request it return 400", func(t *testing.T) {
		createExerciseReq := createRequest{}

		b, err := json.Marshal(createExerciseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/api/v1/exercises", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})
}
