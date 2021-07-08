// Code generated by entc, DO NOT EDIT.

package ent

import (
	"pingpong/ent/schema"
	"pingpong/ent/user"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[0].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescPingCnt is the schema descriptor for pingCnt field.
	userDescPingCnt := userFields[2].Descriptor()
	// user.PingCntValidator is a validator for the "pingCnt" field. It is called by the builders before save.
	user.PingCntValidator = userDescPingCnt.Validators[0].(func(int) error)
}
