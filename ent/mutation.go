// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/superchris/ent-ogen-example/ent/person"
	"github.com/superchris/ent-ogen-example/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePerson = "Person"
)

// PersonMutation represents an operation that mutates the Person nodes in the graph.
type PersonMutation struct {
	config
	op            Op
	typ           string
	id            *int
	firstName     *string
	lastName      *string
	email         *string
	salary        *float32
	addsalary     *float32
	birthDate     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Person, error)
	predicates    []predicate.Person
}

var _ ent.Mutation = (*PersonMutation)(nil)

// personOption allows management of the mutation configuration using functional options.
type personOption func(*PersonMutation)

// newPersonMutation creates new mutation for the Person entity.
func newPersonMutation(c config, op Op, opts ...personOption) *PersonMutation {
	m := &PersonMutation{
		config:        c,
		op:            op,
		typ:           TypePerson,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPersonID sets the ID field of the mutation.
func withPersonID(id int) personOption {
	return func(m *PersonMutation) {
		var (
			err   error
			once  sync.Once
			value *Person
		)
		m.oldValue = func(ctx context.Context) (*Person, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Person.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPerson sets the old Person of the mutation.
func withPerson(node *Person) personOption {
	return func(m *PersonMutation) {
		m.oldValue = func(context.Context) (*Person, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PersonMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PersonMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PersonMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PersonMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Person.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetFirstName sets the "firstName" field.
func (m *PersonMutation) SetFirstName(s string) {
	m.firstName = &s
}

// FirstName returns the value of the "firstName" field in the mutation.
func (m *PersonMutation) FirstName() (r string, exists bool) {
	v := m.firstName
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "firstName" field's value of the Person entity.
// If the Person object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PersonMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "firstName" field.
func (m *PersonMutation) ResetFirstName() {
	m.firstName = nil
}

// SetLastName sets the "lastName" field.
func (m *PersonMutation) SetLastName(s string) {
	m.lastName = &s
}

// LastName returns the value of the "lastName" field in the mutation.
func (m *PersonMutation) LastName() (r string, exists bool) {
	v := m.lastName
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "lastName" field's value of the Person entity.
// If the Person object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PersonMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "lastName" field.
func (m *PersonMutation) ResetLastName() {
	m.lastName = nil
}

// SetEmail sets the "email" field.
func (m *PersonMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *PersonMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the Person entity.
// If the Person object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PersonMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *PersonMutation) ResetEmail() {
	m.email = nil
}

// SetSalary sets the "salary" field.
func (m *PersonMutation) SetSalary(f float32) {
	m.salary = &f
	m.addsalary = nil
}

// Salary returns the value of the "salary" field in the mutation.
func (m *PersonMutation) Salary() (r float32, exists bool) {
	v := m.salary
	if v == nil {
		return
	}
	return *v, true
}

// OldSalary returns the old "salary" field's value of the Person entity.
// If the Person object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PersonMutation) OldSalary(ctx context.Context) (v float32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSalary is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSalary requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSalary: %w", err)
	}
	return oldValue.Salary, nil
}

// AddSalary adds f to the "salary" field.
func (m *PersonMutation) AddSalary(f float32) {
	if m.addsalary != nil {
		*m.addsalary += f
	} else {
		m.addsalary = &f
	}
}

// AddedSalary returns the value that was added to the "salary" field in this mutation.
func (m *PersonMutation) AddedSalary() (r float32, exists bool) {
	v := m.addsalary
	if v == nil {
		return
	}
	return *v, true
}

// ResetSalary resets all changes to the "salary" field.
func (m *PersonMutation) ResetSalary() {
	m.salary = nil
	m.addsalary = nil
}

// SetBirthDate sets the "birthDate" field.
func (m *PersonMutation) SetBirthDate(t time.Time) {
	m.birthDate = &t
}

// BirthDate returns the value of the "birthDate" field in the mutation.
func (m *PersonMutation) BirthDate() (r time.Time, exists bool) {
	v := m.birthDate
	if v == nil {
		return
	}
	return *v, true
}

// OldBirthDate returns the old "birthDate" field's value of the Person entity.
// If the Person object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PersonMutation) OldBirthDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBirthDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBirthDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBirthDate: %w", err)
	}
	return oldValue.BirthDate, nil
}

// ResetBirthDate resets all changes to the "birthDate" field.
func (m *PersonMutation) ResetBirthDate() {
	m.birthDate = nil
}

// Where appends a list predicates to the PersonMutation builder.
func (m *PersonMutation) Where(ps ...predicate.Person) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PersonMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PersonMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Person, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PersonMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PersonMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Person).
func (m *PersonMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PersonMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.firstName != nil {
		fields = append(fields, person.FieldFirstName)
	}
	if m.lastName != nil {
		fields = append(fields, person.FieldLastName)
	}
	if m.email != nil {
		fields = append(fields, person.FieldEmail)
	}
	if m.salary != nil {
		fields = append(fields, person.FieldSalary)
	}
	if m.birthDate != nil {
		fields = append(fields, person.FieldBirthDate)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PersonMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case person.FieldFirstName:
		return m.FirstName()
	case person.FieldLastName:
		return m.LastName()
	case person.FieldEmail:
		return m.Email()
	case person.FieldSalary:
		return m.Salary()
	case person.FieldBirthDate:
		return m.BirthDate()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PersonMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case person.FieldFirstName:
		return m.OldFirstName(ctx)
	case person.FieldLastName:
		return m.OldLastName(ctx)
	case person.FieldEmail:
		return m.OldEmail(ctx)
	case person.FieldSalary:
		return m.OldSalary(ctx)
	case person.FieldBirthDate:
		return m.OldBirthDate(ctx)
	}
	return nil, fmt.Errorf("unknown Person field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PersonMutation) SetField(name string, value ent.Value) error {
	switch name {
	case person.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case person.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case person.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case person.FieldSalary:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSalary(v)
		return nil
	case person.FieldBirthDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBirthDate(v)
		return nil
	}
	return fmt.Errorf("unknown Person field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PersonMutation) AddedFields() []string {
	var fields []string
	if m.addsalary != nil {
		fields = append(fields, person.FieldSalary)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PersonMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case person.FieldSalary:
		return m.AddedSalary()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PersonMutation) AddField(name string, value ent.Value) error {
	switch name {
	case person.FieldSalary:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSalary(v)
		return nil
	}
	return fmt.Errorf("unknown Person numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PersonMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PersonMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PersonMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Person nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PersonMutation) ResetField(name string) error {
	switch name {
	case person.FieldFirstName:
		m.ResetFirstName()
		return nil
	case person.FieldLastName:
		m.ResetLastName()
		return nil
	case person.FieldEmail:
		m.ResetEmail()
		return nil
	case person.FieldSalary:
		m.ResetSalary()
		return nil
	case person.FieldBirthDate:
		m.ResetBirthDate()
		return nil
	}
	return fmt.Errorf("unknown Person field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PersonMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PersonMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PersonMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PersonMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PersonMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PersonMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PersonMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Person unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PersonMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Person edge %s", name)
}
