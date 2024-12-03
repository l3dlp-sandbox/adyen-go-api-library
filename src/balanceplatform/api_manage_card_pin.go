/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
    "net/http"
    "net/url"
    "strings"
    "github.com/adyen/adyen-go-api-library/v16/src/common"
)

// ManageCardPINApi service
type ManageCardPINApi common.Service

// All parameters accepted by ManageCardPINApi.ChangeCardPin
type ManageCardPINApiChangeCardPinInput struct {
	pinChangeRequest *PinChangeRequest
}

func (r ManageCardPINApiChangeCardPinInput) PinChangeRequest(pinChangeRequest PinChangeRequest) ManageCardPINApiChangeCardPinInput {
	r.pinChangeRequest = &pinChangeRequest
	return r
}


/*
Prepare a request for ChangeCardPin

@return ManageCardPINApiChangeCardPinInput
*/
func (a *ManageCardPINApi) ChangeCardPinInput() ManageCardPINApiChangeCardPinInput {
	return ManageCardPINApiChangeCardPinInput{
	}
}

/*
ChangeCardPin Change a card PIN

Changes the personal identification number (PIN) of a specified card.

To make this request, your API credential must have the following role:
* Bank Issuing PIN Change Webservice role

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageCardPINApiChangeCardPinInput - Request parameters, see ChangeCardPinInput
@return PinChangeResponse, *http.Response, error
*/
func (a *ManageCardPINApi) ChangeCardPin(ctx context.Context, r ManageCardPINApiChangeCardPinInput) (PinChangeResponse, *http.Response, error) {
    res := &PinChangeResponse{}
	path := "/pins/change"
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.pinChangeRequest,
        res,
        http.MethodPost,
        a.BasePath()+path,
        queryParams,
        headerParams,
    )

    if httpRes == nil {
        return *res, httpRes, err
    }

    var serviceError common.RestServiceError
                        if httpRes.StatusCode == 401 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 403 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 422 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 500 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }

    return *res, httpRes, err
}


// All parameters accepted by ManageCardPINApi.PublicKey
type ManageCardPINApiPublicKeyInput struct {
	purpose *string
	format *string
}

// The purpose of the public key.  Possible values: **pinChange**, **pinReveal**, **panReveal**.  Default value: **pinReveal**.
func (r ManageCardPINApiPublicKeyInput) Purpose(purpose string) ManageCardPINApiPublicKeyInput {
	r.purpose = &purpose
	return r
}

// The encoding format of public key.  Possible values: **jwk**, **pem**.  Default value: **pem**.
func (r ManageCardPINApiPublicKeyInput) Format(format string) ManageCardPINApiPublicKeyInput {
	r.format = &format
	return r
}


/*
Prepare a request for PublicKey

@return ManageCardPINApiPublicKeyInput
*/
func (a *ManageCardPINApi) PublicKeyInput() ManageCardPINApiPublicKeyInput {
	return ManageCardPINApiPublicKeyInput{
	}
}

/*
PublicKey Get an RSA public key

Get an [RSA](https://en.wikipedia.org/wiki/RSA_(cryptosystem)) public key to encrypt or decrypt card data.

 You need the RSA public key to generate the `encryptedKey` required to:
- [Change a PIN](https://docs.adyen.com/api-explorer/balanceplatform/2/post/pins/change).
- [Reveal a PIN](https://docs.adyen.com/api-explorer/balanceplatform/2/post/pins/reveal).
- [Reveal a PAN](https://docs.adyen.com/api-explorer/balanceplatform/2/post/paymentInstruments/reveal).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageCardPINApiPublicKeyInput - Request parameters, see PublicKeyInput
@return PublicKeyResponse, *http.Response, error
*/
func (a *ManageCardPINApi) PublicKey(ctx context.Context, r ManageCardPINApiPublicKeyInput) (PublicKeyResponse, *http.Response, error) {
    res := &PublicKeyResponse{}
	path := "/publicKey"
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    if r.purpose != nil {
        common.ParameterAddToQuery(queryParams, "purpose", r.purpose, "")
    }
    if r.format != nil {
        common.ParameterAddToQuery(queryParams, "format", r.format, "")
    }
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        nil,
        res,
        http.MethodGet,
        a.BasePath()+path,
        queryParams,
        headerParams,
    )

    if httpRes == nil {
        return *res, httpRes, err
    }

    var serviceError common.RestServiceError
                        if httpRes.StatusCode == 401 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 403 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 422 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 500 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }

    return *res, httpRes, err
}


// All parameters accepted by ManageCardPINApi.RevealCardPin
type ManageCardPINApiRevealCardPinInput struct {
	revealPinRequest *RevealPinRequest
}

func (r ManageCardPINApiRevealCardPinInput) RevealPinRequest(revealPinRequest RevealPinRequest) ManageCardPINApiRevealCardPinInput {
	r.revealPinRequest = &revealPinRequest
	return r
}


/*
Prepare a request for RevealCardPin

@return ManageCardPINApiRevealCardPinInput
*/
func (a *ManageCardPINApi) RevealCardPinInput() ManageCardPINApiRevealCardPinInput {
	return ManageCardPINApiRevealCardPinInput{
	}
}

/*
RevealCardPin Reveal a card PIN

Returns an encrypted PIN block that contains the PIN of a specified card. You can use the decrypted data to reveal the PIN in your user interface.

To make this request, your API credential must have the following role:
* Bank Issuing PIN Reveal Webservice role

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r ManageCardPINApiRevealCardPinInput - Request parameters, see RevealCardPinInput
@return RevealPinResponse, *http.Response, error
*/
func (a *ManageCardPINApi) RevealCardPin(ctx context.Context, r ManageCardPINApiRevealCardPinInput) (RevealPinResponse, *http.Response, error) {
    res := &RevealPinResponse{}
	path := "/pins/reveal"
    queryParams := url.Values{}
    headerParams := make(map[string]string)
    httpRes, err := common.SendAPIRequest(
        ctx,
        a.Client,
        r.revealPinRequest,
        res,
        http.MethodPost,
        a.BasePath()+path,
        queryParams,
        headerParams,
    )

    if httpRes == nil {
        return *res, httpRes, err
    }

    var serviceError common.RestServiceError
                        if httpRes.StatusCode == 401 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 403 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 422 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }
                        if httpRes.StatusCode == 500 {
                            body, _ := ioutil.ReadAll(httpRes.Body)
                            decodeError := json.Unmarshal([]byte(body), &serviceError)
                            if decodeError != nil {
                                return *res, httpRes, decodeError
                            }
                            return *res, httpRes, serviceError
                        }

    return *res, httpRes, err
}

