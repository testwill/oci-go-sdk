// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// InstanceConfiguration The model deployment instance configuration
type InstanceConfiguration struct {

	// The shape used to launch the model deployment instances.
	InstanceShapeName *string `mandatory:"true" json:"instanceShapeName"`
}

func (m InstanceConfiguration) String() string {
	return common.PointerString(m)
}
