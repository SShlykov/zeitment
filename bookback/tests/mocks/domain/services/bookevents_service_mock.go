// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/services/bookevents_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/SShlykov/zeitment/bookback/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockBookEventsService is a mock of BookEventsService interface.
type MockBookEventsService struct {
	ctrl     *gomock.Controller
	recorder *MockBookEventsServiceMockRecorder
}

// MockBookEventsServiceMockRecorder is the mock recorder for MockBookEventsService.
type MockBookEventsServiceMockRecorder struct {
	mock *MockBookEventsService
}

// NewMockBookEventsService creates a new mock instance.
func NewMockBookEventsService(ctrl *gomock.Controller) *MockBookEventsService {
	mock := &MockBookEventsService{ctrl: ctrl}
	mock.recorder = &MockBookEventsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookEventsService) EXPECT() *MockBookEventsServiceMockRecorder {
	return m.recorder
}

// CreateBookEvent mocks base method.
func (m *MockBookEventsService) CreateBookEvent(ctx context.Context, request models.CreateBookEventRequest) (*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBookEvent", ctx, request)
	ret0, _ := ret[0].(*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBookEvent indicates an expected call of CreateBookEvent.
func (mr *MockBookEventsServiceMockRecorder) CreateBookEvent(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBookEvent", reflect.TypeOf((*MockBookEventsService)(nil).CreateBookEvent), ctx, request)
}

// DeleteBookEvent mocks base method.
func (m *MockBookEventsService) DeleteBookEvent(ctx context.Context, id string) (*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookEvent", ctx, id)
	ret0, _ := ret[0].(*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBookEvent indicates an expected call of DeleteBookEvent.
func (mr *MockBookEventsServiceMockRecorder) DeleteBookEvent(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookEvent", reflect.TypeOf((*MockBookEventsService)(nil).DeleteBookEvent), ctx, id)
}

// GetBookEventByID mocks base method.
func (m *MockBookEventsService) GetBookEventByID(ctx context.Context, id string) (*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookEventByID", ctx, id)
	ret0, _ := ret[0].(*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookEventByID indicates an expected call of GetBookEventByID.
func (mr *MockBookEventsServiceMockRecorder) GetBookEventByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookEventByID", reflect.TypeOf((*MockBookEventsService)(nil).GetBookEventByID), ctx, id)
}

// GetBookEventsByBookID mocks base method.
func (m *MockBookEventsService) GetBookEventsByBookID(ctx context.Context, bookID string, request models.RequestBookEvent) ([]*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookEventsByBookID", ctx, bookID, request)
	ret0, _ := ret[0].([]*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookEventsByBookID indicates an expected call of GetBookEventsByBookID.
func (mr *MockBookEventsServiceMockRecorder) GetBookEventsByBookID(ctx, bookID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookEventsByBookID", reflect.TypeOf((*MockBookEventsService)(nil).GetBookEventsByBookID), ctx, bookID, request)
}

// GetBookEventsByChapterID mocks base method.
func (m *MockBookEventsService) GetBookEventsByChapterID(ctx context.Context, chapterID string, request models.RequestBookEvent) ([]*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookEventsByChapterID", ctx, chapterID, request)
	ret0, _ := ret[0].([]*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookEventsByChapterID indicates an expected call of GetBookEventsByChapterID.
func (mr *MockBookEventsServiceMockRecorder) GetBookEventsByChapterID(ctx, chapterID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookEventsByChapterID", reflect.TypeOf((*MockBookEventsService)(nil).GetBookEventsByChapterID), ctx, chapterID, request)
}

// GetBookEventsByPageID mocks base method.
func (m *MockBookEventsService) GetBookEventsByPageID(ctx context.Context, pageID string, request models.RequestBookEvent) ([]*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookEventsByPageID", ctx, pageID, request)
	ret0, _ := ret[0].([]*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookEventsByPageID indicates an expected call of GetBookEventsByPageID.
func (mr *MockBookEventsServiceMockRecorder) GetBookEventsByPageID(ctx, pageID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookEventsByPageID", reflect.TypeOf((*MockBookEventsService)(nil).GetBookEventsByPageID), ctx, pageID, request)
}

// GetBookEventsByParagraphID mocks base method.
func (m *MockBookEventsService) GetBookEventsByParagraphID(ctx context.Context, paragraphID string, request models.RequestBookEvent) ([]*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookEventsByParagraphID", ctx, paragraphID, request)
	ret0, _ := ret[0].([]*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookEventsByParagraphID indicates an expected call of GetBookEventsByParagraphID.
func (mr *MockBookEventsServiceMockRecorder) GetBookEventsByParagraphID(ctx, paragraphID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookEventsByParagraphID", reflect.TypeOf((*MockBookEventsService)(nil).GetBookEventsByParagraphID), ctx, paragraphID, request)
}

// UpdateBookEvent mocks base method.
func (m *MockBookEventsService) UpdateBookEvent(ctx context.Context, id string, request models.UpdateBookEventRequest) (*models.BookEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBookEvent", ctx, id, request)
	ret0, _ := ret[0].(*models.BookEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBookEvent indicates an expected call of UpdateBookEvent.
func (mr *MockBookEventsServiceMockRecorder) UpdateBookEvent(ctx, id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBookEvent", reflect.TypeOf((*MockBookEventsService)(nil).UpdateBookEvent), ctx, id, request)
}
