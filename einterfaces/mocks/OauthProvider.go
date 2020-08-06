// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	io "io"

	model "github.com/cjdelisle/matterfoss-server/v5/model"
	mock "github.com/stretchr/testify/mock"
)

// OauthProvider is an autogenerated mock type for the OauthProvider type
type OauthProvider struct {
	mock.Mock
}

// GetUserFromJson provides a mock function with given fields: data
func (_m *OauthProvider) GetUserFromJson(data io.Reader) *model.User {
	ret := _m.Called(data)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(io.Reader) *model.User); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	return r0
}
