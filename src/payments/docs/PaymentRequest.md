# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountInfo** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AdditionalAmount** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**AdditionalData** | Pointer to **map[string]string** | This field contains additional data, which may be required for a particular payment request.  The &#x60;additionalData&#x60; object consists of entries, each of which includes the key and value. | [optional] 
**Amount** | [**Amount**](Amount.md) |  | 
**ApplicationInfo** | Pointer to [**ApplicationInfo**](ApplicationInfo.md) |  | [optional] 
**BankAccount** | Pointer to [**BankAccount**](BankAccount.md) |  | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**BrowserInfo** | Pointer to [**BrowserInfo**](BrowserInfo.md) |  | [optional] 
**CaptureDelayHours** | Pointer to **int32** | The delay between the authorisation and scheduled auto-capture, specified in hours. | [optional] 
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**DateOfBirth** | Pointer to **string** | The shopper&#39;s date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD | [optional] 
**DccQuote** | Pointer to [**ForexQuote**](ForexQuote.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**DeliveryDate** | Pointer to **time.Time** | The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00 | [optional] 
**DeviceFingerprint** | Pointer to **string** | A string containing the shopper&#39;s device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting). | [optional] 
**EntityType** | Pointer to **string** | The type of the entity the payment is processed for. | [optional] 
**FraudOffset** | Pointer to **int32** | An integer value that is added to the normal fraud score. The value can be either positive or negative. | [optional] 
**FundDestination** | Pointer to [**FundDestination**](FundDestination.md) |  | [optional] 
**FundSource** | Pointer to [**FundSource**](FundSource.md) |  | [optional] 
**FundingSource** | Pointer to **string** | The funding source that should be used when multiple sources are available. For Brazilian combo cards, by default the funding source is credit. To use debit, set this value to **debit**. | [optional] 
**Installments** | Pointer to [**Installments**](Installments.md) |  | [optional] 
**LocalizedShopperStatement** | Pointer to **map[string]string** | This field allows merchants to use dynamic shopper statement in local character sets. The local shopper statement field can be supplied in markets where localized merchant descriptors are used. Currently, Adyen only supports this in the Japanese market .The available character sets at the moment are: * Processing in Japan: **ja-Kana** The character set **ja-Kana** supports UTF-8 based Katakana and alphanumeric and special characters. Merchants should send the Katakana shopperStatement in full-width characters.  An example request would be: &gt; {   \&quot;shopperStatement\&quot; : \&quot;ADYEN - SELLER-A\&quot;,   \&quot;localizedShopperStatement\&quot; : {     \&quot;ja-Kana\&quot; : \&quot;ADYEN - セラーA\&quot;   } } We recommend merchants to always supply the field localizedShopperStatement in addition to the field shopperStatement.It is issuer dependent whether the localized shopper statement field is supported. In the case of non-domestic transactions (e.g. US-issued cards processed in JP) the field &#x60;shopperStatement&#x60; is used to modify the statement of the shopper. Adyen handles the complexity of ensuring the correct descriptors are assigned. | [optional] 
**Mandate** | Pointer to [**Mandate**](Mandate.md) |  | [optional] 
**Mcc** | Pointer to **string** | The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant. | [optional] 
**MerchantAccount** | **string** | The merchant account identifier, with which you want to process the transaction. | 
**MerchantOrderReference** | Pointer to **string** | This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. &gt; We strongly recommend you send the &#x60;merchantOrderReference&#x60; value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide &#x60;retry.orderAttemptNumber&#x60;, &#x60;retry.chainAttemptNumber&#x60;, and &#x60;retry.skipRetry&#x60; values in &#x60;PaymentRequest.additionalData&#x60;. | [optional] 
**MerchantRiskIndicator** | Pointer to [**MerchantRiskIndicator**](MerchantRiskIndicator.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata consists of entries, each of which includes a key and a value. Limits: * Maximum 20 key-value pairs per request. When exceeding, the \&quot;177\&quot; error occurs: \&quot;Metadata size exceeds limit\&quot;. * Maximum 20 characters per key. * Maximum 80 characters per value.  | [optional] 
**MpiData** | Pointer to [**ThreeDSecureData**](ThreeDSecureData.md) |  | [optional] 
**Nationality** | Pointer to **string** | The two-character country code of the shopper&#39;s nationality. | [optional] 
**OrderReference** | Pointer to **string** | When you are doing multiple partial (gift card) payments, this is the &#x60;pspReference&#x60; of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the &#x60;merchantOrderReference&#x60;instead. | [optional] 
**PlatformChargebackLogic** | Pointer to [**PlatformChargebackLogic**](PlatformChargebackLogic.md) |  | [optional] 
**Recurring** | Pointer to [**Recurring**](Recurring.md) |  | [optional] 
**RecurringProcessingModel** | Pointer to **string** | Defines a recurring payment type. Allowed values: * &#x60;Subscription&#x60; – A transaction for a fixed or variable amount, which follows a fixed schedule. * &#x60;CardOnFile&#x60; – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * &#x60;UnscheduledCardOnFile&#x60; – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder&#39;s balance drops below a certain amount.  | [optional] 
**Reference** | **string** | The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\&quot;-\&quot;). Maximum length: 80 characters. | 
**SelectedBrand** | Pointer to **string** | Some payment methods require defining a value for this field to specify how to process the transaction.  For the Bancontact payment method, it can be set to: * &#x60;maestro&#x60; (default), to be processed like a Maestro card, or * &#x60;bcmc&#x60;, to be processed like a Bancontact card. | [optional] 
**SelectedRecurringDetailReference** | Pointer to **string** | The &#x60;recurringDetailReference&#x60; you want to use for this payment. The value &#x60;LATEST&#x60; can be used to select the most recently stored recurring detail. | [optional] 
**SessionId** | Pointer to **string** | A session ID used to identify a payment session. | [optional] 
**ShopperEmail** | Pointer to **string** | The shopper&#39;s email address. We recommend that you provide this data, as it is used in velocity fraud checks. &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperEmail&#x60; for all browser-based and mobile implementations. | [optional] 
**ShopperIP** | Pointer to **string** | The shopper&#39;s IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). &gt; For 3D Secure 2 transactions, schemes require &#x60;shopperIP&#x60; for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://www.adyen.help/hc/en-us/requests/new). | [optional] 
**ShopperInteraction** | Pointer to **string** | Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * &#x60;Ecommerce&#x60; - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * &#x60;ContAuth&#x60; - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * &#x60;Moto&#x60; - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * &#x60;POS&#x60; - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal. | [optional] 
**ShopperLocale** | Pointer to **string** | The combination of a language code and a country code to specify the language to be used in the payment. | [optional] 
**ShopperName** | Pointer to [**Name**](Name.md) |  | [optional] 
**ShopperReference** | Pointer to **string** | Required for recurring payments.  Your reference to uniquely identify this shopper, for example user ID or account ID. Minimum length: 3 characters. &gt; Your reference must not include personally identifiable information (PII), for example name or email address. | [optional] 
**ShopperStatement** | Pointer to **string** | The text to be shown on the shopper&#39;s bank statement.  We recommend sending a maximum of 22 characters, otherwise banks might truncate the string.  Allowed characters: **a-z**, **A-Z**, **0-9**, spaces, and special characters **. , &#39; _ - ? + * /_**. | [optional] 
**SocialSecurityNumber** | Pointer to **string** | The shopper&#39;s social security number. | [optional] 
**Splits** | Pointer to [**[]Split**](Split.md) | An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/marketplaces-and-platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/add-manage-funds#split). | [optional] 
**Store** | Pointer to **string** | The ecommerce or point-of-sale store that is processing the payment. Used in [partner model integrations](https://docs.adyen.com/marketplaces-and-platforms/classic/platforms-for-partners#route-payments) for Adyen for Platforms. | [optional] 
**TelephoneNumber** | Pointer to **string** | The shopper&#39;s telephone number. | [optional] 
**ThreeDS2RequestData** | Pointer to [**ThreeDS2RequestData**](ThreeDS2RequestData.md) |  | [optional] 
**ThreeDSAuthenticationOnly** | Pointer to **bool** | If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. | [optional] [default to false]
**TotalsGroup** | Pointer to **string** | The reference value to aggregate sales totals in reporting. When not specified, the store field is used (if available). | [optional] 
**TrustedShopper** | Pointer to **bool** | Set to true if the payment should be routed to a trusted MID. | [optional] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest(amount Amount, merchantAccount string, reference string, ) *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountInfo

`func (o *PaymentRequest) GetAccountInfo() AccountInfo`

GetAccountInfo returns the AccountInfo field if non-nil, zero value otherwise.

### GetAccountInfoOk

`func (o *PaymentRequest) GetAccountInfoOk() (*AccountInfo, bool)`

GetAccountInfoOk returns a tuple with the AccountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfo

`func (o *PaymentRequest) SetAccountInfo(v AccountInfo)`

SetAccountInfo sets AccountInfo field to given value.

### HasAccountInfo

`func (o *PaymentRequest) HasAccountInfo() bool`

HasAccountInfo returns a boolean if a field has been set.

### GetAdditionalAmount

`func (o *PaymentRequest) GetAdditionalAmount() Amount`

GetAdditionalAmount returns the AdditionalAmount field if non-nil, zero value otherwise.

### GetAdditionalAmountOk

`func (o *PaymentRequest) GetAdditionalAmountOk() (*Amount, bool)`

GetAdditionalAmountOk returns a tuple with the AdditionalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAmount

`func (o *PaymentRequest) SetAdditionalAmount(v Amount)`

SetAdditionalAmount sets AdditionalAmount field to given value.

### HasAdditionalAmount

`func (o *PaymentRequest) HasAdditionalAmount() bool`

HasAdditionalAmount returns a boolean if a field has been set.

### GetAdditionalData

`func (o *PaymentRequest) GetAdditionalData() map[string]string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *PaymentRequest) GetAdditionalDataOk() (*map[string]string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *PaymentRequest) SetAdditionalData(v map[string]string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *PaymentRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentRequest) GetAmount() Amount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequest) GetAmountOk() (*Amount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequest) SetAmount(v Amount)`

SetAmount sets Amount field to given value.


### GetApplicationInfo

`func (o *PaymentRequest) GetApplicationInfo() ApplicationInfo`

GetApplicationInfo returns the ApplicationInfo field if non-nil, zero value otherwise.

### GetApplicationInfoOk

`func (o *PaymentRequest) GetApplicationInfoOk() (*ApplicationInfo, bool)`

GetApplicationInfoOk returns a tuple with the ApplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInfo

`func (o *PaymentRequest) SetApplicationInfo(v ApplicationInfo)`

SetApplicationInfo sets ApplicationInfo field to given value.

### HasApplicationInfo

`func (o *PaymentRequest) HasApplicationInfo() bool`

HasApplicationInfo returns a boolean if a field has been set.

### GetBankAccount

`func (o *PaymentRequest) GetBankAccount() BankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *PaymentRequest) GetBankAccountOk() (*BankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *PaymentRequest) SetBankAccount(v BankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *PaymentRequest) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.

### GetBillingAddress

`func (o *PaymentRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *PaymentRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *PaymentRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *PaymentRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetBrowserInfo

`func (o *PaymentRequest) GetBrowserInfo() BrowserInfo`

GetBrowserInfo returns the BrowserInfo field if non-nil, zero value otherwise.

### GetBrowserInfoOk

`func (o *PaymentRequest) GetBrowserInfoOk() (*BrowserInfo, bool)`

GetBrowserInfoOk returns a tuple with the BrowserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserInfo

`func (o *PaymentRequest) SetBrowserInfo(v BrowserInfo)`

SetBrowserInfo sets BrowserInfo field to given value.

### HasBrowserInfo

`func (o *PaymentRequest) HasBrowserInfo() bool`

HasBrowserInfo returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *PaymentRequest) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *PaymentRequest) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *PaymentRequest) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *PaymentRequest) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetCard

