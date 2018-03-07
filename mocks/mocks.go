// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/leotse/srtfix (interfaces: FileReader,FileWriter,Parser,Resolver)

// Package mock_srtfix is a generated GoMock package.
package mock_srtfix

import (
	gomock "github.com/golang/mock/gomock"
	srtfix "github.com/leotse/srtfix"
	reflect "reflect"
)

// MockFileReader is a mock of FileReader interface
type MockFileReader struct {
	ctrl     *gomock.Controller
	recorder *MockFileReaderMockRecorder
}

// MockFileReaderMockRecorder is the mock recorder for MockFileReader
type MockFileReaderMockRecorder struct {
	mock *MockFileReader
}

// NewMockFileReader creates a new mock instance
func NewMockFileReader(ctrl *gomock.Controller) *MockFileReader {
	mock := &MockFileReader{ctrl: ctrl}
	mock.recorder = &MockFileReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileReader) EXPECT() *MockFileReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockFileReader) Read() (string, error) {
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockFileReaderMockRecorder) Read() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileReader)(nil).Read))
}

// MockFileWriter is a mock of FileWriter interface
type MockFileWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileWriterMockRecorder
}

// MockFileWriterMockRecorder is the mock recorder for MockFileWriter
type MockFileWriterMockRecorder struct {
	mock *MockFileWriter
}

// NewMockFileWriter creates a new mock instance
func NewMockFileWriter(ctrl *gomock.Controller) *MockFileWriter {
	mock := &MockFileWriter{ctrl: ctrl}
	mock.recorder = &MockFileWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileWriter) EXPECT() *MockFileWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockFileWriter) Write(arg0 string) error {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockFileWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileWriter)(nil).Write), arg0)
}

// MockParser is a mock of Parser interface
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserMockRecorder
}

// MockParserMockRecorder is the mock recorder for MockParser
type MockParserMockRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockParser) EXPECT() *MockParserMockRecorder {
	return m.recorder
}

// Parse mocks base method
func (m *MockParser) Parse(arg0 string) ([]*srtfix.Caption, error) {
	ret := m.ctrl.Call(m, "Parse", arg0)
	ret0, _ := ret[0].([]*srtfix.Caption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockParserMockRecorder) Parse(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), arg0)
}

// MockResolver is a mock of Resolver interface
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method
func (m *MockResolver) Resolve(arg0 []*srtfix.Caption) []*srtfix.Caption {
	ret := m.ctrl.Call(m, "Resolve", arg0)
	ret0, _ := ret[0].([]*srtfix.Caption)
	return ret0
}

// Resolve indicates an expected call of Resolve
func (mr *MockResolverMockRecorder) Resolve(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockResolver)(nil).Resolve), arg0)
}