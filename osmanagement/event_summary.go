// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v51/common"
)

// EventSummary Summary of the Event.
type EventSummary struct {

	// OCID identifier of the event
	Id *string `mandatory:"true" json:"id"`

	// Unique OCI identifier of the instance where the event occurred
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Type of the event.
	EventType EventTypeEnum `mandatory:"true" json:"eventType"`

	// human readable description of the event
	Summary *string `mandatory:"false" json:"summary"`

	// Event occurrence count. Number of time the same event happened on the system.
	Count *int `mandatory:"false" json:"count"`

	// Time of the occurrence of the event
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EventSummary) String() string {
	return common.PointerString(m)
}
