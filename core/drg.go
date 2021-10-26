// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Drg A dynamic routing gateway (DRG) is a virtual router that provides a path for private
// network traffic between networks. You use it with other Networking
// Service components to create a connection to your on-premises network using VPN Connect (https://docs.cloud.oracle.com/Content/Network/Tasks/managingIPsec.htm) or a connection that uses
// FastConnect (https://docs.cloud.oracle.com/Content/Network/Concepts/fastconnect.htm). For more information, see
// Overview of the Networking Service (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type Drg struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the DRG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The DRG's Oracle ID (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The DRG's current state.
	LifecycleState DrgLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The date and time the DRG was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	DefaultDrgRouteTables *DefaultDrgRouteTables `mandatory:"false" json:"defaultDrgRouteTables"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of this DRG's default export route distribution for the DRG attachments.
	DefaultExportDrgRouteDistributionId *string `mandatory:"false" json:"defaultExportDrgRouteDistributionId"`
}

func (m Drg) String() string {
	return common.PointerString(m)
}

// DrgLifecycleStateEnum Enum with underlying type: string
type DrgLifecycleStateEnum string

// Set of constants representing the allowable values for DrgLifecycleStateEnum
const (
	DrgLifecycleStateProvisioning DrgLifecycleStateEnum = "PROVISIONING"
	DrgLifecycleStateAvailable    DrgLifecycleStateEnum = "AVAILABLE"
	DrgLifecycleStateTerminating  DrgLifecycleStateEnum = "TERMINATING"
	DrgLifecycleStateTerminated   DrgLifecycleStateEnum = "TERMINATED"
)

var mappingDrgLifecycleState = map[string]DrgLifecycleStateEnum{
	"PROVISIONING": DrgLifecycleStateProvisioning,
	"AVAILABLE":    DrgLifecycleStateAvailable,
	"TERMINATING":  DrgLifecycleStateTerminating,
	"TERMINATED":   DrgLifecycleStateTerminated,
}

// GetDrgLifecycleStateEnumValues Enumerates the set of values for DrgLifecycleStateEnum
func GetDrgLifecycleStateEnumValues() []DrgLifecycleStateEnum {
	values := make([]DrgLifecycleStateEnum, 0)
	for _, v := range mappingDrgLifecycleState {
		values = append(values, v)
	}
	return values
}
