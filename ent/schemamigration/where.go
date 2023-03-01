// Code generated by ent, DO NOT EDIT.

package schemamigration

import (
	"server/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(func(s *sql.Selector) {
		p(s.Not())
	})
}
