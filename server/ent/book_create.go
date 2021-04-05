// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetPublishedAt sets the "publishedAt" field.
func (bc *BookCreate) SetPublishedAt(t time.Time) *BookCreate {
	bc.mutation.SetPublishedAt(t)
	return bc
}

// SetNillablePublishedAt sets the "publishedAt" field if the given value is not nil.
func (bc *BookCreate) SetNillablePublishedAt(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetPublishedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updatedAt" field.
func (bc *BookCreate) SetUpdatedAt(t time.Time) *BookCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bc *BookCreate) SetNillableUpdatedAt(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetAuthor sets the "author" field.
func (bc *BookCreate) SetAuthor(s string) *BookCreate {
	bc.mutation.SetAuthor(s)
	return bc
}

// SetDescription sets the "description" field.
func (bc *BookCreate) SetDescription(s string) *BookCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetEdition sets the "edition" field.
func (bc *BookCreate) SetEdition(i int) *BookCreate {
	bc.mutation.SetEdition(i)
	return bc
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetCollectionID sets the "collection" edge to the Collection entity by ID.
func (bc *BookCreate) SetCollectionID(id int) *BookCreate {
	bc.mutation.SetCollectionID(id)
	return bc
}

// SetNillableCollectionID sets the "collection" edge to the Collection entity by ID if the given value is not nil.
func (bc *BookCreate) SetNillableCollectionID(id *int) *BookCreate {
	if id != nil {
		bc = bc.SetCollectionID(*id)
	}
	return bc
}

// SetCollection sets the "collection" edge to the Collection entity.
func (bc *BookCreate) SetCollection(c *Collection) *BookCreate {
	return bc.SetCollectionID(c.ID)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	var (
		err  error
		node *Book
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() {
	if _, ok := bc.mutation.PublishedAt(); !ok {
		v := book.DefaultPublishedAt()
		bc.mutation.SetPublishedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := book.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "publishedAt", err: errors.New("ent: missing required field \"publishedAt\"")}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New("ent: missing required field \"updatedAt\"")}
	}
	if _, ok := bc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New("ent: missing required field \"author\"")}
	}
	if _, ok := bc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := bc.mutation.Edition(); !ok {
		return &ValidationError{Name: "edition", err: errors.New("ent: missing required field \"edition\"")}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: book.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.PublishedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldPublishedAt,
		})
		_node.PublishedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.Author(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldAuthor,
		})
		_node.Author = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := bc.mutation.Edition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: book.FieldEdition,
		})
		_node.Edition = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldTitle,
		})
		_node.Title = value
	}
	if nodes := bc.mutation.CollectionIDs(); len(nodes) > 0 {
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
		_node.collection_books = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