`func (o *PaymentRequest) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentRequest) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentRequest) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentRequest) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *PaymentRequest) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *PaymentRequest) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *PaymentRequest) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *PaymentRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDccQuote

`func (o *PaymentRequest) GetDccQuote() ForexQuote`

GetDccQuote returns the DccQuote field if non-nil, zero value otherwise.

### GetDccQuoteOk

`func (o *PaymentRequest) GetDccQuoteOk() (*ForexQuote, bool)`

GetDccQuoteOk returns a tuple with the DccQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccQuote

`func (o *PaymentRequest) SetDccQuote(v ForexQuote)`

SetDccQuote sets DccQuote field to given value.

### HasDccQuote

`func (o *PaymentRequest) HasDccQuote() bool`

HasDccQuote returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *PaymentRequest) GetDeliveryAddress() Address`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *PaymentRequest) GetDeliveryAddressOk() (*Address, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *PaymentRequest) SetDeliveryAddress(v Address)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *PaymentRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *PaymentRequest) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *PaymentRequest) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *PaymentRequest) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *PaymentRequest) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetDeviceFingerprint

`func (o *PaymentRequest) GetDeviceFingerprint() string`

GetDeviceFingerprint returns the DeviceFingerprint field if non-nil, zero value otherwise.

### GetDeviceFingerprintOk

