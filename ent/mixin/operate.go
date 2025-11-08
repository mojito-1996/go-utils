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
