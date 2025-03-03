// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file gomoneypb/currency/v1/currency.proto (package gomoneypb.currency.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Currency } from "../../v1/currency_pb";

/**
 * @generated from message gomoneypb.currency.v1.DeleteCurrencyRequest
 */
export class DeleteCurrencyRequest extends Message<DeleteCurrencyRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteCurrencyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.DeleteCurrencyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCurrencyRequest {
    return new DeleteCurrencyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCurrencyRequest {
    return new DeleteCurrencyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCurrencyRequest {
    return new DeleteCurrencyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteCurrencyRequest | PlainMessage<DeleteCurrencyRequest> | undefined, b: DeleteCurrencyRequest | PlainMessage<DeleteCurrencyRequest> | undefined): boolean {
    return proto3.util.equals(DeleteCurrencyRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.DeleteCurrencyResponse
 */
export class DeleteCurrencyResponse extends Message<DeleteCurrencyResponse> {
  /**
   * @generated from field: gomoneypb.v1.Currency currency = 1;
   */
  currency?: Currency;

  constructor(data?: PartialMessage<DeleteCurrencyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.DeleteCurrencyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "currency", kind: "message", T: Currency },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCurrencyResponse {
    return new DeleteCurrencyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCurrencyResponse {
    return new DeleteCurrencyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCurrencyResponse {
    return new DeleteCurrencyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteCurrencyResponse | PlainMessage<DeleteCurrencyResponse> | undefined, b: DeleteCurrencyResponse | PlainMessage<DeleteCurrencyResponse> | undefined): boolean {
    return proto3.util.equals(DeleteCurrencyResponse, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.UpdateCurrencyRequest
 */
export class UpdateCurrencyRequest extends Message<UpdateCurrencyRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string rate = 2;
   */
  rate = "";

  /**
   * @generated from field: bool is_active = 3;
   */
  isActive = false;

  /**
   * @generated from field: int32 decimal_places = 4;
   */
  decimalPlaces = 0;

  constructor(data?: PartialMessage<UpdateCurrencyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.UpdateCurrencyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "rate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "decimal_places", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCurrencyRequest {
    return new UpdateCurrencyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCurrencyRequest {
    return new UpdateCurrencyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCurrencyRequest {
    return new UpdateCurrencyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateCurrencyRequest | PlainMessage<UpdateCurrencyRequest> | undefined, b: UpdateCurrencyRequest | PlainMessage<UpdateCurrencyRequest> | undefined): boolean {
    return proto3.util.equals(UpdateCurrencyRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.UpdateCurrencyResponse
 */
export class UpdateCurrencyResponse extends Message<UpdateCurrencyResponse> {
  /**
   * @generated from field: gomoneypb.v1.Currency currency = 1;
   */
  currency?: Currency;

  constructor(data?: PartialMessage<UpdateCurrencyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.UpdateCurrencyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "currency", kind: "message", T: Currency },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCurrencyResponse {
    return new UpdateCurrencyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCurrencyResponse {
    return new UpdateCurrencyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCurrencyResponse {
    return new UpdateCurrencyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateCurrencyResponse | PlainMessage<UpdateCurrencyResponse> | undefined, b: UpdateCurrencyResponse | PlainMessage<UpdateCurrencyResponse> | undefined): boolean {
    return proto3.util.equals(UpdateCurrencyResponse, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.CreateCurrencyResponse
 */
export class CreateCurrencyResponse extends Message<CreateCurrencyResponse> {
  /**
   * @generated from field: gomoneypb.v1.Currency currency = 1;
   */
  currency?: Currency;

  constructor(data?: PartialMessage<CreateCurrencyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.CreateCurrencyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "currency", kind: "message", T: Currency },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCurrencyResponse {
    return new CreateCurrencyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCurrencyResponse {
    return new CreateCurrencyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCurrencyResponse {
    return new CreateCurrencyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateCurrencyResponse | PlainMessage<CreateCurrencyResponse> | undefined, b: CreateCurrencyResponse | PlainMessage<CreateCurrencyResponse> | undefined): boolean {
    return proto3.util.equals(CreateCurrencyResponse, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.CreateCurrencyRequest
 */
export class CreateCurrencyRequest extends Message<CreateCurrencyRequest> {
  /**
   * @generated from field: gomoneypb.v1.Currency currency = 1;
   */
  currency?: Currency;

  constructor(data?: PartialMessage<CreateCurrencyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.CreateCurrencyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "currency", kind: "message", T: Currency },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCurrencyRequest {
    return new CreateCurrencyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCurrencyRequest {
    return new CreateCurrencyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCurrencyRequest {
    return new CreateCurrencyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateCurrencyRequest | PlainMessage<CreateCurrencyRequest> | undefined, b: CreateCurrencyRequest | PlainMessage<CreateCurrencyRequest> | undefined): boolean {
    return proto3.util.equals(CreateCurrencyRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.GetCurrenciesRequest
 */
export class GetCurrenciesRequest extends Message<GetCurrenciesRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string rate = 2;
   */
  rate = "";

  /**
   * @generated from field: bool is_active = 3;
   */
  isActive = false;

  /**
   * @generated from field: int32 decimal_places = 4;
   */
  decimalPlaces = 0;

  constructor(data?: PartialMessage<GetCurrenciesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.GetCurrenciesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "rate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "decimal_places", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCurrenciesRequest {
    return new GetCurrenciesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCurrenciesRequest {
    return new GetCurrenciesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCurrenciesRequest {
    return new GetCurrenciesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetCurrenciesRequest | PlainMessage<GetCurrenciesRequest> | undefined, b: GetCurrenciesRequest | PlainMessage<GetCurrenciesRequest> | undefined): boolean {
    return proto3.util.equals(GetCurrenciesRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.GetCurrenciesResponse
 */
export class GetCurrenciesResponse extends Message<GetCurrenciesResponse> {
  /**
   * @generated from field: repeated gomoneypb.v1.Currency currencies = 1;
   */
  currencies: Currency[] = [];

  constructor(data?: PartialMessage<GetCurrenciesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.GetCurrenciesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "currencies", kind: "message", T: Currency, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCurrenciesResponse {
    return new GetCurrenciesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCurrenciesResponse {
    return new GetCurrenciesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCurrenciesResponse {
    return new GetCurrenciesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetCurrenciesResponse | PlainMessage<GetCurrenciesResponse> | undefined, b: GetCurrenciesResponse | PlainMessage<GetCurrenciesResponse> | undefined): boolean {
    return proto3.util.equals(GetCurrenciesResponse, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.ExchangeRequest
 */
export class ExchangeRequest extends Message<ExchangeRequest> {
  /**
   * @generated from field: string from_currency = 1;
   */
  fromCurrency = "";

  /**
   * @generated from field: string to_currency = 2;
   */
  toCurrency = "";

  /**
   * @generated from field: string amount = 3;
   */
  amount = "";

  constructor(data?: PartialMessage<ExchangeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.ExchangeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "from_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "to_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExchangeRequest {
    return new ExchangeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExchangeRequest {
    return new ExchangeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExchangeRequest {
    return new ExchangeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ExchangeRequest | PlainMessage<ExchangeRequest> | undefined, b: ExchangeRequest | PlainMessage<ExchangeRequest> | undefined): boolean {
    return proto3.util.equals(ExchangeRequest, a, b);
  }
}

/**
 * @generated from message gomoneypb.currency.v1.ExchangeResponse
 */
export class ExchangeResponse extends Message<ExchangeResponse> {
  /**
   * @generated from field: string amount = 1;
   */
  amount = "";

  constructor(data?: PartialMessage<ExchangeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "gomoneypb.currency.v1.ExchangeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "amount", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExchangeResponse {
    return new ExchangeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExchangeResponse {
    return new ExchangeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExchangeResponse {
    return new ExchangeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ExchangeResponse | PlainMessage<ExchangeResponse> | undefined, b: ExchangeResponse | PlainMessage<ExchangeResponse> | undefined): boolean {
    return proto3.util.equals(ExchangeResponse, a, b);
  }
}