`func (o *PaymentRequest) GetDeviceFingerprintOk() (*string, bool)`

GetDeviceFingerprintOk returns a tuple with the DeviceFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprint

`func (o *PaymentRequest) SetDeviceFingerprint(v string)`

SetDeviceFingerprint sets DeviceFingerprint field to given value.

### HasDeviceFingerprint

`func (o *PaymentRequest) HasDeviceFingerprint() bool`

HasDeviceFingerprint returns a boolean if a field has been set.

### GetEntityType

`func (o *PaymentRequest) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *PaymentRequest) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *PaymentRequest) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *PaymentRequest) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetFraudOffset

`func (o *PaymentRequest) GetFraudOffset() int32`

GetFraudOffset returns the FraudOffset field if non-nil, zero value otherwise.

### GetFraudOffsetOk

`func (o *PaymentRequest) GetFraudOffsetOk() (*int32, bool)`

GetFraudOffsetOk returns a tuple with the FraudOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudOffset

`func (o *PaymentRequest) SetFraudOffset(v int32)`

SetFraudOffset sets FraudOffset field to given value.

### HasFraudOffset

`func (o *PaymentRequest) HasFraudOffset() bool`

HasFraudOffset returns a boolean if a field has been set.

### GetFundDestination

`func (o *PaymentRequest) GetFundDestination() FundDestination`

GetFundDestination returns the FundDestination field if non-nil, zero value otherwise.

### GetFundDestinationOk

`func (o *PaymentRequest) GetFundDestinationOk() (*FundDestination, bool)`

GetFundDestinationOk returns a tuple with the FundDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundDestination

`func (o *PaymentRequest) SetFundDestination(v FundDestination)`

SetFundDestination sets FundDestination field to given value.

### HasFundDestination

`func (o *PaymentRequest) HasFundDestination() bool`

HasFundDestination returns a boolean if a field has been set.

### GetFundSource

`func (o *PaymentRequest) GetFundSource() FundSource`

GetFundSource returns the FundSource field if non-nil, zero value otherwise.

### GetFundSourceOk

`func (o *PaymentRequest) GetFundSourceOk() (*FundSource, bool)`

GetFundSourceOk returns a tuple with the FundSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundSource

`func (o *PaymentRequest) SetFundSource(v FundSource)`

SetFundSource sets FundSource field to given value.

### HasFundSource

`func (o *PaymentRequest) HasFundSource() bool`

HasFundSource returns a boolean if a field has been set.

### GetFundingSource

`func (o *PaymentRequest) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *PaymentRequest) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *PaymentRequest) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *PaymentRequest) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetInstallments

