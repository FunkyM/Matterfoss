// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/cjdelisle/matterfoss-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// IndexerJobInterface is an autogenerated mock type for the IndexerJobInterface type
type IndexerJobInterface struct {
	mock.Mock
}

// MakeWorker provides a mock function with given fields:
func (_m *IndexerJobInterface) MakeWorker() model.Worker {
	ret := _m.Called()

	var r0 model.Worker
	if rf, ok := ret.Get(0).(func() model.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Worker)
		}
	}

	return r0
}
