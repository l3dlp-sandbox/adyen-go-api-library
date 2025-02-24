/*
 * Adyen for Platforms: Notification Configuration API
 *
 * The Notification Configuration API provides endpoints for setting up and testing notifications that inform you of events on your platform, for example when a KYC check or a payout has been completed.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications). ## Authentication To connect to the Notification Configuration API, you must use basic authentication credentials of your web service user. If you don't have one, contact our [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Notification Configuration API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Notification/v6/createNotificationConfiguration ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsnotificationconfiguration

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// PlatformsNotificationConfiguration PlatformsNotificationConfiguration service
// Deprecated: Please migrate to the new Adyen For Platforms.
type PlatformsNotificationConfiguration common.Service

/*
PostCreateNotificationConfiguration Subscribe to notifications.
Creates a subscription to notifications informing you of events on your platform. After the subscription is created, the events specified in the configuration will be sent to the URL specified in the configuration. Subscriptions must be configured on a per-event basis (as opposed to, for example, a per-account holder basis), so all event notifications of a marketplace and of a given type will be sent to the same endpoint(s). A marketplace may have multiple endpoints if desired; an event notification may be sent to as many or as few different endpoints as configured.
  - @param request CreateNotificationConfigurationRequest - reference of CreateNotificationConfigurationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetNotificationConfigurationResponse
*/
func (a PlatformsNotificationConfiguration) CreateNotificationConfiguration(req *CreateNotificationConfigurationRequest, ctxs ..._context.Context) (GetNotificationConfigurationResponse, *_nethttp.Response, error) {
	res := &GetNotificationConfigurationResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/createNotificationConfiguration", ctxs...)
	return *res, httpRes, err
}

/*
PostDeleteNotificationConfigurations Delete an existing notification subscription configuration.
This endpoint is used to delete an existing notification subscription configuration. After the subscription is deleted, no further event notifications will be sent to the URL that was in the subscription.
  - @param request DeleteNotificationConfigurationRequest - reference of DeleteNotificationConfigurationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GenericResponse
*/
func (a PlatformsNotificationConfiguration) DeleteNotificationConfigurations(req *DeleteNotificationConfigurationRequest, ctxs ..._context.Context) (GenericResponse, *_nethttp.Response, error) {
	res := &GenericResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/deleteNotificationConfigurations", ctxs...)
	return *res, httpRes, err
}

/*
PostGetNotificationConfiguration Retrieve an existing notification subscription configuration.
This endpoint is used to retrieve the details of the configuration of a notification subscription.
  - @param request GetNotificationConfigurationRequest - reference of GetNotificationConfigurationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetNotificationConfigurationResponse
*/
func (a PlatformsNotificationConfiguration) GetNotificationConfiguration(req *GetNotificationConfigurationRequest, ctxs ..._context.Context) (GetNotificationConfigurationResponse, *_nethttp.Response, error) {
	res := &GetNotificationConfigurationResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/getNotificationConfiguration", ctxs...)
	return *res, httpRes, err
}

/*
PostGetNotificationConfigurationList Retrieve a list of existing notification subscription configurations.
This endpoint is used to retrieve the details of the configurations of all of the notification subscriptions in the marketplace of the executing user.
  - @param request Body - reference of interface{}).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetNotificationConfigurationListResponse
*/
func (a PlatformsNotificationConfiguration) GetNotificationConfigurationList(req interface{}, ctxs ..._context.Context) (GetNotificationConfigurationListResponse, *_nethttp.Response, error) {
	res := &GetNotificationConfigurationListResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/getNotificationConfigurationList", ctxs...)
	return *res, httpRes, err
}

/*
PostTestNotificationConfiguration Test an existing notification configuration.
This endpoint is used to test an existing notification subscription configuration. For each event type specified, a test notification will be generated and sent to the URL configured in the subscription specified.
  - @param request TestNotificationConfigurationRequest - reference of TestNotificationConfigurationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return TestNotificationConfigurationResponse
*/
func (a PlatformsNotificationConfiguration) TestNotificationConfiguration(req *TestNotificationConfigurationRequest, ctxs ..._context.Context) (TestNotificationConfigurationResponse, *_nethttp.Response, error) {
	res := &TestNotificationConfigurationResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/testNotificationConfiguration", ctxs...)
	return *res, httpRes, err
}

/*
PostUpdateNotificationConfiguration Update an existing notification subscription configuration.
This endpoint is used to update an existing notification subscription configuration. If updating the event types, all event types desired must be provided, otherwise the previous event type configuration will be overwritten.
  - @param request UpdateNotificationConfigurationRequest - reference of UpdateNotificationConfigurationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetNotificationConfigurationResponse
*/
func (a PlatformsNotificationConfiguration) UpdateNotificationConfiguration(req *UpdateNotificationConfigurationRequest, ctxs ..._context.Context) (GetNotificationConfigurationResponse, *_nethttp.Response, error) {
	res := &GetNotificationConfigurationResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/updateNotificationConfiguration", ctxs...)
	return *res, httpRes, err
}
