// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

// Request wrapper for the ListAvailabilityDomains operation
type ListAvailabilityDomainsRequest struct {
        
 // The OCID of the compartment (remember that the tenancy is simply the root compartment). 
        CompartmentID string
}

// Response wrapper for the ListAvailabilityDomains operation
type ListAvailabilityDomainsResponse struct {
        
 // Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
 // particular request, please provide the request ID.
        OpcRequestID string
        
 // For pagination of a list of items. When paging through a list, if this header appears in the response,
 // then a partial list might have been returned. Include this value as the `page` parameter for the
 // subsequent GET request to get the next batch of items.
        OpcNextPage string

        
 // The []AvailabilityDomain instance
        ListAvailabilityDomains []AvailabilityDomain


}

