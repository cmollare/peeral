// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StreamCallbacks is an autogenerated mock type for the StreamCallbacks type
type StreamCallbacks struct {
	mock.Mock
}

// OnReceive provides a mock function with given fields: s, err
func (_m *StreamCallbacks) OnReceive(s string, err string) {
	_m.Called(s, err)
}