// Code generated by ent, DO NOT EDIT.

package event

import (
	"geulSsi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// Contents applies equality check predicate on the "contents" field. It's identical to ContentsEQ.
func Contents(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldContents, v))
}

// Writer applies equality check predicate on the "writer" field. It's identical to WriterEQ.
func Writer(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldWriter, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContentsEQ applies the EQ predicate on the "contents" field.
func ContentsEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldContents, v))
}

// ContentsNEQ applies the NEQ predicate on the "contents" field.
func ContentsNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldContents, v))
}

// ContentsIn applies the In predicate on the "contents" field.
func ContentsIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldContents, vs...))
}

// ContentsNotIn applies the NotIn predicate on the "contents" field.
func ContentsNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldContents, vs...))
}

// ContentsGT applies the GT predicate on the "contents" field.
func ContentsGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldContents, v))
}

// ContentsGTE applies the GTE predicate on the "contents" field.
func ContentsGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldContents, v))
}

// ContentsLT applies the LT predicate on the "contents" field.
func ContentsLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldContents, v))
}

// ContentsLTE applies the LTE predicate on the "contents" field.
func ContentsLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldContents, v))
}

// ContentsContains applies the Contains predicate on the "contents" field.
func ContentsContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldContents, v))
}

// ContentsHasPrefix applies the HasPrefix predicate on the "contents" field.
func ContentsHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldContents, v))
}

// ContentsHasSuffix applies the HasSuffix predicate on the "contents" field.
func ContentsHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldContents, v))
}

// ContentsEqualFold applies the EqualFold predicate on the "contents" field.
func ContentsEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldContents, v))
}

// ContentsContainsFold applies the ContainsFold predicate on the "contents" field.
func ContentsContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldContents, v))
}

// WriterEQ applies the EQ predicate on the "writer" field.
func WriterEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldWriter, v))
}

// WriterNEQ applies the NEQ predicate on the "writer" field.
func WriterNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldWriter, v))
}

// WriterIn applies the In predicate on the "writer" field.
func WriterIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldWriter, vs...))
}

// WriterNotIn applies the NotIn predicate on the "writer" field.
func WriterNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldWriter, vs...))
}

// WriterGT applies the GT predicate on the "writer" field.
func WriterGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldWriter, v))
}

// WriterGTE applies the GTE predicate on the "writer" field.
func WriterGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldWriter, v))
}

// WriterLT applies the LT predicate on the "writer" field.
func WriterLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldWriter, v))
}

// WriterLTE applies the LTE predicate on the "writer" field.
func WriterLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldWriter, v))
}

// WriterContains applies the Contains predicate on the "writer" field.
func WriterContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldWriter, v))
}

// WriterHasPrefix applies the HasPrefix predicate on the "writer" field.
func WriterHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldWriter, v))
}

// WriterHasSuffix applies the HasSuffix predicate on the "writer" field.
func WriterHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldWriter, v))
}

// WriterEqualFold applies the EqualFold predicate on the "writer" field.
func WriterEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldWriter, v))
}

// WriterContainsFold applies the ContainsFold predicate on the "writer" field.
func WriterContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldWriter, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}
