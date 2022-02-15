// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"strings"
)

// SecurityRatingEnum Enum with underlying type: string
type SecurityRatingEnum string

// Set of constants representing the allowable values for SecurityRatingEnum
const (
	SecurityRatingExcellent SecurityRatingEnum = "EXCELLENT"
	SecurityRatingGood      SecurityRatingEnum = "GOOD"
	SecurityRatingFair      SecurityRatingEnum = "FAIR"
	SecurityRatingPoor      SecurityRatingEnum = "POOR"
	SecurityRatingNa        SecurityRatingEnum = "NA"
)

var mappingSecurityRatingEnum = map[string]SecurityRatingEnum{
	"EXCELLENT": SecurityRatingExcellent,
	"GOOD":      SecurityRatingGood,
	"FAIR":      SecurityRatingFair,
	"POOR":      SecurityRatingPoor,
	"NA":        SecurityRatingNa,
}

// GetSecurityRatingEnumValues Enumerates the set of values for SecurityRatingEnum
func GetSecurityRatingEnumValues() []SecurityRatingEnum {
	values := make([]SecurityRatingEnum, 0)
	for _, v := range mappingSecurityRatingEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityRatingEnumStringValues Enumerates the set of values in String for SecurityRatingEnum
func GetSecurityRatingEnumStringValues() []string {
	return []string{
		"EXCELLENT",
		"GOOD",
		"FAIR",
		"POOR",
		"NA",
	}
}

// GetMappingSecurityRatingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityRatingEnum(val string) (SecurityRatingEnum, bool) {
	mappingSecurityRatingEnumIgnoreCase := make(map[string]SecurityRatingEnum)
	for k, v := range mappingSecurityRatingEnum {
		mappingSecurityRatingEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSecurityRatingEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
