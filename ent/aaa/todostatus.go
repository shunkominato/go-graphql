// Code generated by ent, DO NOT EDIT.

package aaa

import (
	"fmt"
	"server/ent/aaa/todostatus"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// TodoStatus is the model entity for the TodoStatus schema.
type TodoStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoStatusQuery when eager-loading is set.
	Edges TodoStatusEdges `json:"edges"`
}

// TodoStatusEdges holds the relations/edges for other nodes in the graph.
type TodoStatusEdges struct {
	// Todos holds the value of the todos edge.
	Todos []*Todo `json:"todos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TodosOrErr returns the Todos value or an error if the edge
// was not loaded in eager-loading.
func (e TodoStatusEdges) TodosOrErr() ([]*Todo, error) {
	if e.loadedTypes[0] {
		return e.Todos, nil
	}
	return nil, &NotLoadedError{edge: "todos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TodoStatus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todostatus.FieldID:
			values[i] = new(sql.NullInt64)
		case todostatus.FieldStatus:
			values[i] = new(sql.NullString)
		case todostatus.FieldCreatedAt, todostatus.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TodoStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TodoStatus fields.
func (ts *TodoStatus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todostatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int(value.Int64)
		case todostatus.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ts.Status = value.String
			}
		case todostatus.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ts.CreatedAt = value.Time
			}
		case todostatus.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ts.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTodos queries the "todos" edge of the TodoStatus entity.
func (ts *TodoStatus) QueryTodos() *TodoQuery {
	return NewTodoStatusClient(ts.config).QueryTodos(ts)
}

// Update returns a builder for updating this TodoStatus.
// Note that you need to call TodoStatus.Unwrap() before calling this method if this TodoStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TodoStatus) Update() *TodoStatusUpdateOne {
	return NewTodoStatusClient(ts.config).UpdateOne(ts)
}

// Unwrap unwraps the TodoStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TodoStatus) Unwrap() *TodoStatus {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("aaa: TodoStatus is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TodoStatus) String() string {
	var builder strings.Builder
	builder.WriteString("TodoStatus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("status=")
	builder.WriteString(ts.Status)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ts.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ts.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TodoStatusSlice is a parsable slice of TodoStatus.
type TodoStatusSlice []*TodoStatus
