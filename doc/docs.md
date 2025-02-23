# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gomoneypb/v1/account.proto](#gomoneypb_v1_account-proto)
    - [Account](#gomoneypb-v1-Account)
    - [Account.ExtraEntry](#gomoneypb-v1-Account-ExtraEntry)
  
- [gomoneypb/accounts/v1/accounts.proto](#gomoneypb_accounts_v1_accounts-proto)
    - [CreateAccountRequest](#gomoneypb-accounts-v1-CreateAccountRequest)
    - [CreateAccountResponse](#gomoneypb-accounts-v1-CreateAccountResponse)
    - [ListAccountsRequest](#gomoneypb-accounts-v1-ListAccountsRequest)
    - [ListAccountsResponse](#gomoneypb-accounts-v1-ListAccountsResponse)
    - [ListAccountsResponse.AccountItem](#gomoneypb-accounts-v1-ListAccountsResponse-AccountItem)
  
    - [AccountsService](#gomoneypb-accounts-v1-AccountsService)
  
- [gomoneypb/currency/v1/currency.proto](#gomoneypb_currency_v1_currency-proto)
    - [ExchangeRequest](#gomoneypb-currency-v1-ExchangeRequest)
    - [ExchangeResponse](#gomoneypb-currency-v1-ExchangeResponse)
  
    - [CurrencyService](#gomoneypb-currency-v1-CurrencyService)
  
- [gomoneypb/users/v1/users.proto](#gomoneypb_users_v1_users-proto)
    - [LoginRequest](#gomoneypb-users-v1-LoginRequest)
    - [LoginResponse](#gomoneypb-users-v1-LoginResponse)
  
    - [UsersService](#gomoneypb-users-v1-UsersService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="gomoneypb_v1_account-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/v1/account.proto



<a name="gomoneypb-v1-Account"></a>

### Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| currency | [string](#string) |  |  |
| currency_balance | [string](#string) |  |  |
| extra | [Account.ExtraEntry](#gomoneypb-v1-Account-ExtraEntry) | repeated |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional |  |
| type | [string](#string) |  |  |






<a name="gomoneypb-v1-Account-ExtraEntry"></a>

### Account.ExtraEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="gomoneypb_accounts_v1_accounts-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/accounts/v1/accounts.proto



<a name="gomoneypb-accounts-v1-CreateAccountRequest"></a>

### CreateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [gomoneypb.v1.Account](#gomoneypb-v1-Account) |  |  |






<a name="gomoneypb-accounts-v1-CreateAccountResponse"></a>

### CreateAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [gomoneypb.v1.Account](#gomoneypb-v1-Account) |  |  |






<a name="gomoneypb-accounts-v1-ListAccountsRequest"></a>

### ListAccountsRequest







<a name="gomoneypb-accounts-v1-ListAccountsResponse"></a>

### ListAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [ListAccountsResponse.AccountItem](#gomoneypb-accounts-v1-ListAccountsResponse-AccountItem) | repeated |  |






<a name="gomoneypb-accounts-v1-ListAccountsResponse-AccountItem"></a>

### ListAccountsResponse.AccountItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [gomoneypb.v1.Account](#gomoneypb-v1-Account) |  |  |





 

 

 


<a name="gomoneypb-accounts-v1-AccountsService"></a>

### AccountsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAccount | [CreateAccountRequest](#gomoneypb-accounts-v1-CreateAccountRequest) | [CreateAccountResponse](#gomoneypb-accounts-v1-CreateAccountResponse) |  |
| ListAccounts | [ListAccountsRequest](#gomoneypb-accounts-v1-ListAccountsRequest) | [ListAccountsResponse](#gomoneypb-accounts-v1-ListAccountsResponse) |  |

 



<a name="gomoneypb_currency_v1_currency-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/currency/v1/currency.proto



<a name="gomoneypb-currency-v1-ExchangeRequest"></a>

### ExchangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from_currency | [string](#string) |  |  |
| to_currency | [string](#string) |  |  |
| amount | [string](#string) |  |  |






<a name="gomoneypb-currency-v1-ExchangeResponse"></a>

### ExchangeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | [string](#string) |  |  |





 

 

 


<a name="gomoneypb-currency-v1-CurrencyService"></a>

### CurrencyService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Exchange | [ExchangeRequest](#gomoneypb-currency-v1-ExchangeRequest) | [ExchangeResponse](#gomoneypb-currency-v1-ExchangeResponse) |  |

 



<a name="gomoneypb_users_v1_users-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/users/v1/users.proto



<a name="gomoneypb-users-v1-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="gomoneypb-users-v1-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |





 

 

 


<a name="gomoneypb-users-v1-UsersService"></a>

### UsersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#gomoneypb-users-v1-LoginRequest) | [LoginResponse](#gomoneypb-users-v1-LoginResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

