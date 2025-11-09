package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type Id struct {
	mixin.Schema
}

func (Id) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Comment("主键"),
	}
}

type UUId struct {
	mixin.Schema
}

func (UUId) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Comment("主键"),
	}
}
