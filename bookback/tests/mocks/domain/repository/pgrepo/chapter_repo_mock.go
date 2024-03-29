// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo (interfaces: ChapterRepo)
//
// Generated by this command:
//
//	mockgen -destination=../../../../tests/mocks/domain/repository/pgrepo/chapter_repo_mock.go -package=mocks github.com/SShlykov/zeitment/bookback/internal/domain/repository/pgrepo ChapterRepo
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

// MockChapterRepo is a mock of ChapterRepo interface.
type MockChapterRepo struct {
	ctrl     *gomock.Controller
	recorder *MockChapterRepoMockRecorder
}

// MockChapterRepoMockRecorder is the mock recorder for MockChapterRepo.
type MockChapterRepoMockRecorder struct {
	mock *MockChapterRepo
}

// NewMockChapterRepo creates a new mock instance.
func NewMockChapterRepo(ctrl *gomock.Controller) *MockChapterRepo {
	mock := &MockChapterRepo{ctrl: ctrl}
	mock.recorder = &MockChapterRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChapterRepo) EXPECT() *MockChapterRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChapterRepo) Create(arg0 context.Context, arg1 *entity.Chapter) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChapterRepoMockRecorder) Create(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChapterRepo)(nil).Create), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockChapterRepo) FindByID(arg0 context.Context, arg1 string) (*entity.Chapter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.Chapter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockChapterRepoMockRecorder) FindByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockChapterRepo)(nil).FindByID), arg0, arg1)
}

// FindByKV mocks base method.
func (m *MockChapterRepo) FindByKV(arg0 context.Context, arg1 dbutils.QueryOptions) ([]*entity.Chapter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByKV", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Chapter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKV indicates an expected call of FindByKV.
func (mr *MockChapterRepoMockRecorder) FindByKV(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKV", reflect.TypeOf((*MockChapterRepo)(nil).FindByKV), arg0, arg1)
}

// HardDelete mocks base method.
func (m *MockChapterRepo) HardDelete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDelete indicates an expected call of HardDelete.
func (mr *MockChapterRepoMockRecorder) HardDelete(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDelete", reflect.TypeOf((*MockChapterRepo)(nil).HardDelete), arg0, arg1)
}

// List mocks base method.
func (m *MockChapterRepo) List(arg0 context.Context, arg1 dbutils.Pagination) ([]*entity.Chapter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*entity.Chapter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockChapterRepoMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockChapterRepo)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockChapterRepo) Update(arg0 context.Context, arg1 string, arg2 *entity.Chapter) (*entity.Chapter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*entity.Chapter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockChapterRepoMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockChapterRepo)(nil).Update), arg0, arg1, arg2)
}
