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
	"github.com/oracle/oci-go-sdk/v42/common"
)

// InstanceReservationConfig Data that defines the instance reservation configuration.
type InstanceReservationConfig struct {

	// The shape to use when launching instances using compute capacity reservations. The shape determines the number of CPUs, the amount of memory,
	// and other resources allocated to the instance.
	// You can list all available shapes by calling ListComputeCapacityReservationInstanceShapes.
	InstanceShape *string `mandatory:"true" json:"instanceShape"`

	// The amount of capacity reserved in this configuration.
	ReservedCount *int64 `mandatory:"true" json:"reservedCount"`

	// The amount of capacity in use out of the total capacity reserved in this reservation configuration.
	UsedCount *int64 `mandatory:"true" json:"usedCount"`

	// The fault domain of this reservation configuration.
	// If a value is not supplied, this reservation configuration is applicable to all fault domains in the specified availability domain.
	// For more information, see Capacity Reservations (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm).
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	InstanceShapeConfig *InstanceReservationShapeConfigDetails `mandatory:"false" json:"instanceShapeConfig"`
}

func (m InstanceReservationConfig) String() string {
	return common.PointerString(m)
}
