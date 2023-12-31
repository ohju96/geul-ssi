// Code generated by ent, DO NOT EDIT.

package wisesaying

import (
	"geulSsi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLTE(FieldID, id))
}

// WiseSaying applies equality check predicate on the "wise_saying" field. It's identical to WiseSayingEQ.
func WiseSaying(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldWiseSaying, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldUpdatedAt, v))
}

// WiseSayingEQ applies the EQ predicate on the "wise_saying" field.
func WiseSayingEQ(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldWiseSaying, v))
}

// WiseSayingNEQ applies the NEQ predicate on the "wise_saying" field.
func WiseSayingNEQ(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNEQ(FieldWiseSaying, v))
}

// WiseSayingIn applies the In predicate on the "wise_saying" field.
func WiseSayingIn(vs ...string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldIn(FieldWiseSaying, vs...))
}

// WiseSayingNotIn applies the NotIn predicate on the "wise_saying" field.
func WiseSayingNotIn(vs ...string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNotIn(FieldWiseSaying, vs...))
}

// WiseSayingGT applies the GT predicate on the "wise_saying" field.
func WiseSayingGT(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGT(FieldWiseSaying, v))
}

// WiseSayingGTE applies the GTE predicate on the "wise_saying" field.
func WiseSayingGTE(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGTE(FieldWiseSaying, v))
}

// WiseSayingLT applies the LT predicate on the "wise_saying" field.
func WiseSayingLT(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLT(FieldWiseSaying, v))
}

// WiseSayingLTE applies the LTE predicate on the "wise_saying" field.
func WiseSayingLTE(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLTE(FieldWiseSaying, v))
}

// WiseSayingContains applies the Contains predicate on the "wise_saying" field.
func WiseSayingContains(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldContains(FieldWiseSaying, v))
}

// WiseSayingHasPrefix applies the HasPrefix predicate on the "wise_saying" field.
func WiseSayingHasPrefix(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldHasPrefix(FieldWiseSaying, v))
}

// WiseSayingHasSuffix applies the HasSuffix predicate on the "wise_saying" field.
func WiseSayingHasSuffix(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldHasSuffix(FieldWiseSaying, v))
}

// WiseSayingEqualFold applies the EqualFold predicate on the "wise_saying" field.
func WiseSayingEqualFold(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEqualFold(FieldWiseSaying, v))
}

// WiseSayingContainsFold applies the ContainsFold predicate on the "wise_saying" field.
func WiseSayingContainsFold(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldContainsFold(FieldWiseSaying, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldContainsFold(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.WiseSaying {
	return predicate.WiseSaying(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WiseSaying) predicate.WiseSaying {
	return predicate.WiseSaying(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WiseSaying) predicate.WiseSaying {
	return predicate.WiseSaying(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WiseSaying) predicate.WiseSaying {
	return predicate.WiseSaying(sql.NotPredicates(p))
}
