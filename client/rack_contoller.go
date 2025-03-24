package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
	"gopkg.in/mgo.v2/bson"
)

// RackController contains functionality for manipulating the RackController entity.
type RackController struct {
	APIClient APIClient
}

func (r *RackController) client(systemID string) *APIClient {
	return r.APIClient.SubClient("rackcontrollers").SubClient(systemID)
}

// Get rack controller details.
func (r *RackController) Get(ctx context.Context, systemID string) (*entity.RackController, error) {
	rackController := new(entity.RackController)
	err := r.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Update rack controller.
func (r *RackController) Update(ctx context.Context, systemID string, rackControllerParams *entity.RackControllerParams, powerParams map[string]interface{}) (*entity.RackController, error) {
	qsp, err := query.Values(rackControllerParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Delete rack controller.
func (r *RackController) Delete(ctx context.Context, systemID string) error {
	return r.client(systemID).Delete(ctx)
}

// GetPowerParameters fetches the power parameters of a given rack controller.
func (r *RackController) GetPowerParameters(ctx context.Context, systemID string) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	err := r.client(systemID).Get(ctx, "power_parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}

// PowerOn rack controller.
func (r *RackController) PowerOn(ctx context.Context, systemID string, params *entity.RackControllerPowerOnParams) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post(ctx, "power_on", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// PowerOff rack controller.
func (r *RackController) PowerOff(ctx context.Context, systemID string, params *entity.RackControllerPowerOffParams) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post(ctx, "power_off", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// GetPowerState of the rack controller.
func (r *RackController) GetPowerState(ctx context.Context, systemID string) (*entity.RackControllerPowerState, error) {
	ps := new(entity.RackControllerPowerState)
	err := r.client(systemID).Get(ctx, "query_power_state", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ps)
	})

	return ps, err
}

// Abort rack controller current operation.
func (r *RackController) Abort(ctx context.Context, systemID string, comment string) (*entity.RackController, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	rackController := new(entity.RackController)
	err := r.client(systemID).Post(ctx, "abort", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Details gets the details for a given rack controller.
func (r *RackController) Details(ctx context.Context, systemID string) (*entity.RackControllerDetails, error) {
	rackControllerDetails := new(entity.RackControllerDetails)
	err := r.client(systemID).Get(ctx, "details", url.Values{}, func(data []byte) error {
		return bson.Unmarshal(data, rackControllerDetails)
	})

	return rackControllerDetails, err
}

// OverrideFailedTesting ignores failed tests of the rack controller.
func (r *RackController) OverrideFailedTesting(ctx context.Context, systemID string, comment string) (*entity.RackController, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	rackController := new(entity.RackController)
	err := r.client(systemID).Post(ctx, "override_failed_testing", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Test the rack controller.
func (r *RackController) Test(ctx context.Context, systemID string, params map[string]interface{}) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post(ctx, "test", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}
