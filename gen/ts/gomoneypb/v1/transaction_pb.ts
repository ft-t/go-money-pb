// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file gomoneypb/v1/transaction.proto (package gomoneypb.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum gomoneypb.v1.TransactionType
 */
export enum TransactionType {
  /**
   * @generated from enum value: TRANSACTION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS = 1;
   */
  TRANSFER_BETWEEN_ACCOUNTS = 1,

  /**
   * @generated from enum value: TRANSACTION_TYPE_DEPOSIT = 2;
   */
  DEPOSIT = 2,

  /**
   * @generated from enum value: TRANSACTION_TYPE_WITHDRAWAL = 3;
   */
  WITHDRAWAL = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(TransactionType)
proto3.util.setEnumType(TransactionType, "gomoneypb.v1.TransactionType", [
  { no: 0, name: "TRANSACTION_TYPE_UNSPECIFIED" },
  { no: 1, name: "TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS" },
  { no: 2, name: "TRANSACTION_TYPE_DEPOSIT" },
  { no: 3, name: "TRANSACTION_TYPE_WITHDRAWAL" },
]);

/**
 * @generated from message gomoneypb.v1.Transaction
 */
export class Transaction extends Message<Transaction> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: optional string source_amount = 2;
   */
  sourceAmount?: string;

  /**
   * @generated from field: optional string source_currency = 3;
   */
  sourceCurrency?: string;

  /**
   * @generated from field: optional string destination_amount = 4;
   */
  destinationAmount?: string;

  /**
   * @generated from field: optional string destination_currency = 5;
   */
  destinationCurrency?: string;

  /**
   * @generated from field: optional int32 source_account_id = 6;
   */
  sourceAccountId?: number;

  /**
   * @generated from field: optional int32 destination_account_id = 7;
   */
  destinationAccountId?: number;

  /**
   * @generated from field: repeated int32 label_ids = 8;
   */
  labelIds: number[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 9;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 10;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp transaction_date = 11;
   */
  transactionDate?: Timestamp;

  /**
   * @generated from field: gomoneypb.v1.TransactionType type = 12;
   */
  type = TransactionType.UNSPECIFIED;

  /**
   * @generated from field: string notes = 13;
   */
  notes = "";

  /**
   * @generated from field: map<string, string> extra = 14;
   */
  extra: { [key: string]: string } = {};

  constructor(data?: PartialMessage<Transaction>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.v1.Transaction";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "source_amount", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "source_currency", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "destination_amount", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "destination_currency", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "source_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 7, name: "destination_account_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 8, name: "label_ids", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 9, name: "created_at", kind: "message", T: Timestamp },
    { no: 10, name: "updated_at", kind: "message", T: Timestamp },
    { no: 11, name: "transaction_date", kind: "message", T: Timestamp },
    { no: 12, name: "type", kind: "enum", T: proto3.getEnumType(TransactionType) },
    { no: 13, name: "notes", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 14, name: "extra", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Transaction {
    return new Transaction().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Transaction {
    return new Transaction().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Transaction {
    return new Transaction().fromJsonString(jsonString, options);
  }

  static equals(a: Transaction | PlainMessage<Transaction> | undefined, b: Transaction | PlainMessage<Transaction> | undefined): boolean {
    return proto3.util.equals(Transaction, a, b);
  }
}

