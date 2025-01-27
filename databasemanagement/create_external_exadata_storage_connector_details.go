// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalExadataStorageConnectorDetails The details of creating the connector to the Exadata storage server.
type CreateExternalExadataStorageConnectorDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata storage server.
	StorageServerId *string `mandatory:"true" json:"storageServerId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The connector name if OCI connector is created.
	ConnectorName *string `mandatory:"true" json:"connectorName"`

	// The unique connection string of the connection. For example, "https://slcm21celadm02.us.oracle.com:443/MS/RESTService/".
	ConnectionUri *string `mandatory:"true" json:"connectionUri"`

	CredentialInfo *RestCredential `mandatory:"true" json:"credentialInfo"`
}

func (m CreateExternalExadataStorageConnectorDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalExadataStorageConnectorDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
