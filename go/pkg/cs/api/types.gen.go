// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// CA defines model for CA.
type CA struct {
	CertValidity *struct {
		NotAfter  *time.Time `json:"not_after,omitempty"`
		NotBefore *time.Time `json:"not_before,omitempty"`
	} `json:"cert_validity,omitempty"`
	Policy *struct {
		ChainLifetime *string `json:"chain_lifetime,omitempty"`
	} `json:"policy,omitempty"`
	Subject *struct {
		IsdAs *string `json:"isd_as,omitempty"`
	} `json:"subject,omitempty"`
	SubjectKeyId *string `json:"subject_key_id,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// Error message
	Error string `json:"error"`
}

// LogLevel defines model for LogLevel.
type LogLevel struct {

	// Logging level
	Level string `json:"level"`
}

// Signer defines model for Signer.
type Signer struct {
	CertValidity *struct {
		NotAfter  *time.Time `json:"not_after,omitempty"`
		NotBefore *time.Time `json:"not_before,omitempty"`
	} `json:"cert_validity,omitempty"`
	Expiration    *time.Time `json:"expiration,omitempty"`
	InGracePeriod *bool      `json:"in_grace_period,omitempty"`
	Subject       *struct {
		IsdAs *string `json:"isd_as,omitempty"`
	} `json:"subject,omitempty"`
	SubjectKeyId *string `json:"subject_key_id,omitempty"`
	TrcId        *struct {
		BaseNumber   *float32 `json:"base_number,omitempty"`
		Isd          *float32 `json:"isd,omitempty"`
		SerialNumber *float32 `json:"serial_number,omitempty"`
	} `json:"trc_id,omitempty"`
}

// Topology defines model for Topology.
type Topology struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// BadRequest defines model for BadRequest.
type BadRequest Error

// SetLogLevelJSONBody defines parameters for SetLogLevel.
type SetLogLevelJSONBody LogLevel

// SetLogLevelRequestBody defines body for SetLogLevel for application/json ContentType.
type SetLogLevelJSONRequestBody SetLogLevelJSONBody

// Getter for additional properties for Topology. Returns the specified
// element and whether it was found
func (a Topology) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Topology
func (a *Topology) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Topology to handle AdditionalProperties
func (a *Topology) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Topology to handle AdditionalProperties
func (a Topology) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
