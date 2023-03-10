// Code generated by ent, DO NOT EDIT.

package arinternalmetadatum

const (
	// Label holds the string label denoting the arinternalmetadatum type in the database.
	Label = "ar_internal_metadatum"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the arinternalmetadatum in the database.
	Table = "ar_internal_metadata"
)

// Columns holds all SQL columns for arinternalmetadatum fields.
var Columns = []string{
	FieldID,
	FieldValue,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
