// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetTargetResponderRecipeResponderRuleRequest wrapper for the GetTargetResponderRecipeResponderRule operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/GetTargetResponderRecipeResponderRule.go.html to see an example of how to use GetTargetResponderRecipeResponderRuleRequest.
type GetTargetResponderRecipeResponderRuleRequest struct {

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// OCID of TargetResponderRecipe
	TargetResponderRecipeId *string `mandatory:"true" contributesTo:"path" name:"targetResponderRecipeId"`

	// The id of ResponderRule
	ResponderRuleId *string `mandatory:"true" contributesTo:"path" name:"responderRuleId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTargetResponderRecipeResponderRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTargetResponderRecipeResponderRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetTargetResponderRecipeResponderRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTargetResponderRecipeResponderRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTargetResponderRecipeResponderRuleResponse wrapper for the GetTargetResponderRecipeResponderRule operation
type GetTargetResponderRecipeResponderRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TargetResponderRecipeResponderRule instance
	TargetResponderRecipeResponderRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTargetResponderRecipeResponderRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTargetResponderRecipeResponderRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
