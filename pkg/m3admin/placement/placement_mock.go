// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db-operator/pkg/m3admin/placement/types.go

package placement

import (
	gomock "github.com/golang/mock/gomock"
	admin "github.com/m3db/m3/src/query/generated/proto/admin"
	placementpb "github.com/m3db/m3cluster/generated/proto/placementpb"
	placement "github.com/m3db/m3cluster/placement"
)

// Mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

// Recorder for MockClient (not exported)
type _MockClientRecorder struct {
	mock *MockClient
}

func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

func (_m *MockClient) EXPECT() *_MockClientRecorder {
	return _m.recorder
}

func (_m *MockClient) Init(request *admin.PlacementInitRequest) error {
	ret := _m.ctrl.Call(_m, "Init", request)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockClient) Get() (placement.Placement, error) {
	ret := _m.ctrl.Call(_m, "Get")
	ret0, _ := ret[0].(placement.Placement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClientRecorder) Get() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get")
}

func (_m *MockClient) Delete() error {
	ret := _m.ctrl.Call(_m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Delete() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete")
}

func (_m *MockClient) Add(instance placementpb.Instance) error {
	ret := _m.ctrl.Call(_m, "Add", instance)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Add", arg0)
}