// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package keymanagement

import (
	"github.com/oracle/oci-go-sdk/v46/common"
	"net/http"
)

// ListVaultReplicasRequest wrapper for the ListVaultReplicas operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/keymanagement/ListVaultReplicas.go.html to see an example of how to use ListVaultReplicasRequest.
type ListVaultReplicasRequest struct {

	// The OCID of the vault.
	VaultId *string `mandatory:"true" contributesTo:"path" name:"vaultId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a
	// resource, set the `if-match` parameter to the value of the etag from a
	// previous GET or POST response for that resource. The resource will be
	// updated or deleted only if the etag you provide matches the resource's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated
	// before then due to conflicting operations (e.g., if a resource has been
	// deleted and purged from the system, then a retry of the original
	// creation request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `TIMECREATED` is descending. The default order for `DISPLAYNAME`
	// is ascending.
	SortBy ListVaultReplicasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListVaultReplicasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVaultReplicasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVaultReplicasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVaultReplicasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVaultReplicasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVaultReplicasResponse wrapper for the ListVaultReplicas operation
type ListVaultReplicasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VaultReplicaSummary instances
	Items []VaultReplicaSummary `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVaultReplicasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVaultReplicasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVaultReplicasSortByEnum Enum with underlying type: string
type ListVaultReplicasSortByEnum string

// Set of constants representing the allowable values for ListVaultReplicasSortByEnum
const (
	ListVaultReplicasSortByTimecreated ListVaultReplicasSortByEnum = "TIMECREATED"
	ListVaultReplicasSortByDisplayname ListVaultReplicasSortByEnum = "DISPLAYNAME"
)

var mappingListVaultReplicasSortBy = map[string]ListVaultReplicasSortByEnum{
	"TIMECREATED": ListVaultReplicasSortByTimecreated,
	"DISPLAYNAME": ListVaultReplicasSortByDisplayname,
}

// GetListVaultReplicasSortByEnumValues Enumerates the set of values for ListVaultReplicasSortByEnum
func GetListVaultReplicasSortByEnumValues() []ListVaultReplicasSortByEnum {
	values := make([]ListVaultReplicasSortByEnum, 0)
	for _, v := range mappingListVaultReplicasSortBy {
		values = append(values, v)
	}
	return values
}

// ListVaultReplicasSortOrderEnum Enum with underlying type: string
type ListVaultReplicasSortOrderEnum string

// Set of constants representing the allowable values for ListVaultReplicasSortOrderEnum
const (
	ListVaultReplicasSortOrderAsc  ListVaultReplicasSortOrderEnum = "ASC"
	ListVaultReplicasSortOrderDesc ListVaultReplicasSortOrderEnum = "DESC"
)

var mappingListVaultReplicasSortOrder = map[string]ListVaultReplicasSortOrderEnum{
	"ASC":  ListVaultReplicasSortOrderAsc,
	"DESC": ListVaultReplicasSortOrderDesc,
}

// GetListVaultReplicasSortOrderEnumValues Enumerates the set of values for ListVaultReplicasSortOrderEnum
func GetListVaultReplicasSortOrderEnumValues() []ListVaultReplicasSortOrderEnum {
	values := make([]ListVaultReplicasSortOrderEnum, 0)
	for _, v := range mappingListVaultReplicasSortOrder {
		values = append(values, v)
	}
	return values
}
