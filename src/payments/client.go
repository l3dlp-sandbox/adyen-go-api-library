/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// APIClient manages communication with the Adyen Payment API API v68
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ModificationsApi *ModificationsApi

	PaymentsApi *PaymentsApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
	c.common.Client = client
	c.common.BasePath = func() string {
		return client.Cfg.Endpoint
	}

	// API Services
	c.ModificationsApi = (*ModificationsApi)(&c.common)
	c.PaymentsApi = (*PaymentsApi)(&c.common)

	return c
}
