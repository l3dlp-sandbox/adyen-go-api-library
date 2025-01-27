/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v6/createAccountHolder ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount

import (
	_context "context"
	_nethttp "net/http"

	"github.com/adyen/adyen-go-api-library/v17/src/common"
)

// PlatformsAccount PlatformsAccount service
// Deprecated: Please migrate to the new Adyen For Platforms.
type PlatformsAccount common.Service

/*
PostCheckAccountHolder Request to perform verification for an account holder.
This endpoint allows to trigger the verification of the account holder earlier than it&#39;s required by the currently processed volume.
  - @param request PerformVerificationRequest - reference of PerformVerificationRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GenericResponse
*/
func (a PlatformsAccount) CheckAccountHolder(req *PerformVerificationRequest, ctxs ..._context.Context) (GenericResponse, *_nethttp.Response, error) {
	res := &GenericResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/checkAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostCloseAccount Close an existing account under an account holder.
This endpoint is used to close an existing account under an account holder. If an account is closed, it may not process transactions or have its funds paid out,and it may not be reopened. Any payments made to a closed account will be directed to the merchant&#39;s liable account.
  - @param request CloseAccountRequest - reference of CloseAccountRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return CloseAccountResponse
*/
func (a PlatformsAccount) CloseAccount(req *CloseAccountRequest, ctxs ..._context.Context) (CloseAccountResponse, *_nethttp.Response, error) {
	res := &CloseAccountResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/closeAccount", ctxs...)
	return *res, httpRes, err
}

/*
PostCloseAccountHolder Close an existing account holder.
This endpoint is used to close an existing account holder and its accounts. If an account holder is closed, it may not process transactions or pay out funds, and it may not be reopened. Any payments made to a closed account will be directed to the merchant&#39;s liable account.
  - @param request CloseAccountHolderRequest - reference of CloseAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return CloseAccountHolderResponse
*/
func (a PlatformsAccount) CloseAccountHolder(req *CloseAccountHolderRequest, ctxs ..._context.Context) (CloseAccountHolderResponse, *_nethttp.Response, error) {
	res := &CloseAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/closeAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostCreateAccount Create a new account under an existing account holder.
This endpoint is used to create an account under an existing account holder. An account holder may have multiple accounts.
  - @param request CreateAccountRequest - reference of CreateAccountRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return CreateAccountResponse
*/
func (a PlatformsAccount) CreateAccount(req *CreateAccountRequest, ctxs ..._context.Context) (CreateAccountResponse, *_nethttp.Response, error) {
	res := &CreateAccountResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/createAccount", ctxs...)
	return *res, httpRes, err
}

/*
PostCreateAccountHolder Create a new account holder.
This endpoint is used to create an account holder. Each account holder represents a single sub-merchant, and each sub-merchant must be represented by an account holder. Depending on the legal entity type, different details are required to be provided in the call to this endpoint.
  - @param request CreateAccountHolderRequest - reference of CreateAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return CreateAccountHolderResponse
*/
func (a PlatformsAccount) CreateAccountHolder(req *CreateAccountHolderRequest, ctxs ..._context.Context) (CreateAccountHolderResponse, *_nethttp.Response, error) {
	res := &CreateAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/createAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostDeleteBankAccounts Delete bank accounts of an existing account holder.
This endpoint is used to delete existing bank accounts from an account holder. For this, pass the &#x60;accountHolderCode&#x60; you got on the account holder creation, and one or more &#x60;bankAccountUUIDs&#x60; specifying bank accounts to delete.
  - @param request DeleteBankAccountRequest - reference of DeleteBankAccountRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GenericResponse
*/
func (a PlatformsAccount) DeleteBankAccounts(req *DeleteBankAccountRequest, ctxs ..._context.Context) (GenericResponse, *_nethttp.Response, error) {
	res := &GenericResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/deleteBankAccounts", ctxs...)
	return *res, httpRes, err
}

/*
PostDeletePayoutMethods Delete payout methods of an existing account holder.
This endpoint is used to delete existing payout method from an account holder. For this, pass the &#x60;accountHolderCode&#x60; you got on the account holder creation, and one or more &#x60;payoutMethodCodes&#x60; specifying payout methods to delete.
  - @param request DeletePayoutMethodRequest - reference of DeletePayoutMethodRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GenericResponse
*/
func (a PlatformsAccount) DeletePayoutMethods(req *DeletePayoutMethodRequest, ctxs ..._context.Context) (GenericResponse, *_nethttp.Response, error) {
	res := &GenericResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/deletePayoutMethods", ctxs...)
	return *res, httpRes, err
}

/*
PostDeleteShareholders Delete shareholders of an existing account holder.
This endpoint is used to delete existing shareholders from an account holder.
  - @param request DeleteShareholderRequest - reference of DeleteShareholderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GenericResponse
*/
func (a PlatformsAccount) DeleteShareholders(req *DeleteShareholderRequest, ctxs ..._context.Context) (GenericResponse, *_nethttp.Response, error) {
	res := &GenericResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/deleteShareholders", ctxs...)
	return *res, httpRes, err
}

/*
PostGetAccountHolder Retrieve the details of an account holder.
This endpoint is used to retrieve the details of an account holder.
  - @param request GetAccountHolderRequest - reference of GetAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetAccountHolderResponse
*/
func (a PlatformsAccount) GetAccountHolder(req *GetAccountHolderRequest, ctxs ..._context.Context) (GetAccountHolderResponse, *_nethttp.Response, error) {
	res := &GetAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/getAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostGetUploadedDocuments Retrieve the uploaded documents of an existing account holder.
This endpoint is used to retrieve documents previously uploaded for use in the KYC Verification of an account holder.  For further information regarding KYC Verification, please refer to [Verification checks](https://docs.adyen.com/platforms/onboarding-and-verification/verification-checks).
  - @param request GetUploadedDocumentsRequest - reference of GetUploadedDocumentsRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetUploadedDocumentsResponse
*/
func (a PlatformsAccount) GetUploadedDocuments(req *GetUploadedDocumentsRequest, ctxs ..._context.Context) (GetUploadedDocumentsResponse, *_nethttp.Response, error) {
	res := &GetUploadedDocumentsResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/getUploadedDocuments", ctxs...)
	return *res, httpRes, err
}

/*
PostSuspendAccountHolder Suspend an existing account holder.
This endpoint is used to suspend an existing account holder. If an account holder is suspended, it may not process transactions or pay out funds. Any payments made to a suspended account holder will be directed to the merchant&#39;s liable account.
  - @param request SuspendAccountHolderRequest - reference of SuspendAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return SuspendAccountHolderResponse
*/
func (a PlatformsAccount) SuspendAccountHolder(req *SuspendAccountHolderRequest, ctxs ..._context.Context) (SuspendAccountHolderResponse, *_nethttp.Response, error) {
	res := &SuspendAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/suspendAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostUnSuspendAccountHolder Reinstate a disabled account holder.
This endpoint is used to reinstate an existing account holder that has been suspended either through the &#x60;/suspendAccountHolder&#x60; endpoint or because a KYC deadline expired.  However, an account holder that has been suspended by Adyen because of KYC verification issues (indicated by a **FAILED** verification [&#x60;status&#x60;](https://docs.adyen.com/api-explorer/#/Account/latest/post/getAccountHolder__resParam_verification-accountHolder-checks-status)) cannot be reinstated through this endpoint.
  - @param request UnSuspendAccountHolderRequest - reference of UnSuspendAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return UnSuspendAccountHolderResponse
*/
func (a PlatformsAccount) UnSuspendAccountHolder(req *UnSuspendAccountHolderRequest, ctxs ..._context.Context) (UnSuspendAccountHolderResponse, *_nethttp.Response, error) {
	res := &UnSuspendAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/unSuspendAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostUpdateAccount Update an existing account under an account holder.
This endpoint is used to update the description or payout schedule of an existing account.
  - @param request UpdateAccountRequest - reference of UpdateAccountRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return UpdateAccountResponse
*/
func (a PlatformsAccount) UpdateAccount(req *UpdateAccountRequest, ctxs ..._context.Context) (UpdateAccountResponse, *_nethttp.Response, error) {
	res := &UpdateAccountResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/updateAccount", ctxs...)
	return *res, httpRes, err
}

/*
PostUpdateAccountHolder Update an existing account holder.
This endpoint is used to update the &#x60;accountHolderDetails&#x60; or &#x60;processingTier&#x60; of an account holder.  If updating the &#x60;accountHolderDetails&#x60;, only the details which have been provided will be updated. Other details will be left as-is with the exception of the following fields: * &#x60;accountHolderDetails.address&#x60; * &#x60;accountHolderDetails.fullPhoneNumber&#x60; * &#x60;accountHolderDetails.bankAccountDetails.BankAccountDetail&#x60; * &#x60;accountHolderDetails.businessDetails.shareholders.ShareholderContact&#x60;, which requires all fields necessary for validation (i.e. in order to update only the &#x60;accountHolderDetails.address.postalCode&#x60;, the fields &#x60;accountHolderDetails.address.country&#x60;, &#x60;.city&#x60;, &#x60;.street&#x60;, &#x60;.postalCode&#x60;, and possibly &#x60;.stateOrProvince&#x60; must be provided as well, so that the address can be properly validated).  Note that this endpoint can also be used to create new bank accounts. For this, provide details of a bank account without providing a &#x60;bankAccountUUID&#x60;.  Similarly, it can also be used to create new shareholders by providing details of a shareholder without providing a &#x60;shareholderCode&#x60;. &gt; The updating of the &#x60;metadata&#x60; field will overwite all of the existing account holder metadata. In order to update an existing metadata key-value pair, all otherkey-value pairs should be provided in order to not delete them.
  - @param request UpdateAccountHolderRequest - reference of UpdateAccountHolderRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return UpdateAccountHolderResponse
*/
func (a PlatformsAccount) UpdateAccountHolder(req *UpdateAccountHolderRequest, ctxs ..._context.Context) (UpdateAccountHolderResponse, *_nethttp.Response, error) {
	res := &UpdateAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/updateAccountHolder", ctxs...)
	return *res, httpRes, err
}

/*
PostUpdateAccountHolderState Update the state of an existing account holder.
This endpoint is used to disable or enable the processing or payout state of an account holder.  For more information about processing and payout states of an account holder, refer to [our documentation](https://docs.adyen.com/platforms).
  - @param request UpdateAccountHolderStateRequest - reference of UpdateAccountHolderStateRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return GetAccountHolderStatusResponse
*/
func (a PlatformsAccount) UpdateAccountHolderState(req *UpdateAccountHolderStateRequest, ctxs ..._context.Context) (GetAccountHolderStatusResponse, *_nethttp.Response, error) {
	res := &GetAccountHolderStatusResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/updateAccountHolderState", ctxs...)
	return *res, httpRes, err
}

/*
PostUploadDocument Upload a document for an existing account holder.
This endpoint is used to upload a document for use in the KYC verification of an account holder.  For further information regarding KYC Verification, please refer to [Verification checks](https://docs.adyen.com/platforms/onboarding-and-verification/verification-checks).
  - @param request UploadDocumentRequest - reference of UploadDocumentRequest).
  - @param ctxs ..._context.Context - optional, for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return UpdateAccountHolderResponse
*/
func (a PlatformsAccount) UploadDocument(req *UploadDocumentRequest, ctxs ..._context.Context) (UpdateAccountHolderResponse, *_nethttp.Response, error) {
	res := &UpdateAccountHolderResponse{}
	httpRes, err := a.Client.MakeHTTPPostRequest(req, res, a.BasePath()+"/uploadDocument", ctxs...)
	return *res, httpRes, err
}
