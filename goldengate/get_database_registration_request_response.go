// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetDatabaseRegistrationRequest wrapper for the GetDatabaseRegistration operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/GetDatabaseRegistration.go.html to see an example of how to use GetDatabaseRegistrationRequest.
type GetDatabaseRegistrationRequest struct {

	// A unique DatabaseRegistration identifier.
	DatabaseRegistrationId *string `mandatory:"true" contributesTo:"path" name:"databaseRegistrationId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDatabaseRegistrationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseRegistrationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseRegistrationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseRegistrationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDatabaseRegistrationResponse wrapper for the GetDatabaseRegistration operation
type GetDatabaseRegistrationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseRegistration instance
	DatabaseRegistration `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDatabaseRegistrationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseRegistrationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