`func (o *PaymentRequest) GetInstallments() Installments`

GetInstallments returns the Installments field if non-nil, zero value otherwise.

### GetInstallmentsOk

`func (o *PaymentRequest) GetInstallmentsOk() (*Installments, bool)`

GetInstallmentsOk returns a tuple with the Installments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallments

`func (o *PaymentRequest) SetInstallments(v Installments)`

SetInstallments sets Installments field to given value.

### HasInstallments

`func (o *PaymentRequest) HasInstallments() bool`

HasInstallments returns a boolean if a field has been set.

### GetLocalizedShopperStatement

`func (o *PaymentRequest) GetLocalizedShopperStatement() map[string]string`

GetLocalizedShopperStatement returns the LocalizedShopperStatement field if non-nil, zero value otherwise.

### GetLocalizedShopperStatementOk

`func (o *PaymentRequest) GetLocalizedShopperStatementOk() (*map[string]string, bool)`

GetLocalizedShopperStatementOk returns a tuple with the LocalizedShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedShopperStatement

`func (o *PaymentRequest) SetLocalizedShopperStatement(v map[string]string)`

SetLocalizedShopperStatement sets LocalizedShopperStatement field to given value.

### HasLocalizedShopperStatement

`func (o *PaymentRequest) HasLocalizedShopperStatement() bool`

