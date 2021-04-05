// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"lxdAssessmentServer/ent/predicate"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBook       = "Book"
	TypeCollection = "Collection"
)

// BookMutation represents an operation that mutates the Book nodes in the graph.
type BookMutation struct {
	config
	op            Op
	typ           string
	id            *int
	author        *string
	description   *string
	title         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Book, error)
	predicates    []predicate.Book
}

var _ ent.Mutation = (*BookMutation)(nil)

// bookOption allows management of the mutation configuration using functional options.
type bookOption func(*BookMutation)

// newBookMutation creates new mutation for the Book entity.
func newBookMutation(c config, op Op, opts ...bookOption) *BookMutation {
	m := &BookMutation{
		config:        c,
		op:            op,
		typ:           TypeBook,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBookID sets the ID field of the mutation.
func withBookID(id int) bookOption {
	return func(m *BookMutation) {
		var (
			err   error
			once  sync.Once
			value *Book
		)
		m.oldValue = func(ctx context.Context) (*Book, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Book.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBook sets the old Book of the mutation.
func withBook(node *Book) bookOption {
	return func(m *BookMutation) {
		m.oldValue = func(context.Context) (*Book, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BookMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BookMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *BookMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAuthor sets the "author" field.
func (m *BookMutation) SetAuthor(s string) {
	m.author = &s
}

// Author returns the value of the "author" field in the mutation.
func (m *BookMutation) Author() (r string, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthor returns the old "author" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldAuthor(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAuthor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAuthor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthor: %w", err)
	}
	return oldValue.Author, nil
}

// ResetAuthor resets all changes to the "author" field.
func (m *BookMutation) ResetAuthor() {
	m.author = nil
}

// SetDescription sets the "description" field.
func (m *BookMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *BookMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *BookMutation) ResetDescription() {
	m.description = nil
}

// SetTitle sets the "title" field.
func (m *BookMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *BookMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Book entity.
// If the Book object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BookMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *BookMutation) ResetTitle() {
	m.title = nil
}

// Op returns the operation name.
func (m *BookMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Book).
func (m *BookMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BookMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.author != nil {
		fields = append(fields, book.FieldAuthor)
	}
	if m.description != nil {
		fields = append(fields, book.FieldDescription)
	}
	if m.title != nil {
		fields = append(fields, book.FieldTitle)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BookMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case book.FieldAuthor:
		return m.Author()
	case book.FieldDescription:
		return m.Description()
	case book.FieldTitle:
		return m.Title()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BookMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case book.FieldAuthor:
		return m.OldAuthor(ctx)
	case book.FieldDescription:
		return m.OldDescription(ctx)
	case book.FieldTitle:
		return m.OldTitle(ctx)
	}
	return nil, fmt.Errorf("unknown Book field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookMutation) SetField(name string, value ent.Value) error {
	switch name {
	case book.FieldAuthor:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthor(v)
		return nil
	case book.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case book.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BookMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BookMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BookMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Book numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BookMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BookMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BookMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Book nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BookMutation) ResetField(name string) error {
	switch name {
	case book.FieldAuthor:
		m.ResetAuthor()
		return nil
	case book.FieldDescription:
		m.ResetDescription()
		return nil
	case book.FieldTitle:
		m.ResetTitle()
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BookMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BookMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BookMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BookMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BookMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BookMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BookMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Book unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BookMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Book edge %s", name)
}

// CollectionMutation represents an operation that mutates the Collection nodes in the graph.
type CollectionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Collection, error)
	predicates    []predicate.Collection
}

var _ ent.Mutation = (*CollectionMutation)(nil)

// collectionOption allows management of the mutation configuration using functional options.
type collectionOption func(*CollectionMutation)

// newCollectionMutation creates new mutation for the Collection entity.
func newCollectionMutation(c config, op Op, opts ...collectionOption) *CollectionMutation {
	m := &CollectionMutation{
		config:        c,
		op:            op,
		typ:           TypeCollection,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCollectionID sets the ID field of the mutation.
func withCollectionID(id int) collectionOption {
	return func(m *CollectionMutation) {
		var (
			err   error
			once  sync.Once
			value *Collection
		)
		m.oldValue = func(ctx context.Context) (*Collection, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Collection.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCollection sets the old Collection of the mutation.
func withCollection(node *Collection) collectionOption {
	return func(m *CollectionMutation) {
		m.oldValue = func(context.Context) (*Collection, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CollectionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CollectionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *CollectionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *CollectionMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CollectionMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Collection entity.
// If the Collection object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CollectionMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *CollectionMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *CollectionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Collection).
func (m *CollectionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CollectionMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, collection.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CollectionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case collection.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CollectionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case collection.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Collection field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CollectionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case collection.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Collection field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CollectionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CollectionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CollectionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Collection numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CollectionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CollectionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CollectionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Collection nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CollectionMutation) ResetField(name string) error {
	switch name {
	case collection.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Collection field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CollectionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CollectionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CollectionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CollectionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CollectionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CollectionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CollectionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Collection unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CollectionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Collection edge %s", name)
}
