// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v42/common"
)

// Action An object that represents an action to trigger for events that match a rule.
type Action interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the action.
	GetId() *string

	// A message generated by the Events service about the current state of this action.
	GetLifecycleMessage() *string

	// The current state of the rule.
	GetLifecycleState() ActionLifecycleStateEnum

	// Whether or not this action is currently enabled.
	// Example: `true`
	GetIsEnabled() *bool

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	GetDescription() *string
}

type action struct {
	JsonData         []byte
	Id               *string                  `mandatory:"true" json:"id"`
	LifecycleMessage *string                  `mandatory:"true" json:"lifecycleMessage"`
	LifecycleState   ActionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	IsEnabled        *bool                    `mandatory:"false" json:"isEnabled"`
	Description      *string                  `mandatory:"false" json:"description"`
	ActionType       string                   `json:"actionType"`
}

// UnmarshalJSON unmarshals json
func (m *action) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraction action
	s := struct {
		Model Unmarshaleraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.LifecycleMessage = s.Model.LifecycleMessage
	m.LifecycleState = s.Model.LifecycleState
	m.IsEnabled = s.Model.IsEnabled
	m.Description = s.Model.Description
	m.ActionType = s.Model.ActionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *action) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActionType {
	case "OSS":
		mm := StreamingServiceAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ONS":
		mm := NotificationServiceAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FAAS":
		mm := FaaSAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m action) GetId() *string {
	return m.Id
}

//GetLifecycleMessage returns LifecycleMessage
func (m action) GetLifecycleMessage() *string {
	return m.LifecycleMessage
}

//GetLifecycleState returns LifecycleState
func (m action) GetLifecycleState() ActionLifecycleStateEnum {
	return m.LifecycleState
}

//GetIsEnabled returns IsEnabled
func (m action) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m action) GetDescription() *string {
	return m.Description
}

func (m action) String() string {
	return common.PointerString(m)
}

// ActionLifecycleStateEnum Enum with underlying type: string
type ActionLifecycleStateEnum string

// Set of constants representing the allowable values for ActionLifecycleStateEnum
const (
	ActionLifecycleStateCreating ActionLifecycleStateEnum = "CREATING"
	ActionLifecycleStateActive   ActionLifecycleStateEnum = "ACTIVE"
	ActionLifecycleStateInactive ActionLifecycleStateEnum = "INACTIVE"
	ActionLifecycleStateUpdating ActionLifecycleStateEnum = "UPDATING"
	ActionLifecycleStateDeleting ActionLifecycleStateEnum = "DELETING"
	ActionLifecycleStateDeleted  ActionLifecycleStateEnum = "DELETED"
	ActionLifecycleStateFailed   ActionLifecycleStateEnum = "FAILED"
)

var mappingActionLifecycleState = map[string]ActionLifecycleStateEnum{
	"CREATING": ActionLifecycleStateCreating,
	"ACTIVE":   ActionLifecycleStateActive,
	"INACTIVE": ActionLifecycleStateInactive,
	"UPDATING": ActionLifecycleStateUpdating,
	"DELETING": ActionLifecycleStateDeleting,
	"DELETED":  ActionLifecycleStateDeleted,
	"FAILED":   ActionLifecycleStateFailed,
}

// GetActionLifecycleStateEnumValues Enumerates the set of values for ActionLifecycleStateEnum
func GetActionLifecycleStateEnumValues() []ActionLifecycleStateEnum {
	values := make([]ActionLifecycleStateEnum, 0)
	for _, v := range mappingActionLifecycleState {
		values = append(values, v)
	}
	return values
}

// ActionActionTypeEnum Enum with underlying type: string
type ActionActionTypeEnum string

// Set of constants representing the allowable values for ActionActionTypeEnum
const (
	ActionActionTypeOns  ActionActionTypeEnum = "ONS"
	ActionActionTypeOss  ActionActionTypeEnum = "OSS"
	ActionActionTypeFaas ActionActionTypeEnum = "FAAS"
)

var mappingActionActionType = map[string]ActionActionTypeEnum{
	"ONS":  ActionActionTypeOns,
	"OSS":  ActionActionTypeOss,
	"FAAS": ActionActionTypeFaas,
}

// GetActionActionTypeEnumValues Enumerates the set of values for ActionActionTypeEnum
func GetActionActionTypeEnumValues() []ActionActionTypeEnum {
	values := make([]ActionActionTypeEnum, 0)
	for _, v := range mappingActionActionType {
		values = append(values, v)
	}
	return values
}
