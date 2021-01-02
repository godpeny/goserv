// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/godpeny/goserv/ent/apiresponse"
)

// APIResponse is the model entity for the APIResponse schema.
type APIResponse struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code int `json:"code,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*APIResponse) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // code
		&sql.NullString{}, // type
		&sql.NullString{}, // message
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the APIResponse fields.
func (ar *APIResponse) assignValues(values ...interface{}) error {
	if m, n := len(values), len(apiresponse.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ar.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field code", values[0])
	} else if value.Valid {
		ar.Code = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		ar.Type = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field message", values[2])
	} else if value.Valid {
		ar.Message = value.String
	}
	return nil
}

// Update returns a builder for updating this APIResponse.
// Note that, you need to call APIResponse.Unwrap() before calling this method, if this APIResponse
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *APIResponse) Update() *APIResponseUpdateOne {
	return (&APIResponseClient{config: ar.config}).UpdateOne(ar)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ar *APIResponse) Unwrap() *APIResponse {
	tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: APIResponse is not a transactional entity")
	}
	ar.config.driver = tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *APIResponse) String() string {
	var builder strings.Builder
	builder.WriteString("APIResponse(")
	builder.WriteString(fmt.Sprintf("id=%v", ar.ID))
	builder.WriteString(", code=")
	builder.WriteString(fmt.Sprintf("%v", ar.Code))
	builder.WriteString(", type=")
	builder.WriteString(ar.Type)
	builder.WriteString(", message=")
	builder.WriteString(ar.Message)
	builder.WriteByte(')')
	return builder.String()
}

// APIResponses is a parsable slice of APIResponse.
type APIResponses []*APIResponse

func (ar APIResponses) config(cfg config) {
	for _i := range ar {
		ar[_i].config = cfg
	}
}