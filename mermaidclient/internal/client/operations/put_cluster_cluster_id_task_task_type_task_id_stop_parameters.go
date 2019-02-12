// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopParams creates a new PutClusterClusterIDTaskTaskTypeTaskIDStopParams object
// with the default values initialized.
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopParams() *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithTimeout creates a new PutClusterClusterIDTaskTaskTypeTaskIDStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithTimeout(timeout time.Duration) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopParams{

		timeout: timeout,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithContext creates a new PutClusterClusterIDTaskTaskTypeTaskIDStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithContext(ctx context.Context) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopParams{

		Context: ctx,
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithHTTPClient creates a new PutClusterClusterIDTaskTaskTypeTaskIDStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopParamsWithHTTPClient(client *http.Client) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	var ()
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopParams{
		HTTPClient: client,
	}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopParams contains all the parameters to send to the API endpoint
for the put cluster cluster ID task task type task ID stop operation typically these are written to a http.Request
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopParams struct {

	/*ClusterID
	  cluster ID this API is performing on

	*/
	ClusterID string
	/*Disable
	  do not run in future

	*/
	Disable *bool
	/*TaskID
	  task ID (UUID) or name

	*/
	TaskID string
	/*TaskType
	  task type

	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithTimeout(timeout time.Duration) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithContext(ctx context.Context) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithHTTPClient(client *http.Client) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithClusterID(clusterID string) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDisable adds the disable to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithDisable(disable *bool) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetDisable(disable)
	return o
}

// SetDisable adds the disable to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetDisable(disable *bool) {
	o.Disable = disable
}

// WithTaskID adds the taskID to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithTaskID(taskID string) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WithTaskType adds the taskType to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WithTaskType(taskType string) *PutClusterClusterIDTaskTaskTypeTaskIDStopParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the put cluster cluster ID task task type task ID stop params
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Disable != nil {

		// query param disable
		var qrDisable bool
		if o.Disable != nil {
			qrDisable = *o.Disable
		}
		qDisable := swag.FormatBool(qrDisable)
		if qDisable != "" {
			if err := r.SetQueryParam("disable", qDisable); err != nil {
				return err
			}
		}

	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	// path param task_type
	if err := r.SetPathParam("task_type", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
