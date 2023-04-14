package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstName"),
		field.String("lastName"),
		field.String("email"),
		field.Float32("salary"),
		field.Time("birthDate"),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}