HasLocalizedShopperStatement returns a boolean if a field has been set.

### GetMandate

`func (o *PaymentRequest) GetMandate() Mandate`

GetMandate returns the Mandate field if non-nil, zero value otherwise.

### GetMandateOk

`func (o *PaymentRequest) GetMandateOk() (*Mandate, bool)`

GetMandateOk returns a tuple with the Mandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandate

`func (o *PaymentRequest) SetMandate(v Mandate)`

SetMandate sets Mandate field to given value.

### HasMandate

`func (o *PaymentRequest) HasMandate() bool`

HasMandate returns a boolean if a field has been set.

### GetMcc

`func (o *PaymentRequest) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *PaymentRequest) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *PaymentRequest) SetMcc(v string)`

SetMcc sets Mcc field to given value.

### HasMcc

`func (o *PaymentRequest) HasMcc() bool`

HasMcc returns a boolean if a field has been set.

### GetMerchantAccount

`func (o *PaymentRequest) GetMerchantAccount() string`

GetMerchantAccount returns the MerchantAccount field if non-nil, zero value otherwise.

### GetMerchantAccountOk

`func (o *PaymentRequest) GetMerchantAccountOk() (*string, bool)`

GetMerchantAccountOk returns a tuple with the MerchantAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccount

`func (o *PaymentRequest) SetMerchantAccount(v string)`

SetMerchantAccount sets MerchantAccount field to given value.


### GetMerchantOrderReference

`func (o *PaymentRequest) GetMerchantOrderReference() string`

GetMerchantOrderReference returns the MerchantOrderReference field if non-nil, zero value otherwise.

### GetMerchantOrderReferenceOk

`func (o *PaymentRequest) GetMerchantOrderReferenceOk() (*string, bool)`

GetMerchantOrderReferenceOk returns a tuple with the MerchantOrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOrderReference

`func (o *PaymentRequest) SetMerchantOrderReference(v string)`

SetMerchantOrderReference sets MerchantOrderReference field to given value.

### HasMerchantOrderReference

`func (o *PaymentRequest) HasMerchantOrderReference() bool`

HasMerchantOrderReference returns a boolean if a field has been set.

### GetMerchantRiskIndicator

`func (o *PaymentRequest) GetMerchantRiskIndicator() MerchantRiskIndicator`

GetMerchantRiskIndicator returns the MerchantRiskIndicator field if non-nil, zero value otherwise.

### GetMerchantRiskIndicatorOk

`func (o *PaymentRequest) GetMerchantRiskIndicatorOk() (*MerchantRiskIndicator, bool)`

GetMerchantRiskIndicatorOk returns a tuple with the MerchantRiskIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRiskIndicator

`func (o *PaymentRequest) SetMerchantRiskIndicator(v MerchantRiskIndicator)`

SetMerchantRiskIndicator sets MerchantRiskIndicator field to given value.

### HasMerchantRiskIndicator

`func (o *PaymentRequest) HasMerchantRiskIndicator() bool`

HasMerchantRiskIndicator returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMpiData

`func (o *PaymentRequest) GetMpiData() ThreeDSecureData`

GetMpiData returns the MpiData field if non-nil, zero value otherwise.

### GetMpiDataOk

`func (o *PaymentRequest) GetMpiDataOk() (*ThreeDSecureData, bool)`

GetMpiDataOk returns a tuple with the MpiData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpiData

`func (o *PaymentRequest) SetMpiData(v ThreeDSecureData)`

SetMpiData sets MpiData field to given value.

### HasMpiData

`func (o *PaymentRequest) HasMpiData() bool`

HasMpiData returns a boolean if a field has been set.

### GetNationality

