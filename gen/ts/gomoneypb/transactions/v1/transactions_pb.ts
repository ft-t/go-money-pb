// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file gomoneypb/transactions/v1/transactions.proto (package gomoneypb.transactions.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Transaction } from "../../v1/transaction_pb";

/**
 * @generated from message gomoneypb.transactions.v1.CreateTransactionRequest
 */
export class CreateTransactionRequest extends Message<CreateTransactionRequest> {
  /**
   * @generated from field: string notes = 1;
   */
  notes = "";

  /**
   * @generated from field: map<string, string> extra = 2;
   */
  extra: { [key: string]: string } = {};

  /**
   * @generated from field: repeated int32 label_ids = 3;
   */
  labelIds: number[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp transaction_date = 7;
   */
  transactionDate?: Timestamp;

  /**
   * @generated from oneof gomoneypb.transactions.v1.CreateTransactionRequest.transaction
   */
  transaction: {
    /**
     * @generated from field: gomoneypb.transactions.v1.TransferBetweenAccounts transfer_between_accounts = 4;
     */
    value: TransferBetweenAccounts;
    case: "transferBetweenAccounts";
  } | {
    /**
     * @generated from field: gomoneypb.transactions.v1.Deposit deposit = 5;
     */
    value: Deposit;
    case: "deposit";
  } | {
    /**
     * @generated from field: gomoneypb.transactions.v1.Withdrawal withdrawal = 6;
     */
    value: Withdrawal;
    case: "withdrawal";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<CreateTransactionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.transactions.v1.CreateTransactionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "extra", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 3, name: "label_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 7, name: "transaction_date", kind: "message", T: Timestamp },
    { no: 4, name: "transfer_between_accounts", kind: "message", T: TransferBetweenAccounts, oneof: "transaction" },
    { no: 5, name: "deposit", kind: "message", T: Deposit, oneof: "transaction" },
    { no: 6, name: "withdrawal", kind: "message", T: Withdrawal, oneof: "transaction" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTransactionRequest {
    return new CreateTransactionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTransactionRequest {
    return new CreateTransactionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTransactionRequest {
    return new CreateTransactionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTransactionRequest | PlainMessage<CreateTransactionRequest> | undefined, b: CreateTransactionRequest | PlainMessage<CreateTransactionRequest> | undefined): boolean {
    return proto3.util.equals(CreateTransactionRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.transactions.v1.Withdrawal
 */
export class Withdrawal extends Message<Withdrawal> {
  /**
   * @generated from field: string source_amount = 1;
   */
  sourceAmount = "";

  /**
   * @generated from field: string source_currency = 2;
   */
  sourceCurrency = "";

  /**
   * @generated from field: int32 source_account_id = 3;
   */
  sourceAccountId = 0;

  constructor(data?: PartialMessage<Withdrawal>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.transactions.v1.Withdrawal";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source_amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "source_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "source_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Withdrawal {
    return new Withdrawal().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Withdrawal {
    return new Withdrawal().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Withdrawal {
    return new Withdrawal().fromJsonString(jsonString, options);
  }

  static equals(a: Withdrawal | PlainMessage<Withdrawal> | undefined, b: Withdrawal | PlainMessage<Withdrawal> | undefined): boolean {
    return proto3.util.equals(Withdrawal, a, b);
  }
}

/**
 * @generated from message gomoneypb.transactions.v1.TransferBetweenAccounts
 */
export class TransferBetweenAccounts extends Message<TransferBetweenAccounts> {
  /**
   * @generated from field: int32 source_account_id = 1;
   */
  sourceAccountId = 0;

  /**
   * @generated from field: int32 destination_account_id = 2;
   */
  destinationAccountId = 0;

  /**
   * @generated from field: string source_amount = 3;
   */
  sourceAmount = "";

  /**
   * @generated from field: string destination_amount = 4;
   */
  destinationAmount = "";

  /**
   * @generated from field: string source_currency = 5;
   */
  sourceCurrency = "";

  /**
   * @generated from field: string destination_currency = 6;
   */
  destinationCurrency = "";

  constructor(data?: PartialMessage<TransferBetweenAccounts>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.transactions.v1.TransferBetweenAccounts";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "destination_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "source_amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "destination_amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "source_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "destination_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TransferBetweenAccounts {
    return new TransferBetweenAccounts().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TransferBetweenAccounts {
    return new TransferBetweenAccounts().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TransferBetweenAccounts {
    return new TransferBetweenAccounts().fromJsonString(jsonString, options);
  }

  static equals(a: TransferBetweenAccounts | PlainMessage<TransferBetweenAccounts> | undefined, b: TransferBetweenAccounts | PlainMessage<TransferBetweenAccounts> | undefined): boolean {
    return proto3.util.equals(TransferBetweenAccounts, a, b);
  }
}

/**
 * @generated from message gomoneypb.transactions.v1.Deposit
 */
export class Deposit extends Message<Deposit> {
  /**
   * @generated from field: string destination_amount = 1;
   */
  destinationAmount = "";

  /**
   * @generated from field: string destination_currency = 2;
   */
  destinationCurrency = "";

  /**
   * @generated from field: int32 destination_account_id = 3;
   */
  destinationAccountId = 0;

  constructor(data?: PartialMessage<Deposit>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.transactions.v1.Deposit";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "destination_amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "destination_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "destination_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Deposit {
    return new Deposit().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Deposit {
    return new Deposit().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Deposit {
    return new Deposit().fromJsonString(jsonString, options);
  }

  static equals(a: Deposit | PlainMessage<Deposit> | undefined, b: Deposit | PlainMessage<Deposit> | undefined): boolean {
    return proto3.util.equals(Deposit, a, b);
  }
}

/**
 * @generated from message gomoneypb.transactions.v1.CreateTransactionResponse
 */
export class CreateTransactionResponse extends Message<CreateTransactionResponse> {
  /**
   * @generated from field: gomoneypb.v1.Transaction transaction = 1;
   */
  transaction?: Transaction;

  constructor(data?: PartialMessage<CreateTransactionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.transactions.v1.CreateTransactionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "transaction", kind: "message", T: Transaction },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTransactionResponse {
    return new CreateTransactionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTransactionResponse {
    return new CreateTransactionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTransactionResponse {
    return new CreateTransactionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTransactionResponse | PlainMessage<CreateTransactionResponse> | undefined, b: CreateTransactionResponse | PlainMessage<CreateTransactionResponse> | undefined): boolean {
    return proto3.util.equals(CreateTransactionResponse, a, b);
  }
}

