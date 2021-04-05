// Code generated by entc, DO NOT EDIT.

package book

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeCollection holds the string denoting the collection edge name in mutations.
	EdgeCollection = "collection"
	// Table holds the table name of the book in the database.
	Table = "books"
	// CollectionTable is the table the holds the collection relation/edge.
	CollectionTable = "books"
	// CollectionInverseTable is the table name for the Collection entity.
	// It exists in this package in order to avoid circular dependency with the "collection" package.
	CollectionInverseTable = "collections"
	// CollectionColumn is the table column denoting the collection relation/edge.
	CollectionColumn = "collection_books"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldAuthor,
	FieldDescription,
	FieldTitle,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "books"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"collection_books",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
