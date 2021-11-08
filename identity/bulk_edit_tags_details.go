// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// BulkEditTagsDetails The representation of BulkEditTagsDetails
type BulkEditTagsDetails struct {

	// The OCID of the compartment where the bulk tag edit request is submitted.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources to be updated.
	Resources []BulkEditResource `mandatory:"true" json:"resources"`

	// The operations associated with the request to bulk edit tags.
	BulkEditOperations []BulkEditOperationDetails `mandatory:"true" json:"bulkEditOperations"`
}

func (m BulkEditTagsDetails) String() string {
	return common.PointerString(m)
}
