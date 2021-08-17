// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// AccessRequestHistoryCollection Results of access request history search, which contains summary of the access request.
type AccessRequestHistoryCollection struct {

	// contains AccessRequestHistorySummary
	Items []AccessRequestHistorySummary `mandatory:"true" json:"items"`
}

func (m AccessRequestHistoryCollection) String() string {
	return common.PointerString(m)
}
