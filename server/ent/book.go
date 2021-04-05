// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"lxdAssessmentServer/ent/book"
	"lxdAssessmentServer/ent/collection"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PublishedAt holds the value of the "publishedAt" field.
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edition holds the value of the "edition" field.
	Edition int `json:"edition,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges            BookEdges `json:"edges"`
	collection_books *int
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// Collection holds the value of the collection edge.
	Collection *Collection `json:"collection,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CollectionOrErr returns the Collection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) CollectionOrErr() (*Collection, error) {
	if e.loadedTypes[0] {
		if e.Collection == nil {
			// The edge collection was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: collection.Label}
		}
		return e.Collection, nil
	}
	return nil, &NotLoadedError{edge: "collection"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldID, book.FieldEdition:
			values[i] = &sql.NullInt64{}
		case book.FieldAuthor, book.FieldDescription, book.FieldTitle:
			values[i] = &sql.NullString{}
		case book.FieldPublishedAt, book.FieldUpdatedAt:
			values[i] = &sql.NullTime{}
		case book.ForeignKeys[0]: // collection_books
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Book", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case book.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field publishedAt", values[i])
			} else if value.Valid {
				b.PublishedAt = value.Time
			}
		case book.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case book.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				b.Author = value.String
			}
		case book.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case book.FieldEdition:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field edition", values[i])
			} else if value.Valid {
				b.Edition = int(value.Int64)
			}
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field collection_books", value)
			} else if value.Valid {
				b.collection_books = new(int)
				*b.collection_books = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCollection queries the "collection" edge of the Book entity.
func (b *Book) QueryCollection() *CollectionQuery {
	return (&BookClient{config: b.config}).QueryCollection(b)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", publishedAt=")
	builder.WriteString(b.PublishedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", author=")
	builder.WriteString(b.Author)
	builder.WriteString(", description=")
	builder.WriteString(b.Description)
	builder.WriteString(", edition=")
	builder.WriteString(fmt.Sprintf("%v", b.Edition))
	builder.WriteString(", title=")
	builder.WriteString(b.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
