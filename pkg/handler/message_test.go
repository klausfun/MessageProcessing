package handler

import (
	"MessageProcessing/models"
	"MessageProcessing/pkg/service"
	mock_service "MessageProcessing/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createMessage(t *testing.T) {
	type mockBehavior func(s *mock_service.MockMessage, input models.Message)

	testTable := []struct {
		name                string
		inputBody           string
		inputMessage        models.Message
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"content":"content"}`,
			inputMessage: models.Message{
				Content: "content",
			},
			mockBehavior: func(s *mock_service.MockMessage, input models.Message) {
				s.EXPECT().Create(input).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty Fields",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_service.MockMessage, input models.Message) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"content":"content"}`,
			inputMessage: models.Message{
				Content: "content",
			},
			mockBehavior: func(s *mock_service.MockMessage, input models.Message) {
				s.EXPECT().Create(input).Return(0, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			message := mock_service.NewMockMessage(c)
			testCase.mockBehavior(message, testCase.inputMessage)

			services := &service.Service{Message: message}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/api/message", func(c *gin.Context) {
				handler.createMessage(c)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/message", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_getAllMessages(t *testing.T) {
	type mockBehaviorCur func(s *mock_service.MockMessage)
	type mockBehaviorComp func(s *mock_service.MockMessage)

	testTable := []struct {
		name                string
		mockBehaviorCur     mockBehaviorCur
		mockBehaviorComp    mockBehaviorComp
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehaviorCur: func(s *mock_service.MockMessage) {
				s.EXPECT().GetCurMessages().Return([]models.Message{
					{
						Id:      1,
						Content: "Content1",
					},
				}, nil)
			},
			mockBehaviorComp: func(s *mock_service.MockMessage) {
				s.EXPECT().GetCompMessages().Return([]models.Message{
					{
						Id:      2,
						Content: "Content2",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"number_of_current_messages":1,"current":[{"id":1,"content":"Content1"}],"number_of_completed_messages":1,"completed":[{"id":2,"content":"Content2"}]}`,
		},
		{
			name: "Service Failure",
			mockBehaviorCur: func(s *mock_service.MockMessage) {
				s.EXPECT().GetCurMessages().Return(nil, errors.New("service failure"))
			},
			mockBehaviorComp:    func(s *mock_service.MockMessage) {},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
		{
			name:            "Service Failure",
			mockBehaviorCur: func(s *mock_service.MockMessage) {},
			mockBehaviorComp: func(s *mock_service.MockMessage) {
				s.EXPECT().GetCurMessages().Return(nil, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			message := mock_service.NewMockMessage(c)
			testCase.mockBehaviorCur(message)
			testCase.mockBehaviorComp(message)

			services := &service.Service{Message: message}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/api/message", func(c *gin.Context) {
				handler.getAllMessages(c)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/message", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
