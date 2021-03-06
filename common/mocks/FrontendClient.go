// Copyright (c) 2017 Uber Technologies, Inc.
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

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	"github.com/uber/cadence/.gen/go/cadence/workflowserviceclient"
	shared "github.com/uber/cadence/.gen/go/shared"
	yarpc "go.uber.org/yarpc"
)

// FrontendClient is an autogenerated mock type for the Client type
type FrontendClient struct {
	mock.Mock
}

var _ workflowserviceclient.Interface = (*FrontendClient)(nil)

// DeprecateDomain provides a mock function with given fields: ctx, DeprecateRequest, opts
func (_m *FrontendClient) DeprecateDomain(ctx context.Context, DeprecateRequest *shared.DeprecateDomainRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, DeprecateRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.DeprecateDomainRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, DeprecateRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DescribeDomain provides a mock function with given fields: ctx, DescribeRequest, opts
func (_m *FrontendClient) DescribeDomain(ctx context.Context, DescribeRequest *shared.DescribeDomainRequest, opts ...yarpc.CallOption) (*shared.DescribeDomainResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, DescribeRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.DescribeDomainResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.DescribeDomainRequest, ...yarpc.CallOption) *shared.DescribeDomainResponse); ok {
		r0 = rf(ctx, DescribeRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DescribeDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.DescribeDomainRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, DescribeRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeTaskList provides a mock function with given fields: ctx, Request, opts
func (_m *FrontendClient) DescribeTaskList(ctx context.Context, Request *shared.DescribeTaskListRequest, opts ...yarpc.CallOption) (*shared.DescribeTaskListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, Request)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.DescribeTaskListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.DescribeTaskListRequest, ...yarpc.CallOption) *shared.DescribeTaskListResponse); ok {
		r0 = rf(ctx, Request, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DescribeTaskListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.DescribeTaskListRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, Request, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeWorkflowExecution provides a mock function with given fields: ctx, DescribeRequest, opts
func (_m *FrontendClient) DescribeWorkflowExecution(ctx context.Context, DescribeRequest *shared.DescribeWorkflowExecutionRequest, opts ...yarpc.CallOption) (*shared.DescribeWorkflowExecutionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, DescribeRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.DescribeWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.DescribeWorkflowExecutionRequest, ...yarpc.CallOption) *shared.DescribeWorkflowExecutionResponse); ok {
		r0 = rf(ctx, DescribeRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.DescribeWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.DescribeWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, DescribeRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowExecutionHistory provides a mock function with given fields: ctx, GetRequest, opts
func (_m *FrontendClient) GetWorkflowExecutionHistory(ctx context.Context, GetRequest *shared.GetWorkflowExecutionHistoryRequest, opts ...yarpc.CallOption) (*shared.GetWorkflowExecutionHistoryResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, GetRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.GetWorkflowExecutionHistoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.GetWorkflowExecutionHistoryRequest, ...yarpc.CallOption) *shared.GetWorkflowExecutionHistoryResponse); ok {
		r0 = rf(ctx, GetRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetWorkflowExecutionHistoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.GetWorkflowExecutionHistoryRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, GetRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClosedWorkflowExecutions provides a mock function with given fields: ctx, ListRequest, opts
func (_m *FrontendClient) ListClosedWorkflowExecutions(ctx context.Context, ListRequest *shared.ListClosedWorkflowExecutionsRequest, opts ...yarpc.CallOption) (*shared.ListClosedWorkflowExecutionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ListRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ListClosedWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ListClosedWorkflowExecutionsRequest, ...yarpc.CallOption) *shared.ListClosedWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, ListRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListClosedWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ListClosedWorkflowExecutionsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ListRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDomains provides a mock function with given fields: ctx, ListRequest, opts
func (_m *FrontendClient) ListDomains(ctx context.Context, ListRequest *shared.ListDomainsRequest, opts ...yarpc.CallOption) (*shared.ListDomainsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ListRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ListDomainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ListDomainsRequest, ...yarpc.CallOption) *shared.ListDomainsResponse); ok {
		r0 = rf(ctx, ListRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListDomainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ListDomainsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ListRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOpenWorkflowExecutions provides a mock function with given fields: ctx, ListRequest, opts
func (_m *FrontendClient) ListOpenWorkflowExecutions(ctx context.Context, ListRequest *shared.ListOpenWorkflowExecutionsRequest, opts ...yarpc.CallOption) (*shared.ListOpenWorkflowExecutionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ListRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ListOpenWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ListOpenWorkflowExecutionsRequest, ...yarpc.CallOption) *shared.ListOpenWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, ListRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListOpenWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ListOpenWorkflowExecutionsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ListRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWorkflowExecutions provides a mock function with given fields: ctx, ListRequest, opts
func (_m *FrontendClient) ListWorkflowExecutions(ctx context.Context, ListRequest *shared.ListWorkflowExecutionsRequest, opts ...yarpc.CallOption) (*shared.ListWorkflowExecutionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ListRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ListWorkflowExecutionsRequest, ...yarpc.CallOption) *shared.ListWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, ListRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ListWorkflowExecutionsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ListRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanWorkflowExecutions provides a mock function with given fields: ctx, ListRequest, opts
func (_m *FrontendClient) ScanWorkflowExecutions(ctx context.Context, ListRequest *shared.ListWorkflowExecutionsRequest, opts ...yarpc.CallOption) (*shared.ListWorkflowExecutionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ListRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ListWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ListWorkflowExecutionsRequest, ...yarpc.CallOption) *shared.ListWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, ListRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ListWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ListWorkflowExecutionsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ListRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountWorkflowExecutions provides a mock function with given fields: ctx, CountRequest, opts
func (_m *FrontendClient) CountWorkflowExecutions(ctx context.Context, CountRequest *shared.CountWorkflowExecutionsRequest, opts ...yarpc.CallOption) (*shared.CountWorkflowExecutionsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CountRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.CountWorkflowExecutionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.CountWorkflowExecutionsRequest, ...yarpc.CallOption) *shared.CountWorkflowExecutionsResponse); ok {
		r0 = rf(ctx, CountRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.CountWorkflowExecutionsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.CountWorkflowExecutionsRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, CountRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSearchAttributes provides a mock function with given fields: ctx, opts
func (_m *FrontendClient) GetSearchAttributes(ctx context.Context, opts ...yarpc.CallOption) (*shared.GetSearchAttributesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.GetSearchAttributesResponse
	if rf, ok := ret.Get(0).(func(context.Context, ...yarpc.CallOption) *shared.GetSearchAttributesResponse); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.GetSearchAttributesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForActivityTask provides a mock function with given fields: ctx, PollRequest, opts
func (_m *FrontendClient) PollForActivityTask(ctx context.Context, PollRequest *shared.PollForActivityTaskRequest, opts ...yarpc.CallOption) (*shared.PollForActivityTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, PollRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.PollForActivityTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.PollForActivityTaskRequest, ...yarpc.CallOption) *shared.PollForActivityTaskResponse); ok {
		r0 = rf(ctx, PollRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForActivityTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.PollForActivityTaskRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, PollRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PollForDecisionTask provides a mock function with given fields: ctx, PollRequest, opts
func (_m *FrontendClient) PollForDecisionTask(ctx context.Context, PollRequest *shared.PollForDecisionTaskRequest, opts ...yarpc.CallOption) (*shared.PollForDecisionTaskResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, PollRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.PollForDecisionTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.PollForDecisionTaskRequest, ...yarpc.CallOption) *shared.PollForDecisionTaskResponse); ok {
		r0 = rf(ctx, PollRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.PollForDecisionTaskResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.PollForDecisionTaskRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, PollRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWorkflow provides a mock function with given fields: ctx, QueryRequest, opts
func (_m *FrontendClient) QueryWorkflow(ctx context.Context, QueryRequest *shared.QueryWorkflowRequest, opts ...yarpc.CallOption) (*shared.QueryWorkflowResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, QueryRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.QueryWorkflowResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.QueryWorkflowRequest, ...yarpc.CallOption) *shared.QueryWorkflowResponse); ok {
		r0 = rf(ctx, QueryRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.QueryWorkflowResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.QueryWorkflowRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, QueryRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordActivityTaskHeartbeat provides a mock function with given fields: ctx, HeartbeatRequest, opts
func (_m *FrontendClient) RecordActivityTaskHeartbeat(ctx context.Context, HeartbeatRequest *shared.RecordActivityTaskHeartbeatRequest, opts ...yarpc.CallOption) (*shared.RecordActivityTaskHeartbeatResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, HeartbeatRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.RecordActivityTaskHeartbeatResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RecordActivityTaskHeartbeatRequest, ...yarpc.CallOption) *shared.RecordActivityTaskHeartbeatResponse); ok {
		r0 = rf(ctx, HeartbeatRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RecordActivityTaskHeartbeatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.RecordActivityTaskHeartbeatRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, HeartbeatRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecordActivityTaskHeartbeatByID provides a mock function with given fields: ctx, HeartbeatRequest, opts
func (_m *FrontendClient) RecordActivityTaskHeartbeatByID(ctx context.Context, HeartbeatRequest *shared.RecordActivityTaskHeartbeatByIDRequest, opts ...yarpc.CallOption) (*shared.RecordActivityTaskHeartbeatResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, HeartbeatRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.RecordActivityTaskHeartbeatResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RecordActivityTaskHeartbeatByIDRequest, ...yarpc.CallOption) *shared.RecordActivityTaskHeartbeatResponse); ok {
		r0 = rf(ctx, HeartbeatRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RecordActivityTaskHeartbeatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.RecordActivityTaskHeartbeatByIDRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, HeartbeatRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterDomain provides a mock function with given fields: ctx, RegisterRequest, opts
func (_m *FrontendClient) RegisterDomain(ctx context.Context, RegisterRequest *shared.RegisterDomainRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, RegisterRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RegisterDomainRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, RegisterRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestCancelWorkflowExecution provides a mock function with given fields: ctx, CancelRequest, opts
func (_m *FrontendClient) RequestCancelWorkflowExecution(ctx context.Context, CancelRequest *shared.RequestCancelWorkflowExecutionRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CancelRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RequestCancelWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CancelRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetStickyTaskList provides a mock function with given fields: ctx, ResetRequest, opts
func (_m *FrontendClient) ResetStickyTaskList(ctx context.Context, ResetRequest *shared.ResetStickyTaskListRequest, opts ...yarpc.CallOption) (*shared.ResetStickyTaskListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ResetRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ResetStickyTaskListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ResetStickyTaskListRequest, ...yarpc.CallOption) *shared.ResetStickyTaskListResponse); ok {
		r0 = rf(ctx, ResetRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ResetStickyTaskListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ResetStickyTaskListRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ResetRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetWorkflowExecution provides a mock function with given fields: ctx, ResetRequest, opts
func (_m *FrontendClient) ResetWorkflowExecution(ctx context.Context, ResetRequest *shared.ResetWorkflowExecutionRequest, opts ...yarpc.CallOption) (*shared.ResetWorkflowExecutionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, ResetRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.ResetWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.ResetWorkflowExecutionRequest, ...yarpc.CallOption) *shared.ResetWorkflowExecutionResponse); ok {
		r0 = rf(ctx, ResetRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.ResetWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.ResetWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, ResetRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RespondActivityTaskCanceled provides a mock function with given fields: ctx, CanceledRequest, opts
func (_m *FrontendClient) RespondActivityTaskCanceled(ctx context.Context, CanceledRequest *shared.RespondActivityTaskCanceledRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CanceledRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskCanceledRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CanceledRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskCanceledByID provides a mock function with given fields: ctx, CanceledRequest, opts
func (_m *FrontendClient) RespondActivityTaskCanceledByID(ctx context.Context, CanceledRequest *shared.RespondActivityTaskCanceledByIDRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CanceledRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskCanceledByIDRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CanceledRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskCompleted provides a mock function with given fields: ctx, CompleteRequest, opts
func (_m *FrontendClient) RespondActivityTaskCompleted(ctx context.Context, CompleteRequest *shared.RespondActivityTaskCompletedRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CompleteRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskCompletedRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CompleteRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskCompletedByID provides a mock function with given fields: ctx, CompleteRequest, opts
func (_m *FrontendClient) RespondActivityTaskCompletedByID(ctx context.Context, CompleteRequest *shared.RespondActivityTaskCompletedByIDRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CompleteRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskCompletedByIDRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CompleteRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskFailed provides a mock function with given fields: ctx, FailRequest, opts
func (_m *FrontendClient) RespondActivityTaskFailed(ctx context.Context, FailRequest *shared.RespondActivityTaskFailedRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, FailRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskFailedRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, FailRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondActivityTaskFailedByID provides a mock function with given fields: ctx, FailRequest, opts
func (_m *FrontendClient) RespondActivityTaskFailedByID(ctx context.Context, FailRequest *shared.RespondActivityTaskFailedByIDRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, FailRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondActivityTaskFailedByIDRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, FailRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondDecisionTaskCompleted provides a mock function with given fields: ctx, CompleteRequest, opts
func (_m *FrontendClient) RespondDecisionTaskCompleted(ctx context.Context, CompleteRequest *shared.RespondDecisionTaskCompletedRequest, opts ...yarpc.CallOption) (*shared.RespondDecisionTaskCompletedResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CompleteRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.RespondDecisionTaskCompletedResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondDecisionTaskCompletedRequest, ...yarpc.CallOption) *shared.RespondDecisionTaskCompletedResponse); ok {
		r0 = rf(ctx, CompleteRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.RespondDecisionTaskCompletedResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.RespondDecisionTaskCompletedRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, CompleteRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RespondDecisionTaskFailed provides a mock function with given fields: ctx, FailedRequest, opts
func (_m *FrontendClient) RespondDecisionTaskFailed(ctx context.Context, FailedRequest *shared.RespondDecisionTaskFailedRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, FailedRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondDecisionTaskFailedRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, FailedRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RespondQueryTaskCompleted provides a mock function with given fields: ctx, CompleteRequest, opts
func (_m *FrontendClient) RespondQueryTaskCompleted(ctx context.Context, CompleteRequest *shared.RespondQueryTaskCompletedRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, CompleteRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.RespondQueryTaskCompletedRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, CompleteRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignalWithStartWorkflowExecution provides a mock function with given fields: ctx, SignalWithStartRequest, opts
func (_m *FrontendClient) SignalWithStartWorkflowExecution(ctx context.Context, SignalWithStartRequest *shared.SignalWithStartWorkflowExecutionRequest, opts ...yarpc.CallOption) (*shared.StartWorkflowExecutionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, SignalWithStartRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.StartWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.SignalWithStartWorkflowExecutionRequest, ...yarpc.CallOption) *shared.StartWorkflowExecutionResponse); ok {
		r0 = rf(ctx, SignalWithStartRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.StartWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.SignalWithStartWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, SignalWithStartRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignalWorkflowExecution provides a mock function with given fields: ctx, SignalRequest, opts
func (_m *FrontendClient) SignalWorkflowExecution(ctx context.Context, SignalRequest *shared.SignalWorkflowExecutionRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, SignalRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.SignalWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, SignalRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartWorkflowExecution provides a mock function with given fields: ctx, StartRequest, opts
func (_m *FrontendClient) StartWorkflowExecution(ctx context.Context, StartRequest *shared.StartWorkflowExecutionRequest, opts ...yarpc.CallOption) (*shared.StartWorkflowExecutionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, StartRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.StartWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.StartWorkflowExecutionRequest, ...yarpc.CallOption) *shared.StartWorkflowExecutionResponse); ok {
		r0 = rf(ctx, StartRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.StartWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.StartWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, StartRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TerminateWorkflowExecution provides a mock function with given fields: ctx, TerminateRequest, opts
func (_m *FrontendClient) TerminateWorkflowExecution(ctx context.Context, TerminateRequest *shared.TerminateWorkflowExecutionRequest, opts ...yarpc.CallOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, TerminateRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *shared.TerminateWorkflowExecutionRequest, ...yarpc.CallOption) error); ok {
		r0 = rf(ctx, TerminateRequest, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDomain provides a mock function with given fields: ctx, UpdateRequest, opts
func (_m *FrontendClient) UpdateDomain(ctx context.Context, UpdateRequest *shared.UpdateDomainRequest, opts ...yarpc.CallOption) (*shared.UpdateDomainResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, UpdateRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *shared.UpdateDomainResponse
	if rf, ok := ret.Get(0).(func(context.Context, *shared.UpdateDomainRequest, ...yarpc.CallOption) *shared.UpdateDomainResponse); ok {
		r0 = rf(ctx, UpdateRequest, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*shared.UpdateDomainResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *shared.UpdateDomainRequest, ...yarpc.CallOption) error); ok {
		r1 = rf(ctx, UpdateRequest, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
