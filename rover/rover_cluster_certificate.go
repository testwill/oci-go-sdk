// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// RoverCloudService API
//
// A description of the RoverCloudService API.
//

package rover

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// RoverClusterCertificate The certificate response
type RoverClusterCertificate struct {

	// The certificate that can be installed on a client to do TLS communication to the cluster
	Certificate *string `mandatory:"true" json:"certificate"`
}

func (m RoverClusterCertificate) String() string {
	return common.PointerString(m)
}