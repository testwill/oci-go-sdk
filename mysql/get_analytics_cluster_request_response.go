// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package mysql

import (
	"github.com/oracle/oci-go-sdk/v44/common"
	"net/http"
)

// GetAnalyticsClusterRequest wrapper for the GetAnalyticsCluster operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/mysql/GetAnalyticsCluster.go.html to see an example of how to use GetAnalyticsClusterRequest.
type GetAnalyticsClusterRequest struct {

	// The DB System OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// Customer-defined unique identifier for the request. If you need to
	// contact Oracle about a specific request, please provide the request
	// ID that you supplied in this header with the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For conditional requests. In the GET call for a resource, set the
	// `If-None-Match` header to the value of the ETag from a previous GET (or
	// POST or PUT) response for that resource. The server will return with
	// either a 304 Not Modified response if the resource has not changed, or a
	// 200 OK response with the updated representation.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAnalyticsClusterRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAnalyticsClusterRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAnalyticsClusterRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAnalyticsClusterRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAnalyticsClusterResponse wrapper for the GetAnalyticsCluster operation
type GetAnalyticsClusterResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AnalyticsCluster instance
	AnalyticsCluster `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetAnalyticsClusterResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAnalyticsClusterResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