`func (o *PaymentRequest) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *PaymentRequest) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *PaymentRequest) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *PaymentRequest) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetOrderReference

`func (o *PaymentRequest) GetOrderReference() string`

GetOrderReference returns the OrderReference field if non-nil, zero value otherwise.

### GetOrderReferenceOk

`func (o *PaymentRequest) GetOrderReferenceOk() (*string, bool)`

GetOrderReferenceOk returns a tuple with the OrderReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderReference

`func (o *PaymentRequest) SetOrderReference(v string)`

SetOrderReference sets OrderReference field to given value.

### HasOrderReference

`func (o *PaymentRequest) HasOrderReference() bool`

HasOrderReference returns a boolean if a field has been set.

### GetPlatformChargebackLogic

`func (o *PaymentRequest) GetPlatformChargebackLogic() PlatformChargebackLogic`

GetPlatformChargebackLogic returns the PlatformChargebackLogic field if non-nil, zero value otherwise.

### GetPlatformChargebackLogicOk

`func (o *PaymentRequest) GetPlatformChargebackLogicOk() (*PlatformChargebackLogic, bool)`

GetPlatformChargebackLogicOk returns a tuple with the PlatformChargebackLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformChargebackLogic

`func (o *PaymentRequest) SetPlatformChargebackLogic(v PlatformChargebackLogic)`

SetPlatformChargebackLogic sets PlatformChargebackLogic field to given value.

### HasPlatformChargebackLogic

`func (o *PaymentRequest) HasPlatformChargebackLogic() bool`

HasPlatformChargebackLogic returns a boolean if a field has been set.

### GetRecurring

`func (o *PaymentRequest) GetRecurring() Recurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *PaymentRequest) GetRecurringOk() (*Recurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *PaymentRequest) SetRecurring(v Recurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *PaymentRequest) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetRecurringProcessingModel

`func (o *PaymentRequest) GetRecurringProcessingModel() string`

GetRecurringProcessingModel returns the RecurringProcessingModel field if non-nil, zero value otherwise.

### GetRecurringProcessingModelOk

`func (o *PaymentRequest) GetRecurringProcessingModelOk() (*string, bool)`

GetRecurringProcessingModelOk returns a tuple with the RecurringProcessingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringProcessingModel

`func (o *PaymentRequest) SetRecurringProcessingModel(v string)`

SetRecurringProcessingModel sets RecurringProcessingModel field to given value.

### HasRecurringProcessingModel

`func (o *PaymentRequest) HasRecurringProcessingModel() bool`

HasRecurringProcessingModel returns a boolean if a field has been set.

### GetReference

