// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/rss_feed.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "portfolio-backend/domain"
	reflect "reflect"

	events "github.com/aws/aws-lambda-go/events"
	gomock "github.com/golang/mock/gomock"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

// MockRSSFeedRepository is a mock of RSSFeedRepository interface.
type MockRSSFeedRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRSSFeedRepositoryMockRecorder
}

// MockRSSFeedRepositoryMockRecorder is the mock recorder for MockRSSFeedRepository.
type MockRSSFeedRepositoryMockRecorder struct {
	mock *MockRSSFeedRepository
}

// NewMockRSSFeedRepository creates a new mock instance.
func NewMockRSSFeedRepository(ctrl *gomock.Controller) *MockRSSFeedRepository {
	mock := &MockRSSFeedRepository{ctrl: ctrl}
	mock.recorder = &MockRSSFeedRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRSSFeedRepository) EXPECT() *MockRSSFeedRepositoryMockRecorder {
	return m.recorder
}

// CreateRSSFeed mocks base method.
func (m *MockRSSFeedRepository) CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRSSFeed", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRSSFeed indicates an expected call of CreateRSSFeed.
func (mr *MockRSSFeedRepositoryMockRecorder) CreateRSSFeed(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRSSFeed", reflect.TypeOf((*MockRSSFeedRepository)(nil).CreateRSSFeed), ctx, input)
}

// DeleteRSSFeed mocks base method.
func (m *MockRSSFeedRepository) DeleteRSSFeed(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRSSFeed", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRSSFeed indicates an expected call of DeleteRSSFeed.
func (mr *MockRSSFeedRepositoryMockRecorder) DeleteRSSFeed(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRSSFeed", reflect.TypeOf((*MockRSSFeedRepository)(nil).DeleteRSSFeed), ctx, id)
}

// GetRSSFeed mocks base method.
func (m *MockRSSFeedRepository) GetRSSFeed(ctx context.Context, id string) (*domain.RSSFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRSSFeed", ctx, id)
	ret0, _ := ret[0].(*domain.RSSFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRSSFeed indicates an expected call of GetRSSFeed.
func (mr *MockRSSFeedRepositoryMockRecorder) GetRSSFeed(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRSSFeed", reflect.TypeOf((*MockRSSFeedRepository)(nil).GetRSSFeed), ctx, id)
}

// GetRSSFeeds mocks base method.
func (m *MockRSSFeedRepository) GetRSSFeeds(ctx context.Context) ([]domain.RSSFeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRSSFeeds", ctx)
	ret0, _ := ret[0].([]domain.RSSFeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRSSFeeds indicates an expected call of GetRSSFeeds.
func (mr *MockRSSFeedRepositoryMockRecorder) GetRSSFeeds(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRSSFeeds", reflect.TypeOf((*MockRSSFeedRepository)(nil).GetRSSFeeds), ctx)
}

// IsExistsUrl mocks base method.
func (m *MockRSSFeedRepository) IsExistsUrl(ctx context.Context, url string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistsUrl", ctx, url)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsExistsUrl indicates an expected call of IsExistsUrl.
func (mr *MockRSSFeedRepositoryMockRecorder) IsExistsUrl(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistsUrl", reflect.TypeOf((*MockRSSFeedRepository)(nil).IsExistsUrl), ctx, url)
}

// MockRSSFeedHandler is a mock of RSSFeedHandler interface.
type MockRSSFeedHandler struct {
	ctrl     *gomock.Controller
	recorder *MockRSSFeedHandlerMockRecorder
}

// MockRSSFeedHandlerMockRecorder is the mock recorder for MockRSSFeedHandler.
type MockRSSFeedHandlerMockRecorder struct {
	mock *MockRSSFeedHandler
}

// NewMockRSSFeedHandler creates a new mock instance.
func NewMockRSSFeedHandler(ctrl *gomock.Controller) *MockRSSFeedHandler {
	mock := &MockRSSFeedHandler{ctrl: ctrl}
	mock.recorder = &MockRSSFeedHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRSSFeedHandler) EXPECT() *MockRSSFeedHandlerMockRecorder {
	return m.recorder
}

// BatchGetRSSFeeds mocks base method.
func (m *MockRSSFeedHandler) BatchGetRSSFeeds(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGetRSSFeeds", request)
	ret0, _ := ret[0].(events.APIGatewayProxyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetRSSFeeds indicates an expected call of BatchGetRSSFeeds.
func (mr *MockRSSFeedHandlerMockRecorder) BatchGetRSSFeeds(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetRSSFeeds", reflect.TypeOf((*MockRSSFeedHandler)(nil).BatchGetRSSFeeds), request)
}

// CreateRSSFeed mocks base method.
func (m *MockRSSFeedHandler) CreateRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRSSFeed", request)
	ret0, _ := ret[0].(events.APIGatewayProxyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRSSFeed indicates an expected call of CreateRSSFeed.
func (mr *MockRSSFeedHandlerMockRecorder) CreateRSSFeed(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRSSFeed", reflect.TypeOf((*MockRSSFeedHandler)(nil).CreateRSSFeed), request)
}

// DeleteRSSFeed mocks base method.
func (m *MockRSSFeedHandler) DeleteRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRSSFeed", request)
	ret0, _ := ret[0].(events.APIGatewayProxyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRSSFeed indicates an expected call of DeleteRSSFeed.
func (mr *MockRSSFeedHandlerMockRecorder) DeleteRSSFeed(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRSSFeed", reflect.TypeOf((*MockRSSFeedHandler)(nil).DeleteRSSFeed), request)
}

// GetRSSFeed mocks base method.
func (m *MockRSSFeedHandler) GetRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRSSFeed", request)
	ret0, _ := ret[0].(events.APIGatewayProxyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRSSFeed indicates an expected call of GetRSSFeed.
func (mr *MockRSSFeedHandlerMockRecorder) GetRSSFeed(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRSSFeed", reflect.TypeOf((*MockRSSFeedHandler)(nil).GetRSSFeed), request)
}
