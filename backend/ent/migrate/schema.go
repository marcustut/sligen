// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// InstancesColumns holds the columns for the "instances" table.
	InstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "instance_id", Type: field.TypeString, Unique: true},
		{Name: "instance_type", Type: field.TypeString},
		{Name: "private_dns_name", Type: field.TypeString, Unique: true},
		{Name: "private_ip_address", Type: field.TypeString, Unique: true},
		{Name: "public_dns_name", Type: field.TypeString, Unique: true},
		{Name: "public_ip_address", Type: field.TypeString, Unique: true},
		{Name: "image_id", Type: field.TypeString},
		{Name: "architecture", Type: field.TypeString},
		{Name: "availability_zone", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "slide_instance", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "user_instances", Type: field.TypeString, Nullable: true},
	}
	// InstancesTable holds the schema information for the "instances" table.
	InstancesTable = &schema.Table{
		Name:       "instances",
		Columns:    InstancesColumns,
		PrimaryKey: []*schema.Column{InstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "instances_slides_instance",
				Columns:    []*schema.Column{InstancesColumns[12]},
				RefColumns: []*schema.Column{SlidesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "instances_users_instances",
				Columns:    []*schema.Column{InstancesColumns[13]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SlidesColumns holds the columns for the "slides" table.
	SlidesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "path_token", Type: field.TypeJSON, Nullable: true},
		{Name: "size", Type: field.TypeInt64, Nullable: true},
		{Name: "access_level", Type: field.TypeEnum, Enums: []string{"PRIVATE", "PUBLIC", "VIEW"}, Default: "PRIVATE"},
		{Name: "shared_with", Type: field.TypeJSON, SchemaType: map[string]string{"mysql": "json DEFAULT (JSON_ARRAY())"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
		{Name: "user_slides", Type: field.TypeString, Nullable: true},
	}
	// SlidesTable holds the schema information for the "slides" table.
	SlidesTable = &schema.Table{
		Name:       "slides",
		Columns:    SlidesColumns,
		PrimaryKey: []*schema.Column{SlidesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "slides_users_slides",
				Columns:    []*schema.Column{SlidesColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 50},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "full_name", Type: field.TypeString, Nullable: true, Size: 60},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true, Size: 2083},
		{Name: "bio", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		InstancesTable,
		SlidesTable,
		UsersTable,
	}
)

func init() {
	InstancesTable.ForeignKeys[0].RefTable = SlidesTable
	InstancesTable.ForeignKeys[1].RefTable = UsersTable
	SlidesTable.ForeignKeys[0].RefTable = UsersTable
}
