// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SShlykov/zeitment/bookback/internal/domain/services (interfaces: ParagraphService)
//
// Generated by this command:
//
//	mockgen -destination=../../../tests/mocks/domain/services/paragraph_service_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/services ParagraphService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/SShlykov/zeitment/bookback/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockParagraphService is a mock of ParagraphService interface.
type MockParagraphService struct {
	ctrl     *gomock.Controller
	recorder *MockParagraphServiceMockRecorder
}

// MockParagraphServiceMockRecorder is the mock recorder for MockParagraphService.
type MockParagraphServiceMockRecorder struct {
	mock *MockParagraphService
}

// NewMockParagraphService creates a new mock instance.
func NewMockParagraphService(ctrl *gomock.Controller) *MockParagraphService {
	mock := &MockParagraphService{ctrl: ctrl}
	mock.recorder = &MockParagraphServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParagraphService) EXPECT() *MockParagraphServiceMockRecorder {
	return m.recorder
}

// CreateParagraph mocks base method.
func (m *MockParagraphService) CreateParagraph(arg0 context.Context, arg1 models.CreateParagraphRequest) (*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParagraph", arg0, arg1)
	ret0, _ := ret[0].(*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParagraph indicates an expected call of CreateParagraph.
func (mr *MockParagraphServiceMockRecorder) CreateParagraph(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParagraph", reflect.TypeOf((*MockParagraphService)(nil).CreateParagraph), arg0, arg1)
}

// DeleteParagraph mocks base method.
func (m *MockParagraphService) DeleteParagraph(arg0 context.Context, arg1 string) (*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteParagraph", arg0, arg1)
	ret0, _ := ret[0].(*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteParagraph indicates an expected call of DeleteParagraph.
func (mr *MockParagraphServiceMockRecorder) DeleteParagraph(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteParagraph", reflect.TypeOf((*MockParagraphService)(nil).DeleteParagraph), arg0, arg1)
}

// GetParagraphByID mocks base method.
func (m *MockParagraphService) GetParagraphByID(arg0 context.Context, arg1 string) (*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParagraphByID", arg0, arg1)
	ret0, _ := ret[0].(*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParagraphByID indicates an expected call of GetParagraphByID.
func (mr *MockParagraphServiceMockRecorder) GetParagraphByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParagraphByID", reflect.TypeOf((*MockParagraphService)(nil).GetParagraphByID), arg0, arg1)
}

// GetParagraphsByPageID mocks base method.
func (m *MockParagraphService) GetParagraphsByPageID(arg0 context.Context, arg1 string, arg2 models.RequestParagraph) ([]*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParagraphsByPageID", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParagraphsByPageID indicates an expected call of GetParagraphsByPageID.
func (mr *MockParagraphServiceMockRecorder) GetParagraphsByPageID(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParagraphsByPageID", reflect.TypeOf((*MockParagraphService)(nil).GetParagraphsByPageID), arg0, arg1, arg2)
}

// ListParagraphs mocks base method.
func (m *MockParagraphService) ListParagraphs(arg0 context.Context, arg1 models.RequestParagraph) ([]*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListParagraphs", arg0, arg1)
	ret0, _ := ret[0].([]*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParagraphs indicates an expected call of ListParagraphs.
func (mr *MockParagraphServiceMockRecorder) ListParagraphs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParagraphs", reflect.TypeOf((*MockParagraphService)(nil).ListParagraphs), arg0, arg1)
}

// UpdateParagraph mocks base method.
func (m *MockParagraphService) UpdateParagraph(arg0 context.Context, arg1 string, arg2 models.UpdateParagraphRequest) (*models.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateParagraph", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateParagraph indicates an expected call of UpdateParagraph.
func (mr *MockParagraphServiceMockRecorder) UpdateParagraph(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateParagraph", reflect.TypeOf((*MockParagraphService)(nil).UpdateParagraph), arg0, arg1, arg2)
}
