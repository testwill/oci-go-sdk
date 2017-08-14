// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
        "time"
)


//UserGroupMembership An object that represents the membership of a user in a group. When you add a user to a group, the result is a\n`UserGroupMembership` with its own OCID. To remove a user from a group, you delete the `UserGroupMembership` object.\n

type UserGroupMembership struct {

    // The OCID of the membership.
    Id string `json:"id,omitempty"`

    // The OCID of the tenancy containing the user, group, and membership object.
    CompartmentId string `json:"compartmentId,omitempty"`

    // The OCID of the group.
    GroupId string `json:"groupId,omitempty"`

    // The OCID of the user.
    UserId string `json:"userId,omitempty"`

    // Date and time the membership was created, in the format defined by RFC3339.\n\nExample: `2016-08-25T21:10:29.600Z`\n
    TimeCreated time.Time `json:"timeCreated,omitempty"`

    // The membership's current state.  After creating a membership object, make sure its `lifecycleState` changes\nfrom CREATING to ACTIVE before using it.\n
    LifecycleState string `json:"lifecycleState,omitempty"`

    // The detailed status of INACTIVE lifecycleState.
    InactiveStatus int64 `json:"inactiveStatus,omitempty"`
}