`func (o *PaymentRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetSelectedBrand

`func (o *PaymentRequest) GetSelectedBrand() string`

GetSelectedBrand returns the SelectedBrand field if non-nil, zero value otherwise.

### GetSelectedBrandOk

`func (o *PaymentRequest) GetSelectedBrandOk() (*string, bool)`

GetSelectedBrandOk returns a tuple with the SelectedBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedBrand

`func (o *PaymentRequest) SetSelectedBrand(v string)`

SetSelectedBrand sets SelectedBrand field to given value.

### HasSelectedBrand

`func (o *PaymentRequest) HasSelectedBrand() bool`

HasSelectedBrand returns a boolean if a field has been set.

### GetSelectedRecurringDetailReference

`func (o *PaymentRequest) GetSelectedRecurringDetailReference() string`

GetSelectedRecurringDetailReference returns the SelectedRecurringDetailReference field if non-nil, zero value otherwise.

### GetSelectedRecurringDetailReferenceOk

`func (o *PaymentRequest) GetSelectedRecurringDetailReferenceOk() (*string, bool)`

GetSelectedRecurringDetailReferenceOk returns a tuple with the SelectedRecurringDetailReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRecurringDetailReference

`func (o *PaymentRequest) SetSelectedRecurringDetailReference(v string)`

SetSelectedRecurringDetailReference sets SelectedRecurringDetailReference field to given value.

### HasSelectedRecurringDetailReference

`func (o *PaymentRequest) HasSelectedRecurringDetailReference() bool`

HasSelectedRecurringDetailReference returns a boolean if a field has been set.

### GetSessionId

`func (o *PaymentRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PaymentRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PaymentRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PaymentRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetShopperEmail

`func (o *PaymentRequest) GetShopperEmail() string`

GetShopperEmail returns the ShopperEmail field if non-nil, zero value otherwise.

### GetShopperEmailOk

`func (o *PaymentRequest) GetShopperEmailOk() (*string, bool)`

GetShopperEmailOk returns a tuple with the ShopperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperEmail

`func (o *PaymentRequest) SetShopperEmail(v string)`

SetShopperEmail sets ShopperEmail field to given value.

### HasShopperEmail

`func (o *PaymentRequest) HasShopperEmail() bool`

HasShopperEmail returns a boolean if a field has been set.

### GetShopperIP

`func (o *PaymentRequest) GetShopperIP() string`

GetShopperIP returns the ShopperIP field if non-nil, zero value otherwise.

### GetShopperIPOk

`func (o *PaymentRequest) GetShopperIPOk() (*string, bool)`

GetShopperIPOk returns a tuple with the ShopperIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperIP

`func (o *PaymentRequest) SetShopperIP(v string)`

SetShopperIP sets ShopperIP field to given value.

### HasShopperIP

`func (o *PaymentRequest) HasShopperIP() bool`

HasShopperIP returns a boolean if a field has been set.

### GetShopperInteraction

`func (o *PaymentRequest) GetShopperInteraction() string`

GetShopperInteraction returns the ShopperInteraction field if non-nil, zero value otherwise.

### GetShopperInteractionOk

`func (o *PaymentRequest) GetShopperInteractionOk() (*string, bool)`

GetShopperInteractionOk returns a tuple with the ShopperInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperInteraction

`func (o *PaymentRequest) SetShopperInteraction(v string)`

SetShopperInteraction sets ShopperInteraction field to given value.

### HasShopperInteraction

`func (o *PaymentRequest) HasShopperInteraction() bool`

HasShopperInteraction returns a boolean if a field has been set.

### GetShopperLocale

`func (o *PaymentRequest) GetShopperLocale() string`

GetShopperLocale returns the ShopperLocale field if non-nil, zero value otherwise.

### GetShopperLocaleOk

`func (o *PaymentRequest) GetShopperLocaleOk() (*string, bool)`

GetShopperLocaleOk returns a tuple with the ShopperLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperLocale

`func (o *PaymentRequest) SetShopperLocale(v string)`

SetShopperLocale sets ShopperLocale field to given value.

### HasShopperLocale

`func (o *PaymentRequest) HasShopperLocale() bool`

HasShopperLocale returns a boolean if a field has been set.

### GetShopperName

`func (o *PaymentRequest) GetShopperName() Name`

GetShopperName returns the ShopperName field if non-nil, zero value otherwise.

### GetShopperNameOk

`func (o *PaymentRequest) GetShopperNameOk() (*Name, bool)`

GetShopperNameOk returns a tuple with the ShopperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperName

`func (o *PaymentRequest) SetShopperName(v Name)`

SetShopperName sets ShopperName field to given value.

### HasShopperName

`func (o *PaymentRequest) HasShopperName() bool`

HasShopperName returns a boolean if a field has been set.

### GetShopperReference

`func (o *PaymentRequest) GetShopperReference() string`

GetShopperReference returns the ShopperReference field if non-nil, zero value otherwise.

### GetShopperReferenceOk

`func (o *PaymentRequest) GetShopperReferenceOk() (*string, bool)`

GetShopperReferenceOk returns a tuple with the ShopperReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperReference

`func (o *PaymentRequest) SetShopperReference(v string)`

SetShopperReference sets ShopperReference field to given value.

### HasShopperReference

`func (o *PaymentRequest) HasShopperReference() bool`

HasShopperReference returns a boolean if a field has been set.

### GetShopperStatement

`func (o *PaymentRequest) GetShopperStatement() string`

GetShopperStatement returns the ShopperStatement field if non-nil, zero value otherwise.

### GetShopperStatementOk

`func (o *PaymentRequest) GetShopperStatementOk() (*string, bool)`

GetShopperStatementOk returns a tuple with the ShopperStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperStatement

`func (o *PaymentRequest) SetShopperStatement(v string)`

SetShopperStatement sets ShopperStatement field to given value.

### HasShopperStatement

`func (o *PaymentRequest) HasShopperStatement() bool`

HasShopperStatement returns a boolean if a field has been set.

### GetSocialSecurityNumber

`func (o *PaymentRequest) GetSocialSecurityNumber() string`

GetSocialSecurityNumber returns the SocialSecurityNumber field if non-nil, zero value otherwise.

### GetSocialSecurityNumberOk

`func (o *PaymentRequest) GetSocialSecurityNumberOk() (*string, bool)`

GetSocialSecurityNumberOk returns a tuple with the SocialSecurityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialSecurityNumber

`func (o *PaymentRequest) SetSocialSecurityNumber(v string)`

SetSocialSecurityNumber sets SocialSecurityNumber field to given value.

### HasSocialSecurityNumber

`func (o *PaymentRequest) HasSocialSecurityNumber() bool`

HasSocialSecurityNumber returns a boolean if a field has been set.

### GetSplits

`func (o *PaymentRequest) GetSplits() []Split`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *PaymentRequest) GetSplitsOk() (*[]Split, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *PaymentRequest) SetSplits(v []Split)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *PaymentRequest) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetStore

