// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file gomoneypb/accounts/v1/accounts.proto (package gomoneypb.accounts.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAccountRequest, CreateAccountResponse, DeleteAccountRequest, DeleteAccountResponse, ListAccountsRequest, ListAccountsResponse, ReorderAccountsRequest, ReorderAccountsResponse, UpdateAccountRequest, UpdateAccountResponse } from "./accounts_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gomoneypb.accounts.v1.AccountsService
 */
export const AccountsService = {
  typeName: "gomoneypb.accounts.v1.AccountsService",
  methods: {
    /**
     * @generated from rpc gomoneypb.accounts.v1.AccountsService.CreateAccount
     */
    createAccount: {
      name: "CreateAccount",
      I: CreateAccountRequest,
      O: CreateAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gomoneypb.accounts.v1.AccountsService.UpdateAccount
     */
    updateAccount: {
      name: "UpdateAccount",
      I: UpdateAccountRequest,
      O: UpdateAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gomoneypb.accounts.v1.AccountsService.DeleteAccount
     */
    deleteAccount: {
      name: "DeleteAccount",
      I: DeleteAccountRequest,
      O: DeleteAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gomoneypb.accounts.v1.AccountsService.ListAccounts
     */
    listAccounts: {
      name: "ListAccounts",
      I: ListAccountsRequest,
      O: ListAccountsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gomoneypb.accounts.v1.AccountsService.ReorderAccounts
     */
    reorderAccounts: {
      name: "ReorderAccounts",
      I: ReorderAccountsRequest,
      O: ReorderAccountsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

