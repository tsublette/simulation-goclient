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

	"github.com/3dsim/simulation-goclient/models"
)

// NewPatchSimulationParams creates a new PatchSimulationParams object
// with the default values initialized.
func NewPatchSimulationParams() *PatchSimulationParams {
	var ()
	return &PatchSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSimulationParamsWithTimeout creates a new PatchSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSimulationParamsWithTimeout(timeout time.Duration) *PatchSimulationParams {
	var ()
	return &PatchSimulationParams{

		timeout: timeout,
	}
}

// NewPatchSimulationParamsWithContext creates a new PatchSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSimulationParamsWithContext(ctx context.Context) *PatchSimulationParams {
	var ()
	return &PatchSimulationParams{

		Context: ctx,
	}
}

/*PatchSimulationParams contains all the parameters to send to the API endpoint
for the patch simulation operation typically these are written to a http.Request
*/
type PatchSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int64
	/*SimulationPatch
	  This endpoint uses JSON Patch, RFC 6092.

	*/
	SimulationPatch []*models.PatchDocument

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch simulation params
func (o *PatchSimulationParams) WithTimeout(timeout time.Duration) *PatchSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch simulation params
func (o *PatchSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch simulation params
func (o *PatchSimulationParams) WithContext(ctx context.Context) *PatchSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch simulation params
func (o *PatchSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the patch simulation params
func (o *PatchSimulationParams) WithID(id int64) *PatchSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch simulation params
func (o *PatchSimulationParams) SetID(id int64) {
	o.ID = id
}

// WithSimulationPatch adds the simulationPatch to the patch simulation params
func (o *PatchSimulationParams) WithSimulationPatch(simulationPatch []*models.PatchDocument) *PatchSimulationParams {
	o.SetSimulationPatch(simulationPatch)
	return o
}

// SetSimulationPatch adds the simulationPatch to the patch simulation params
func (o *PatchSimulationParams) SetSimulationPatch(simulationPatch []*models.PatchDocument) {
	o.SimulationPatch = simulationPatch
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.SimulationPatch); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