`func (o *PaymentRequest) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentRequest) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentRequest) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTelephoneNumber

`func (o *PaymentRequest) GetTelephoneNumber() string`

GetTelephoneNumber returns the TelephoneNumber field if non-nil, zero value otherwise.

### GetTelephoneNumberOk

`func (o *PaymentRequest) GetTelephoneNumberOk() (*string, bool)`

GetTelephoneNumberOk returns a tuple with the TelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNumber

`func (o *PaymentRequest) SetTelephoneNumber(v string)`

SetTelephoneNumber sets TelephoneNumber field to given value.

### HasTelephoneNumber

`func (o *PaymentRequest) HasTelephoneNumber() bool`

HasTelephoneNumber returns a boolean if a field has been set.

### GetThreeDS2RequestData

`func (o *PaymentRequest) GetThreeDS2RequestData() ThreeDS2RequestData`

GetThreeDS2RequestData returns the ThreeDS2RequestData field if non-nil, zero value otherwise.

### GetThreeDS2RequestDataOk

`func (o *PaymentRequest) GetThreeDS2RequestDataOk() (*ThreeDS2RequestData, bool)`

GetThreeDS2RequestDataOk returns a tuple with the ThreeDS2RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDS2RequestData

`func (o *PaymentRequest) SetThreeDS2RequestData(v ThreeDS2RequestData)`

SetThreeDS2RequestData sets ThreeDS2RequestData field to given value.

### HasThreeDS2RequestData

`func (o *PaymentRequest) HasThreeDS2RequestData() bool`

HasThreeDS2RequestData returns a boolean if a field has been set.

### GetThreeDSAuthenticationOnly

`func (o *PaymentRequest) GetThreeDSAuthenticationOnly() bool`

GetThreeDSAuthenticationOnly returns the ThreeDSAuthenticationOnly field if non-nil, zero value otherwise.

### GetThreeDSAuthenticationOnlyOk

`func (o *PaymentRequest) GetThreeDSAuthenticationOnlyOk() (*bool, bool)`

GetThreeDSAuthenticationOnlyOk returns a tuple with the ThreeDSAuthenticationOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSAuthenticationOnly

`func (o *PaymentRequest) SetThreeDSAuthenticationOnly(v bool)`

SetThreeDSAuthenticationOnly sets ThreeDSAuthenticationOnly field to given value.

### HasThreeDSAuthenticationOnly

`func (o *PaymentRequest) HasThreeDSAuthenticationOnly() bool`

HasThreeDSAuthenticationOnly returns a boolean if a field has been set.

### GetTotalsGroup

`func (o *PaymentRequest) GetTotalsGroup() string`

GetTotalsGroup returns the TotalsGroup field if non-nil, zero value otherwise.

### GetTotalsGroupOk

`func (o *PaymentRequest) GetTotalsGroupOk() (*string, bool)`

GetTotalsGroupOk returns a tuple with the TotalsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsGroup

`func (o *PaymentRequest) SetTotalsGroup(v string)`

SetTotalsGroup sets TotalsGroup field to given value.

### HasTotalsGroup

`func (o *PaymentRequest) HasTotalsGroup() bool`

HasTotalsGroup returns a boolean if a field has been set.

### GetTrustedShopper

`func (o *PaymentRequest) GetTrustedShopper() bool`

GetTrustedShopper returns the TrustedShopper field if non-nil, zero value otherwise.

### GetTrustedShopperOk

`func (o *PaymentRequest) GetTrustedShopperOk() (*bool, bool)`

GetTrustedShopperOk returns a tuple with the TrustedShopper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedShopper

`func (o *PaymentRequest) SetTrustedShopper(v bool)`

SetTrustedShopper sets TrustedShopper field to given value.

### HasTrustedShopper

`func (o *PaymentRequest) HasTrustedShopper() bool`

HasTrustedShopper returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

