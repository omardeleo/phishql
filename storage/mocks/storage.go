// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jloom6/phishql/storage (interfaces: Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	structs "github.com/jloom6/phishql/structs"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// GetArtists mocks base method
func (m *MockInterface) GetArtists(arg0 context.Context, arg1 structs.GetArtistsRequest) ([]structs.Artist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArtists", arg0, arg1)
	ret0, _ := ret[0].([]structs.Artist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArtists indicates an expected call of GetArtists
func (mr *MockInterfaceMockRecorder) GetArtists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArtists", reflect.TypeOf((*MockInterface)(nil).GetArtists), arg0, arg1)
}

// GetShows mocks base method
func (m *MockInterface) GetShows(arg0 context.Context, arg1 structs.GetShowsRequest) ([]structs.Show, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShows", arg0, arg1)
	ret0, _ := ret[0].([]structs.Show)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShows indicates an expected call of GetShows
func (mr *MockInterfaceMockRecorder) GetShows(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShows", reflect.TypeOf((*MockInterface)(nil).GetShows), arg0, arg1)
}

// GetSongs mocks base method
func (m *MockInterface) GetSongs(arg0 context.Context, arg1 structs.GetSongsRequest) ([]structs.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSongs", arg0, arg1)
	ret0, _ := ret[0].([]structs.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSongs indicates an expected call of GetSongs
func (mr *MockInterfaceMockRecorder) GetSongs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSongs", reflect.TypeOf((*MockInterface)(nil).GetSongs), arg0, arg1)
}

// GetTags mocks base method
func (m *MockInterface) GetTags(arg0 context.Context, arg1 structs.GetTagsRequest) ([]structs.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags", arg0, arg1)
	ret0, _ := ret[0].([]structs.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags
func (mr *MockInterfaceMockRecorder) GetTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockInterface)(nil).GetTags), arg0, arg1)
}
