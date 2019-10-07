// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteClusterClusterID delete cluster cluster ID API
*/
func (a *Client) DeleteClusterClusterID(params *DeleteClusterClusterIDParams) (*DeleteClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteClusterClusterIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClusterClusterIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteClusterClusterIDTaskTaskTypeTaskID delete cluster cluster ID task task type task ID API
*/
func (a *Client) DeleteClusterClusterIDTaskTaskTypeTaskID(params *DeleteClusterClusterIDTaskTaskTypeTaskIDParams) (*DeleteClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterClusterIDTaskTaskTypeTaskID",
		Method:             "DELETE",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteClusterClusterIDTaskTaskTypeTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteClusterClusterIDTaskTaskTypeTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterID get cluster cluster ID API
*/
func (a *Client) GetClusterClusterID(params *GetClusterClusterIDParams) (*GetClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDStatus get cluster cluster ID status API
*/
func (a *Client) GetClusterClusterIDStatus(params *GetClusterClusterIDStatusParams) (*GetClusterClusterIDStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDStatus",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDTaskBackupTaskIDRunID get cluster cluster ID task backup task ID run ID API
*/
func (a *Client) GetClusterClusterIDTaskBackupTaskIDRunID(params *GetClusterClusterIDTaskBackupTaskIDRunIDParams) (*GetClusterClusterIDTaskBackupTaskIDRunIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTaskBackupTaskIDRunIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTaskBackupTaskIDRunID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/task/backup/{task_id}/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTaskBackupTaskIDRunIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDTaskBackupTaskIDRunIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDTaskBackupTaskIDRunIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDTaskRepairTaskIDRunID get cluster cluster ID task repair task ID run ID API
*/
func (a *Client) GetClusterClusterIDTaskRepairTaskIDRunID(params *GetClusterClusterIDTaskRepairTaskIDRunIDParams) (*GetClusterClusterIDTaskRepairTaskIDRunIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTaskRepairTaskIDRunIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTaskRepairTaskIDRunID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/task/repair/{task_id}/{run_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTaskRepairTaskIDRunIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDTaskRepairTaskIDRunIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDTaskRepairTaskIDRunIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDTaskTaskTypeTaskID get cluster cluster ID task task type task ID API
*/
func (a *Client) GetClusterClusterIDTaskTaskTypeTaskID(params *GetClusterClusterIDTaskTaskTypeTaskIDParams) (*GetClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTaskTaskTypeTaskID",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDTaskTaskTypeTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDTaskTaskTypeTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDTaskTaskTypeTaskIDHistory get cluster cluster ID task task type task ID history API
*/
func (a *Client) GetClusterClusterIDTaskTaskTypeTaskIDHistory(params *GetClusterClusterIDTaskTaskTypeTaskIDHistoryParams) (*GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTaskTaskTypeTaskIDHistory",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusterClusterIDTasks get cluster cluster ID tasks API
*/
func (a *Client) GetClusterClusterIDTasks(params *GetClusterClusterIDTasksParams) (*GetClusterClusterIDTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterClusterIDTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterClusterIDTasks",
		Method:             "GET",
		PathPattern:        "/cluster/{cluster_id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterClusterIDTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClusterClusterIDTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClusterClusterIDTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClusters get clusters API
*/
func (a *Client) GetClusters(params *GetClustersParams) (*GetClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetVersion get version API
*/
func (a *Client) GetVersion(params *GetVersionParams) (*GetVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetVersion",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostClusterClusterIDTasks post cluster cluster ID tasks API
*/
func (a *Client) PostClusterClusterIDTasks(params *PostClusterClusterIDTasksParams) (*PostClusterClusterIDTasksCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterClusterIDTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClusterClusterIDTasks",
		Method:             "POST",
		PathPattern:        "/cluster/{cluster_id}/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClusterClusterIDTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClusterClusterIDTasksCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostClusterClusterIDTasksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostClusters post clusters API
*/
func (a *Client) PostClusters(params *PostClustersParams) (*PostClustersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClusters",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClustersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutClusterClusterID put cluster cluster ID API
*/
func (a *Client) PutClusterClusterID(params *PutClusterClusterIDParams) (*PutClusterClusterIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterClusterIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutClusterClusterIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutClusterClusterIDTaskTaskTypeTaskID put cluster cluster ID task task type task ID API
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskID(params *PutClusterClusterIDTaskTaskTypeTaskIDParams) (*PutClusterClusterIDTaskTaskTypeTaskIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskID",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterClusterIDTaskTaskTypeTaskIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutClusterClusterIDTaskTaskTypeTaskIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutClusterClusterIDTaskTaskTypeTaskIDStart put cluster cluster ID task task type task ID start API
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskIDStart(params *PutClusterClusterIDTaskTaskTypeTaskIDStartParams) (*PutClusterClusterIDTaskTaskTypeTaskIDStartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskIDStart",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDStartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterClusterIDTaskTaskTypeTaskIDStartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutClusterClusterIDTaskTaskTypeTaskIDStartDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutClusterClusterIDTaskTaskTypeTaskIDStop put cluster cluster ID task task type task ID stop API
*/
func (a *Client) PutClusterClusterIDTaskTaskTypeTaskIDStop(params *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) (*PutClusterClusterIDTaskTaskTypeTaskIDStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTaskTaskTypeTaskIDStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTaskTaskTypeTaskIDStop",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/task/{task_type}/{task_id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTaskTaskTypeTaskIDStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterClusterIDTaskTaskTypeTaskIDStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutClusterClusterIDTaskTaskTypeTaskIDStopDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutClusterClusterIDTasksRepairTarget put cluster cluster ID tasks repair target API
*/
func (a *Client) PutClusterClusterIDTasksRepairTarget(params *PutClusterClusterIDTasksRepairTargetParams) (*PutClusterClusterIDTasksRepairTargetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterClusterIDTasksRepairTargetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClusterClusterIDTasksRepairTarget",
		Method:             "PUT",
		PathPattern:        "/cluster/{cluster_id}/tasks/repair/target",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClusterClusterIDTasksRepairTargetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutClusterClusterIDTasksRepairTargetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutClusterClusterIDTasksRepairTargetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
