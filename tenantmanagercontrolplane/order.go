// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// Order Order details.
type Order struct {

	// Immutable and unique order number holding customer subscription information.
	OrderNumber *string `mandatory:"true" json:"orderNumber"`

	// Administrator email owning the subscription.
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// State of the order.
	OrderState *string `mandatory:"true" json:"orderState"`

	// Array of subscriptions associated with the order.
	Subscriptions []SubscriptionInfo `mandatory:"true" json:"subscriptions"`

	// Order's data center region.
	DataCenterRegion *string `mandatory:"false" json:"dataCenterRegion"`
}

func (m Order) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Order) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
