package golsptoolkit

import "encoding/json"

// HeaderPart represents the parsed LSP message header.
//
// In LSP, each message is sent as ASCII header lines followed by a JSON body,
// separated by a blank line ("\r\n\r\n"). At minimum, a Content-Length header
// is required.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#baseProtocol
type HeaderPart struct {
	ContentLength int
	ContentType   string
}

type ContentPart struct {
}

// Base Types represents the base types used in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#baseTypes
type (
	Integer   = int32
	UInteger  = uint32
	Decimal   = float64
	LSPAny    = any
	LSPObject = map[string]LSPAny
	LSPArray  = []LSPAny
)

// IsLSPAny checks if the given value is a valid LSPAny type.
func IsLSPAny(v any) bool {
	switch v.(type) {
	case string, Integer, UInteger, Decimal, bool, nil, LSPObject, LSPArray:
		return true
	default:
		return false
	}
}

// IsLSPArrayOrObject checks if the given value is a valid LSPArray or LSPObject type.
func IsLSPArrayOrObject(v any) bool {
	switch v.(type) {
	case LSPArray, LSPObject:
		return true
	default:
		return false
	}
}

// Abstract Message represents a base message structure in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#abstractMessage
type AbstractMessage struct {
	JSONRPC string `json:"jsonrpc"`
}

// Request Message represents a request message structure in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
type RequestMessage struct {
	AbstractMessage
	ID     json.Number `json:"id"`
	Method string      `json:"method"`
	Params LSPAny      `json:"params,omitempty"`
}

// Response Message represents a response message structure in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
type ResponseMessage struct {
	AbstractMessage
	ID     *json.Number   `json:"id"`
	Result LSPAny         `json:"result,omitempty"`
	Error  *ResponseError `json:"error,omitempty"`
}

// Response Error represents an error response structure in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseError
type ResponseError struct {
	Code    Integer `json:"code"`
	Message string  `json:"message"`
	Data    LSPAny  `json:"data,omitempty"`
}

// Notification Message represents a notification message structure in the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notificationMessage
type NotificationMessage struct {
	AbstractMessage
	Method string `json:"method"`
	Params LSPAny `json:"params,omitempty"`
}

// ErrorCodes represents the error codes defined by the Language Server Protocol.
//
// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#errorCodes
const (
	// JSON-RPC defined error codes
	ParseError     Integer = -32700
	InvalidRequest Integer = -32600
	MethodNotFound Integer = -32601
	InvalidParams  Integer = -32602
	InternalError  Integer = -32603

	// JSON-RPC reserved error range
	JSONRPCReservedErrorRangeStart Integer = -32099
	// Deprecated: Use JSONRPCReservedErrorRangeStart instead.
	ServerErrorStart             Integer = -32099
	ServerNotInitialized         Integer = -32002
	UnknownErrorCode             Integer = -32001
	JSONRPCReservedErrorRangeEnd Integer = -32000
	// Deprecated: Use JSONRPCReservedErrorRangeEnd instead.
	ServerErrorEnd Integer = -32000

	// LSP reserved error range
	LSPReservedErrorRangeStart Integer = -32899
	RequestFailed              Integer = -32803
	ServerCancelled            Integer = -32804
	ContentModified            Integer = -32801
	RequestCancelled           Integer = -32800
	LSPReservedErrorRangeEnd   Integer = -32800
)

type CancelParams struct {
	ID json.Number `json:"id"`
}

type ProgressParams[T any] struct {
	Token json.Number `json:"token"`
	Value T           `json:"value"`
}
