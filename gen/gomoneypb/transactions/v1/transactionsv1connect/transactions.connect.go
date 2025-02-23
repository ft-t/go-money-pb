// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gomoneypb/transactions/v1/transactions.proto

package transactionsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ft-t/go-money-pb/gen/gomoneypb/transactions/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// TransactionsServiceName is the fully-qualified name of the TransactionsService service.
	TransactionsServiceName = "gomoneypb.transactions.v1.TransactionsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TransactionsServiceCreateTransactionProcedure is the fully-qualified name of the
	// TransactionsService's CreateTransaction RPC.
	TransactionsServiceCreateTransactionProcedure = "/gomoneypb.transactions.v1.TransactionsService/CreateTransaction"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	transactionsServiceServiceDescriptor                 = v1.File_gomoneypb_transactions_v1_transactions_proto.Services().ByName("TransactionsService")
	transactionsServiceCreateTransactionMethodDescriptor = transactionsServiceServiceDescriptor.Methods().ByName("CreateTransaction")
)

// TransactionsServiceClient is a client for the gomoneypb.transactions.v1.TransactionsService
// service.
type TransactionsServiceClient interface {
	CreateTransaction(context.Context, *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error)
}

// NewTransactionsServiceClient constructs a client for the
// gomoneypb.transactions.v1.TransactionsService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransactionsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TransactionsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transactionsServiceClient{
		createTransaction: connect.NewClient[v1.CreateTransactionRequest, v1.CreateTransactionResponse](
			httpClient,
			baseURL+TransactionsServiceCreateTransactionProcedure,
			connect.WithSchema(transactionsServiceCreateTransactionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// transactionsServiceClient implements TransactionsServiceClient.
type transactionsServiceClient struct {
	createTransaction *connect.Client[v1.CreateTransactionRequest, v1.CreateTransactionResponse]
}

// CreateTransaction calls gomoneypb.transactions.v1.TransactionsService.CreateTransaction.
func (c *transactionsServiceClient) CreateTransaction(ctx context.Context, req *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error) {
	return c.createTransaction.CallUnary(ctx, req)
}

// TransactionsServiceHandler is an implementation of the
// gomoneypb.transactions.v1.TransactionsService service.
type TransactionsServiceHandler interface {
	CreateTransaction(context.Context, *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error)
}

// NewTransactionsServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransactionsServiceHandler(svc TransactionsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	transactionsServiceCreateTransactionHandler := connect.NewUnaryHandler(
		TransactionsServiceCreateTransactionProcedure,
		svc.CreateTransaction,
		connect.WithSchema(transactionsServiceCreateTransactionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/gomoneypb.transactions.v1.TransactionsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TransactionsServiceCreateTransactionProcedure:
			transactionsServiceCreateTransactionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTransactionsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTransactionsServiceHandler struct{}

func (UnimplementedTransactionsServiceHandler) CreateTransaction(context.Context, *connect.Request[v1.CreateTransactionRequest]) (*connect.Response[v1.CreateTransactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("gomoneypb.transactions.v1.TransactionsService.CreateTransaction is not implemented"))
}
