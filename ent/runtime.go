// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/bluayer/GoingGong/ent/schema"
	"github.com/bluayer/GoingGong/ent/user"
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
}
