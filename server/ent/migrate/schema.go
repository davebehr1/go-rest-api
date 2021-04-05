// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "collection_books", Type: field.TypeInt, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "books_collections_books",
				Columns:    []*schema.Column{BooksColumns[4]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:        "collections",
		Columns:     CollectionsColumns,
		PrimaryKey:  []*schema.Column{CollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BooksTable,
		CollectionsTable,
	}
)

func init() {
	BooksTable.ForeignKeys[0].RefTable = CollectionsTable
}
