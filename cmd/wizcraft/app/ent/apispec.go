// Code generated by ent, DO NOT EDIT.

package ent

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/apispec"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/generalspec"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/project"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/service"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// APISpec is the model entity for the APISpec schema.
type APISpec struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// OpenapiSpec holds the value of the "openapi_spec" field.
	OpenapiSpec []uint8 `json:"openapi_spec,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the APISpecQuery when eager-loading is set.
	Edges                APISpecEdges `json:"edges"`
	general_spec_apispec *int
	project_apispecs     *uuid.UUID
	service_apispec      *uuid.UUID
	selectValues         sql.SelectValues
}

// APISpecEdges holds the relations/edges for other nodes in the graph.
type APISpecEdges struct {
	// Service holds the value of the service edge.
	Service *Service `json:"service,omitempty"`
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Generalspec holds the value of the generalspec edge.
	Generalspec *GeneralSpec `json:"generalspec,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ServiceOrErr returns the Service value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APISpecEdges) ServiceOrErr() (*Service, error) {
	if e.Service != nil {
		return e.Service, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: service.Label}
	}
	return nil, &NotLoadedError{edge: "service"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APISpecEdges) ProjectOrErr() (*Project, error) {
	if e.Project != nil {
		return e.Project, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: project.Label}
	}
	return nil, &NotLoadedError{edge: "project"}
}

// GeneralspecOrErr returns the Generalspec value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e APISpecEdges) GeneralspecOrErr() (*GeneralSpec, error) {
	if e.Generalspec != nil {
		return e.Generalspec, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: generalspec.Label}
	}
	return nil, &NotLoadedError{edge: "generalspec"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*APISpec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apispec.FieldOpenapiSpec:
			values[i] = new([]byte)
		case apispec.FieldID:
			values[i] = new(uuid.UUID)
		case apispec.ForeignKeys[0]: // general_spec_apispec
			values[i] = new(sql.NullInt64)
		case apispec.ForeignKeys[1]: // project_apispecs
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case apispec.ForeignKeys[2]: // service_apispec
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the APISpec fields.
func (as *APISpec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apispec.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				as.ID = *value
			}
		case apispec.FieldOpenapiSpec:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field openapi_spec", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &as.OpenapiSpec); err != nil {
					return fmt.Errorf("unmarshal field openapi_spec: %w", err)
				}
			}
		case apispec.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field general_spec_apispec", value)
			} else if value.Valid {
				as.general_spec_apispec = new(int)
				*as.general_spec_apispec = int(value.Int64)
			}
		case apispec.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field project_apispecs", values[i])
			} else if value.Valid {
				as.project_apispecs = new(uuid.UUID)
				*as.project_apispecs = *value.S.(*uuid.UUID)
			}
		case apispec.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field service_apispec", values[i])
			} else if value.Valid {
				as.service_apispec = new(uuid.UUID)
				*as.service_apispec = *value.S.(*uuid.UUID)
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the APISpec.
// This includes values selected through modifiers, order, etc.
func (as *APISpec) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// QueryService queries the "service" edge of the APISpec entity.
func (as *APISpec) QueryService() *ServiceQuery {
	return NewAPISpecClient(as.config).QueryService(as)
}

// QueryProject queries the "project" edge of the APISpec entity.
func (as *APISpec) QueryProject() *ProjectQuery {
	return NewAPISpecClient(as.config).QueryProject(as)
}

// QueryGeneralspec queries the "generalspec" edge of the APISpec entity.
func (as *APISpec) QueryGeneralspec() *GeneralSpecQuery {
	return NewAPISpecClient(as.config).QueryGeneralspec(as)
}

// Update returns a builder for updating this APISpec.
// Note that you need to call APISpec.Unwrap() before calling this method if this APISpec
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *APISpec) Update() *APISpecUpdateOne {
	return NewAPISpecClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the APISpec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *APISpec) Unwrap() *APISpec {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: APISpec is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *APISpec) String() string {
	var builder strings.Builder
	builder.WriteString("APISpec(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("openapi_spec=")
	builder.WriteString(fmt.Sprintf("%v", as.OpenapiSpec))
	builder.WriteByte(')')
	return builder.String()
}

// APISpecs is a parsable slice of APISpec.
type APISpecs []*APISpec
