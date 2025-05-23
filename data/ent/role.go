// Code generated by ent, DO NOT EDIT.

package ent

import (
	"coolfun-admin/data/ent/role"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID uint64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// status 1 normal 0 ban | 状态 1 正常 0 禁用
	Status uint8 `json:"status,omitempty"`
	// role name | 角色名
	Name string `json:"name,omitempty"`
	// role value for permission control in front end | 角色值，用于前端权限控制
	Value string `json:"value,omitempty"`
	// default menu : dashboard | 默认登录页面
	DefaultRouter string `json:"default_router,omitempty"`
	// remark | 备注
	Remark string `json:"remark,omitempty"`
	// order number | 排序编号
	OrderNo uint32 `json:"order_no,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Menus holds the value of the menus edge.
	Menus []*Menu `json:"menus,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) MenusOrErr() ([]*Menu, error) {
	if e.loadedTypes[0] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID, role.FieldStatus, role.FieldOrderNo:
			values[i] = new(sql.NullInt64)
		case role.FieldName, role.FieldValue, role.FieldDefaultRouter, role.FieldRemark:
			values[i] = new(sql.NullString)
		case role.FieldCreatedAt, role.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Role", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint64(value.Int64)
		case role.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case role.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case role.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = uint8(value.Int64)
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case role.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				r.Value = value.String
			}
		case role.FieldDefaultRouter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_router", values[i])
			} else if value.Valid {
				r.DefaultRouter = value.String
			}
		case role.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				r.Remark = value.String
			}
		case role.FieldOrderNo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_no", values[i])
			} else if value.Valid {
				r.OrderNo = uint32(value.Int64)
			}
		}
	}
	return nil
}

// QueryMenus queries the "menus" edge of the Role entity.
func (r *Role) QueryMenus() *MenuQuery {
	return (&RoleClient{config: r.config}).QueryMenus(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(r.Value)
	builder.WriteString(", ")
	builder.WriteString("default_router=")
	builder.WriteString(r.DefaultRouter)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(r.Remark)
	builder.WriteString(", ")
	builder.WriteString("order_no=")
	builder.WriteString(fmt.Sprintf("%v", r.OrderNo))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
