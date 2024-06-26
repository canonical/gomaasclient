package client

import (
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

func (r *RackController) client(systemID string) APIClient {
	return r.APIClient.GetSubObject("rackcontrollers").GetSubObject(systemID)
}

// Get rack controller details.
func (r *RackController) Get(systemID string) (*entity.RackController, error) {
	rackController := new(entity.RackController)
	err := r.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Update rack controller.
func (r *RackController) Update(systemID string, rackControllerParams *entity.RackControllerParams, powerParams map[string]interface{}) (*entity.RackController, error) {
	qsp, err := query.Values(rackControllerParams)
	if err != nil {
		return nil, err
	}

	for k, v := range powerParamsToURLValues(powerParams) {
		// Since qsp.Add(k, v...) is not allowed
		qsp[k] = append(qsp[k], v...)
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Delete rack controller.
func (r *RackController) Delete(systemID string) error {
	return r.client(systemID).Delete()
}

// GetPowerParameters fetches the power parameters of a given rack controller.
func (r *RackController) GetPowerParameters(systemID string) (map[string]interface{}, error) {
	params := map[string]interface{}{}
	err := r.client(systemID).Get("power_parameters", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &params)
	})

	return params, err
}

// PowerOn rack controller.
func (r *RackController) PowerOn(systemID string, params *entity.RackControllerPowerOnParams) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post("power_on", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// PowerOff rack controller.
func (r *RackController) PowerOff(systemID string, params *entity.RackControllerPowerOffParams) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post("power_off", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// GetPowerState of the rack controller.
func (r *RackController) GetPowerState(systemID string) (*entity.RackControllerPowerState, error) {
	ps := new(entity.RackControllerPowerState)
	err := r.client(systemID).Get("query_power_state", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ps)
	})

	return ps, err
}

// Abort rack controller current operation.
func (r *RackController) Abort(systemID string, comment string) (*entity.RackController, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	rackController := new(entity.RackController)
	err := r.client(systemID).Post("abort", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Details gets the details for a given rack controller.
func (r *RackController) Details(systemID string) (*entity.RackControllerDetails, error) {
	rackControllerDetails := new(entity.RackControllerDetails)
	err := r.client(systemID).Get("details", url.Values{}, func(data []byte) error {
		return bson.Unmarshal(data, rackControllerDetails)
	})

	return rackControllerDetails, err
}

// OverrideFailedTesting ignores failed tests of the rack controller.
func (r *RackController) OverrideFailedTesting(systemID string, comment string) (*entity.RackController, error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}

	rackController := new(entity.RackController)
	err := r.client(systemID).Post("override_failed_testing", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}

// Test the rack controller.
func (r *RackController) Test(systemID string, params map[string]interface{}) (*entity.RackController, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	rackController := new(entity.RackController)
	err = r.client(systemID).Post("test", qsp, func(data []byte) error {
		return json.Unmarshal(data, rackController)
	})

	return rackController, err
}
