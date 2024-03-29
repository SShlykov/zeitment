// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo (interfaces: ParagraphRepo)
//
// Generated by this command:
//
//	mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/paragraph_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo ParagraphRepo
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	"github.com/SShlykov/zeitment/bookback/internal/infrastructure/repository/entity"
	"github.com/SShlykov/zeitment/postgres/dbutils"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockParagraphRepo is a mock of ParagraphRepo interface.
type MockParagraphRepo struct {
	ctrl     *gomock.Controller
	recorder *MockParagraphRepoMockRecorder
}

// MockParagraphRepoMockRecorder is the mock recorder for MockParagraphRepo.
type MockParagraphRepoMockRecorder struct {
	mock *MockParagraphRepo
}

// NewMockParagraphRepo creates a new mock instance.
func NewMockParagraphRepo(ctrl *gomock.Controller) *MockParagraphRepo {
	mock := &MockParagraphRepo{ctrl: ctrl}
	mock.recorder = &MockParagraphRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParagraphRepo) EXPECT() *MockParagraphRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockParagraphRepo) Create(arg0 context.Context, arg1 *entity.Paragraph) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockParagraphRepoMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockParagraphRepo)(nil).Create), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockParagraphRepo) FindByID(arg0 context.Context, arg1 string) (*entity.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockParagraphRepoMockRecorder) FindByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockParagraphRepo)(nil).FindByID), arg0, arg1)
}

// FindByKV mocks base method.
func (m *MockParagraphRepo) FindByKV(arg0 context.Context, arg1 dbutils.QueryOptions) ([]*entity.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKV", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKV indicates an expected call of FindByKV.
func (mr *MockParagraphRepoMockRecorder) FindByKV(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKV", reflect.TypeOf((*MockParagraphRepo)(nil).FindByKV), arg0, arg1)
}

// HardDelete mocks base method.
func (m *MockParagraphRepo) HardDelete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDelete indicates an expected call of HardDelete.
func (mr *MockParagraphRepoMockRecorder) HardDelete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDelete", reflect.TypeOf((*MockParagraphRepo)(nil).HardDelete), arg0, arg1)
}

// List mocks base method.
func (m *MockParagraphRepo) List(arg0 context.Context, arg1 dbutils.Pagination) ([]*entity.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockParagraphRepoMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockParagraphRepo)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockParagraphRepo) Update(arg0 context.Context, arg1 string, arg2 *entity.Paragraph) (*entity.Paragraph, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Paragraph)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockParagraphRepoMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockParagraphRepo)(nil).Update), arg0, arg1, arg2)
}
