package mock_repository

import (
	reflect "reflect"
	kvado "taskserver/pkg/grpc/proto"

	gomock "github.com/golang/mock/gomock"
)

// MockBooksAuthors is a mock of BooksAuthors interface.
type MockBooksAuthors struct {
	ctrl     *gomock.Controller
	recorder *MockBooksAuthorsMockRecorder
}

// MockBooksAuthorsMockRecorder is the mock recorder for MockBooksAuthors.
type MockBooksAuthorsMockRecorder struct {
	mock *MockBooksAuthors
}

// NewMockBooksAuthors creates a new mock instance.
func NewMockBooksAuthors(ctrl *gomock.Controller) *MockBooksAuthors {
	mock := &MockBooksAuthors{ctrl: ctrl}
	mock.recorder = &MockBooksAuthorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBooksAuthors) EXPECT() *MockBooksAuthorsMockRecorder {
	return m.recorder
}

// FindAuthorByBookName mocks base method.
func (m *MockBooksAuthors) FindAuthorByBookName(arg0 string) ([]*kvado.FindAuthorsResponse_Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAuthorByBookName", arg0)
	ret0, _ := ret[0].([]*kvado.FindAuthorsResponse_Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthorByBookName indicates an expected call of FindAuthorByBookName.
func (mr *MockBooksAuthorsMockRecorder) FindAuthorByBookName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthorByBookName", reflect.TypeOf((*MockBooksAuthors)(nil).FindAuthorByBookName), arg0)
}

// FindAuthorsByBookID mocks base method.
func (m *MockBooksAuthors) FindAuthorsByBookID(arg0 int) ([]*kvado.FindAuthorsResponse_Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAuthorsByBookID", arg0)
	ret0, _ := ret[0].([]*kvado.FindAuthorsResponse_Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthorsByBookID indicates an expected call of FindAuthorsByBookID.
func (mr *MockBooksAuthorsMockRecorder) FindAuthorsByBookID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthorsByBookID", reflect.TypeOf((*MockBooksAuthors)(nil).FindAuthorsByBookID), arg0)
}

// FindBooksByAuthorID mocks base method.
func (m *MockBooksAuthors) FindBooksByAuthorID(arg0 int) ([]*kvado.FindBooksResponse_Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBooksByAuthorID", arg0)
	ret0, _ := ret[0].([]*kvado.FindBooksResponse_Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBooksByAuthorID indicates an expected call of FindBooksByAuthorID.
func (mr *MockBooksAuthorsMockRecorder) FindBooksByAuthorID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBooksByAuthorID", reflect.TypeOf((*MockBooksAuthors)(nil).FindBooksByAuthorID), arg0)
}

// FindBooksByAuthorName mocks base method.
func (m *MockBooksAuthors) FindBooksByAuthorName(arg0 string) ([]*kvado.FindBooksResponse_Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBooksByAuthorName", arg0)
	ret0, _ := ret[0].([]*kvado.FindBooksResponse_Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBooksByAuthorName indicates an expected call of FindBooksByAuthorName.
func (mr *MockBooksAuthorsMockRecorder) FindBooksByAuthorName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBooksByAuthorName", reflect.TypeOf((*MockBooksAuthors)(nil).FindBooksByAuthorName), arg0)
}
