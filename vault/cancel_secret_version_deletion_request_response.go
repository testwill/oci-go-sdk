// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package vault

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"net/http"
	"strings"
)

// CancelSecretVersionDeletionRequest wrapper for the CancelSecretVersionDeletion operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/vault/CancelSecretVersionDeletion.go.html to see an example of how to use CancelSecretVersionDeletionRequest.
type CancelSecretVersionDeletionRequest struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// The version number of the secret.
	SecretVersionNumber *int64 `mandatory:"true" contributesTo:"path" name:"secretVersionNumber"`

	// For optimistic concurrency control. In the PUT or DELETE call for a
	// resource, set the `if-match` parameter to the value of the etag from a
	// previous GET or POST response for that resource. The resource will be
	// updated or deleted only if the etag you provide matches the resource's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CancelSecretVersionDeletionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CancelSecretVersionDeletionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CancelSecretVersionDeletionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CancelSecretVersionDeletionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CancelSecretVersionDeletionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CancelSecretVersionDeletionResponse wrapper for the CancelSecretVersionDeletion operation
type CancelSecretVersionDeletionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CancelSecretVersionDeletionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CancelSecretVersionDeletionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
