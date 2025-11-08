package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type CreatedBy struct {
	mixin.Schema
}

func (CreatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("create_by").
			Comment("创建者ID").
			Optional().
			Nillable(),
		field.Time("create_time").
			Comment("创建时间").
			Optional().
			Default(time.Now).
			Nillable(),
	}
}

type UpdatedBy struct {
	mixin.Schema
}

func (UpdatedBy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("update_by").
			Comment("更新者ID").
			Optional().
			Nillable(),
		field.Time("update_time").
			Comment("更新时间").
			Optional().
			Default(time.Now).
			Nillable(),
	}
}

type DeleteBy struct {
	mixin.Schema
}

func (DeleteBy) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("delete_fal").
			Comment("删除标识：0-正常，1-删除").
			Default(false),
		field.Uint32("delete_by").
			Comment("删除者ID").
			Optional().
			Nillable(),
		field.Time("delete_time").
			Comment("删除时间").
			Optional().
			Nillable(),
	}
}
