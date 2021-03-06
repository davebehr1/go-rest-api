// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"lxdAssessmentServer/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where adds a new predicate for the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.predicates = append(bu.mutation.predicates, ps...)
	return bu
}

// SetPublishedAt sets the "publishedAt" field.
func (bu *BookUpdate) SetPublishedAt(t time.Time) *BookUpdate {
	bu.mutation.SetPublishedAt(t)
	return bu
}

// SetUpdatedAt sets the "updatedAt" field.
func (bu *BookUpdate) SetUpdatedAt(t time.Time) *BookUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetAuthor sets the "author" field.
func (bu *BookUpdate) SetAuthor(s string) *BookUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BookUpdate) SetDescription(s string) *BookUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetEdition sets the "edition" field.
func (bu *BookUpdate) SetEdition(i int) *BookUpdate {
	bu.mutation.ResetEdition()
	bu.mutation.SetEdition(i)
	return bu
}

// AddEdition adds i to the "edition" field.
func (bu *BookUpdate) AddEdition(i int) *BookUpdate {
	bu.mutation.AddEdition(i)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetCollectionID sets the "collection" edge to the Collection entity by ID.
func (bu *BookUpdate) SetCollectionID(id int) *BookUpdate {
	bu.mutation.SetCollectionID(id)
	return bu
}

// SetNillableCollectionID sets the "collection" edge to the Collection entity by ID if the given value is not nil.
func (bu *BookUpdate) SetNillableCollectionID(id *int) *BookUpdate {
	if id != nil {
		bu = bu.SetCollectionID(*id)
	}
	return bu
}

// SetCollection sets the "collection" edge to the Collection entity.
func (bu *BookUpdate) SetCollection(c *Collection) *BookUpdate {
	return bu.SetCollectionID(c.ID)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearCollection clears the "collection" edge to the Collection entity.
func (bu *BookUpdate) ClearCollection() *BookUpdate {
	bu.mutation.ClearCollection()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookUpdate) defaults() {
	if _, ok := bu.mutation.PublishedAt(); !ok {
		v := book.UpdateDefaultPublishedAt()
		bu.mutation.SetPublishedAt(v)
	}
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := book.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldPublishedAt,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldUpdatedAt,
		})
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldAuthor,
		})
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldDescription,
		})
	}
	if value, ok := bu.mutation.Edition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: book.FieldEdition,
		})
	}
	if value, ok := bu.mutation.AddedEdition(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: book.FieldEdition,
		})
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldTitle,
		})
	}
	if bu.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.CollectionTable,
			Columns: []string{book.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.CollectionTable,
			Columns: []string{book.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// SetPublishedAt sets the "publishedAt" field.
func (buo *BookUpdateOne) SetPublishedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetPublishedAt(t)
	return buo
}

// SetUpdatedAt sets the "updatedAt" field.
func (buo *BookUpdateOne) SetUpdatedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetAuthor sets the "author" field.
func (buo *BookUpdateOne) SetAuthor(s string) *BookUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BookUpdateOne) SetDescription(s string) *BookUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetEdition sets the "edition" field.
func (buo *BookUpdateOne) SetEdition(i int) *BookUpdateOne {
	buo.mutation.ResetEdition()
	buo.mutation.SetEdition(i)
	return buo
}

// AddEdition adds i to the "edition" field.
func (buo *BookUpdateOne) AddEdition(i int) *BookUpdateOne {
	buo.mutation.AddEdition(i)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetCollectionID sets the "collection" edge to the Collection entity by ID.
func (buo *BookUpdateOne) SetCollectionID(id int) *BookUpdateOne {
	buo.mutation.SetCollectionID(id)
	return buo
}

// SetNillableCollectionID sets the "collection" edge to the Collection entity by ID if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCollectionID(id *int) *BookUpdateOne {
	if id != nil {
		buo = buo.SetCollectionID(*id)
	}
	return buo
}

// SetCollection sets the "collection" edge to the Collection entity.
func (buo *BookUpdateOne) SetCollection(c *Collection) *BookUpdateOne {
	return buo.SetCollectionID(c.ID)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearCollection clears the "collection" edge to the Collection entity.
func (buo *BookUpdateOne) ClearCollection() *BookUpdateOne {
	buo.mutation.ClearCollection()
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	var (
		err  error
		node *Book
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookUpdateOne) defaults() {
	if _, ok := buo.mutation.PublishedAt(); !ok {
		v := book.UpdateDefaultPublishedAt()
		buo.mutation.SetPublishedAt(v)
	}
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := book.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Book.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.PublishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldPublishedAt,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldUpdatedAt,
		})
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldAuthor,
		})
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldDescription,
		})
	}
	if value, ok := buo.mutation.Edition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: book.FieldEdition,
		})
	}
	if value, ok := buo.mutation.AddedEdition(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: book.FieldEdition,
		})
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldTitle,
		})
	}
	if buo.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.CollectionTable,
			Columns: []string{book.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.CollectionTable,
			Columns: []string{book.CollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
