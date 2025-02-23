# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [gomoneypb/v1/account.proto](#gomoneypb_v1_account-proto)
    - [Account](#gomoneypb-v1-Account)
    - [Account.ExtraEntry](#gomoneypb-v1-Account-ExtraEntry)
  
    - [AccountType](#gomoneypb-v1-AccountType)
  
- [gomoneypb/accounts/v1/accounts.proto](#gomoneypb_accounts_v1_accounts-proto)
    - [CreateAccountRequest](#gomoneypb-accounts-v1-CreateAccountRequest)
    - [CreateAccountResponse](#gomoneypb-accounts-v1-CreateAccountResponse)
    - [ListAccountsRequest](#gomoneypb-accounts-v1-ListAccountsRequest)
    - [ListAccountsResponse](#gomoneypb-accounts-v1-ListAccountsResponse)
    - [ListAccountsResponse.AccountItem](#gomoneypb-accounts-v1-ListAccountsResponse-AccountItem)
    - [UpdateAccountRequest](#gomoneypb-accounts-v1-UpdateAccountRequest)
    - [UpdateAccountRequest.ExtraEntry](#gomoneypb-accounts-v1-UpdateAccountRequest-ExtraEntry)
    - [UpdateAccountResponse](#gomoneypb-accounts-v1-UpdateAccountResponse)
  
    - [AccountsService](#gomoneypb-accounts-v1-AccountsService)
  
- [gomoneypb/currency/v1/currency.proto](#gomoneypb_currency_v1_currency-proto)
    - [ExchangeRequest](#gomoneypb-currency-v1-ExchangeRequest)
    - [ExchangeResponse](#gomoneypb-currency-v1-ExchangeResponse)
  
    - [CurrencyService](#gomoneypb-currency-v1-CurrencyService)
  
- [gomoneypb/v1/transaction.proto](#gomoneypb_v1_transaction-proto)
    - [Transaction](#gomoneypb-v1-Transaction)
    - [Transaction.ExtraEntry](#gomoneypb-v1-Transaction-ExtraEntry)
  
    - [TransactionType](#gomoneypb-v1-TransactionType)
  
- [gomoneypb/transactions/v1/transactions.proto](#gomoneypb_transactions_v1_transactions-proto)
    - [CreateTransactionRequest](#gomoneypb-transactions-v1-CreateTransactionRequest)
    - [CreateTransactionRequest.ExtraEntry](#gomoneypb-transactions-v1-CreateTransactionRequest-ExtraEntry)
    - [CreateTransactionResponse](#gomoneypb-transactions-v1-CreateTransactionResponse)
    - [Deposit](#gomoneypb-transactions-v1-Deposit)
    - [TransferBetweenAccounts](#gomoneypb-transactions-v1-TransferBetweenAccounts)
    - [Withdrawal](#gomoneypb-transactions-v1-Withdrawal)
  
    - [TransactionsService](#gomoneypb-transactions-v1-TransactionsService)
  
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
| type | [AccountType](#gomoneypb-v1-AccountType) |  |  |
| note | [string](#string) |  |  |






<a name="gomoneypb-v1-Account-ExtraEntry"></a>

### Account.ExtraEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="gomoneypb-v1-AccountType"></a>

### AccountType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACCOUNT_TYPE_UNSPECIFIED | 0 |  |
| ACCOUNT_TYPE_REGULAR | 1 |  |
| ACCOUNT_TYPE_SAVINGS | 2 |  |
| ACCOUNT_TYPE_BROKERAGE | 3 |  |
| ACCOUNT_TYPE_LIABILITY | 4 |  |


 

 

 



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






<a name="gomoneypb-accounts-v1-UpdateAccountRequest"></a>

### UpdateAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| extra | [UpdateAccountRequest.ExtraEntry](#gomoneypb-accounts-v1-UpdateAccountRequest-ExtraEntry) | repeated |  |
| type | [gomoneypb.v1.AccountType](#gomoneypb-v1-AccountType) |  |  |
| note | [string](#string) |  |  |






<a name="gomoneypb-accounts-v1-UpdateAccountRequest-ExtraEntry"></a>

### UpdateAccountRequest.ExtraEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="gomoneypb-accounts-v1-UpdateAccountResponse"></a>

### UpdateAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [gomoneypb.v1.Account](#gomoneypb-v1-Account) |  |  |





 

 

 


<a name="gomoneypb-accounts-v1-AccountsService"></a>

### AccountsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAccount | [CreateAccountRequest](#gomoneypb-accounts-v1-CreateAccountRequest) | [CreateAccountResponse](#gomoneypb-accounts-v1-CreateAccountResponse) |  |
| UpdateAccount | [UpdateAccountRequest](#gomoneypb-accounts-v1-UpdateAccountRequest) | [UpdateAccountResponse](#gomoneypb-accounts-v1-UpdateAccountResponse) |  |
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

 



<a name="gomoneypb_v1_transaction-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/v1/transaction.proto



<a name="gomoneypb-v1-Transaction"></a>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| source_amount | [string](#string) | optional |  |
| source_currency | [string](#string) | optional |  |
| destination_amount | [string](#string) | optional |  |
| destination_currency | [string](#string) | optional |  |
| source_account_id | [int32](#int32) | optional |  |
| destination_account_id | [int32](#int32) | optional |  |
| label_ids | [int32](#int32) | repeated |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| transaction_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| type | [TransactionType](#gomoneypb-v1-TransactionType) |  |  |
| notes | [string](#string) |  |  |
| extra | [Transaction.ExtraEntry](#gomoneypb-v1-Transaction-ExtraEntry) | repeated |  |






<a name="gomoneypb-v1-Transaction-ExtraEntry"></a>

### Transaction.ExtraEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="gomoneypb-v1-TransactionType"></a>

### TransactionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_TYPE_UNSPECIFIED | 0 |  |
| TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS | 1 |  |
| TRANSACTION_TYPE_DEPOSIT | 2 |  |
| TRANSACTION_TYPE_WITHDRAWAL | 3 |  |


 

 

 



<a name="gomoneypb_transactions_v1_transactions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gomoneypb/transactions/v1/transactions.proto



<a name="gomoneypb-transactions-v1-CreateTransactionRequest"></a>

### CreateTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| notes | [string](#string) |  |  |
| extra | [CreateTransactionRequest.ExtraEntry](#gomoneypb-transactions-v1-CreateTransactionRequest-ExtraEntry) | repeated |  |
| label_ids | [int32](#int32) | repeated |  |
| transaction_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| transfer_between_accounts | [TransferBetweenAccounts](#gomoneypb-transactions-v1-TransferBetweenAccounts) |  |  |
| deposit | [Deposit](#gomoneypb-transactions-v1-Deposit) |  |  |
| withdrawal | [Withdrawal](#gomoneypb-transactions-v1-Withdrawal) |  |  |






<a name="gomoneypb-transactions-v1-CreateTransactionRequest-ExtraEntry"></a>

### CreateTransactionRequest.ExtraEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="gomoneypb-transactions-v1-CreateTransactionResponse"></a>

### CreateTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction | [gomoneypb.v1.Transaction](#gomoneypb-v1-Transaction) |  |  |






<a name="gomoneypb-transactions-v1-Deposit"></a>

### Deposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| destination_amount | [string](#string) |  |  |
| destination_currency | [string](#string) |  |  |
| destination_account_id | [int32](#int32) |  |  |






<a name="gomoneypb-transactions-v1-TransferBetweenAccounts"></a>

### TransferBetweenAccounts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_account_id | [int32](#int32) |  |  |
| destination_account_id | [int32](#int32) |  |  |
| source_amount | [string](#string) |  |  |
| destination_amount | [string](#string) |  |  |
| source_currency | [string](#string) |  |  |
| destination_currency | [string](#string) |  |  |






<a name="gomoneypb-transactions-v1-Withdrawal"></a>

### Withdrawal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_amount | [string](#string) |  |  |
| source_currency | [string](#string) |  |  |
| source_account_id | [int32](#int32) |  |  |





 

 

 


<a name="gomoneypb-transactions-v1-TransactionsService"></a>

### TransactionsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateTransaction | [CreateTransactionRequest](#gomoneypb-transactions-v1-CreateTransactionRequest) | [CreateTransactionResponse](#gomoneypb-transactions-v1-CreateTransactionResponse) |  |

 



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

