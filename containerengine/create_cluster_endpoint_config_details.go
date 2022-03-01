// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// CreateClusterEndpointConfigDetails The properties that define the network configuration for the Cluster endpoint.
type CreateClusterEndpointConfigDetails struct {

	// The OCID of the regional subnet in which to place the Cluster endpoint.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A list of the OCIDs of the network security groups (NSGs) to apply to the cluster endpoint. For more information about NSGs, see NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Whether the cluster should be assigned a public IP address. Defaults to false. If set to true on a private subnet, the cluster provisioning will fail.
	IsPublicIpEnabled *bool `mandatory:"false" json:"isPublicIpEnabled"`
}

func (m CreateClusterEndpointConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateClusterEndpointConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
