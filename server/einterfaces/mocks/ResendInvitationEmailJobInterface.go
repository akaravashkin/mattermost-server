// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make einterfaces-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// ResendInvitationEmailJobInterface is an autogenerated mock type for the ResendInvitationEmailJobInterface type
type ResendInvitationEmailJobInterface struct {
	mock.Mock
}

// MakeScheduler provides a mock function with given fields:
func (_m *ResendInvitationEmailJobInterface) MakeScheduler() model.Scheduler {
	ret := _m.Called()

	var r0 model.Scheduler
	if rf, ok := ret.Get(0).(func() model.Scheduler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Scheduler)
		}
	}

	return r0
}

// MakeWorker provides a mock function with given fields:
func (_m *ResendInvitationEmailJobInterface) MakeWorker() model.Worker {
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
