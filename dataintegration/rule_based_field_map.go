// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v51/common"
)

// RuleBasedFieldMap A map of rule patterns.
type RuleBasedFieldMap struct {

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The pattern to map from.
	FromPattern *string `mandatory:"false" json:"fromPattern"`

	// The pattern to map to.
	ToPattern *string `mandatory:"false" json:"toPattern"`

	// Specifies whether the rule uses a java regex syntax.
	IsJavaRegexSyntax *bool `mandatory:"false" json:"isJavaRegexSyntax"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	FromRuleConfig *RuleTypeConfig `mandatory:"false" json:"fromRuleConfig"`

	ToRuleConfig *RuleTypeConfig `mandatory:"false" json:"toRuleConfig"`

	// mapType
	MapType RuleBasedFieldMapMapTypeEnum `mandatory:"false" json:"mapType,omitempty"`
}

//GetDescription returns Description
func (m RuleBasedFieldMap) GetDescription() *string {
	return m.Description
}

func (m RuleBasedFieldMap) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m RuleBasedFieldMap) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRuleBasedFieldMap RuleBasedFieldMap
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeRuleBasedFieldMap
	}{
		"RULE_BASED_FIELD_MAP",
		(MarshalTypeRuleBasedFieldMap)(m),
	}

	return json.Marshal(&s)
}

// RuleBasedFieldMapMapTypeEnum Enum with underlying type: string
type RuleBasedFieldMapMapTypeEnum string

// Set of constants representing the allowable values for RuleBasedFieldMapMapTypeEnum
const (
	RuleBasedFieldMapMapTypeMapbyname     RuleBasedFieldMapMapTypeEnum = "MAPBYNAME"
	RuleBasedFieldMapMapTypeMapbyposition RuleBasedFieldMapMapTypeEnum = "MAPBYPOSITION"
	RuleBasedFieldMapMapTypeMapbypattern  RuleBasedFieldMapMapTypeEnum = "MAPBYPATTERN"
)

var mappingRuleBasedFieldMapMapType = map[string]RuleBasedFieldMapMapTypeEnum{
	"MAPBYNAME":     RuleBasedFieldMapMapTypeMapbyname,
	"MAPBYPOSITION": RuleBasedFieldMapMapTypeMapbyposition,
	"MAPBYPATTERN":  RuleBasedFieldMapMapTypeMapbypattern,
}

// GetRuleBasedFieldMapMapTypeEnumValues Enumerates the set of values for RuleBasedFieldMapMapTypeEnum
func GetRuleBasedFieldMapMapTypeEnumValues() []RuleBasedFieldMapMapTypeEnum {
	values := make([]RuleBasedFieldMapMapTypeEnum, 0)
	for _, v := range mappingRuleBasedFieldMapMapType {
		values = append(values, v)
	}
	return values
}
