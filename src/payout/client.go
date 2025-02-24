/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// APIClient manages communication with the Adyen Payout API API v68
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	InitializationApi *InitializationApi

	InstantPayoutsApi *InstantPayoutsApi

	ReviewingApi *ReviewingApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
	c.common.Client = client
	c.common.BasePath = func() string {
		return client.Cfg.Endpoint
	}

	// API Services
	c.InitializationApi = (*InitializationApi)(&c.common)
	c.InstantPayoutsApi = (*InstantPayoutsApi)(&c.common)
	c.ReviewingApi = (*ReviewingApi)(&c.common)

	return c
}
