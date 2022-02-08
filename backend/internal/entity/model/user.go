package model

import "github.com/marcustut/fyp/backend/ent"

// User is the model entity for the User schema.
type User = ent.User

// UserConnection is the connection containing edges to User.
type UserConnection = ent.UserConnection

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput = ent.UserWhereInput

// CreateUserInput represents a mutation input for creating User.
type CreateUserInput = ent.CreateUserInput

// UpdateUserInput represents a mutation input for updating User.
type UpdateUserInput = ent.UpdateUserInput
