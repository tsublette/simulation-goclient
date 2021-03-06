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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewPostScanPatternSimulationParams creates a new PostScanPatternSimulationParams object
// with the default values initialized.
func NewPostScanPatternSimulationParams() *PostScanPatternSimulationParams {
	var ()
	return &PostScanPatternSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostScanPatternSimulationParamsWithTimeout creates a new PostScanPatternSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostScanPatternSimulationParamsWithTimeout(timeout time.Duration) *PostScanPatternSimulationParams {
	var ()
	return &PostScanPatternSimulationParams{

		timeout: timeout,
	}
}

// NewPostScanPatternSimulationParamsWithContext creates a new PostScanPatternSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostScanPatternSimulationParamsWithContext(ctx context.Context) *PostScanPatternSimulationParams {
	var ()
	return &PostScanPatternSimulationParams{

		Context: ctx,
	}
}

/*PostScanPatternSimulationParams contains all the parameters to send to the API endpoint
for the post scan pattern simulation operation typically these are written to a http.Request
*/
type PostScanPatternSimulationParams struct {

	/*ScanPatternSimulation
	  ScanPatternSimulation fields required to add a simulation

	*/
	ScanPatternSimulation *models.ScanPatternSimulation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) WithTimeout(timeout time.Duration) *PostScanPatternSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) WithContext(ctx context.Context) *PostScanPatternSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithScanPatternSimulation adds the scanPatternSimulation to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) WithScanPatternSimulation(scanPatternSimulation *models.ScanPatternSimulation) *PostScanPatternSimulationParams {
	o.SetScanPatternSimulation(scanPatternSimulation)
	return o
}

// SetScanPatternSimulation adds the scanPatternSimulation to the post scan pattern simulation params
func (o *PostScanPatternSimulationParams) SetScanPatternSimulation(scanPatternSimulation *models.ScanPatternSimulation) {
	o.ScanPatternSimulation = scanPatternSimulation
}

// WriteToRequest writes these params to a swagger request
func (o *PostScanPatternSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ScanPatternSimulation == nil {
		o.ScanPatternSimulation = new(models.ScanPatternSimulation)
	}

	if err := r.SetBodyParam(o.ScanPatternSimulation); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
