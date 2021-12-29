package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"reflect"
	"strings"
)

const SPACE string = " "
const EdgedbTag string = "edgedb"

// Predicate
// function type that takes a string and return a boolean
type Predicate = func(string) bool

// Replacement
// Representation of remplacement of a GraphQl query field name by another one
//
// Name : name of the edgdDB entity field tag
// HasChild: True, if the field is a complex type
type Replacement struct {
	Name     string
	HasChild bool
}

// Replace
// Function type that takes a string and returns a pointer of Replacement
type Replace = func(string) *Replacement

// RequestedFields
// Return the fields requested by the QraphQl query
// as a string format as below :
// {
//   Name,
//   Age,
//   Friends {
//       Name
//   }
// }
//
func RequestedFields(ctx context.Context, predicate Predicate, replaceFunc Replace) string {

	// Initialize a string builder
	b := &strings.Builder{}

	// For each collected fields in the GraphQl query
	for _, field := range graphql.CollectFieldsCtx(ctx, nil) {

		// do ...
		RequestedField(ctx, field, 1, b, predicate, replaceFunc)
	}

	// Return the builder in string format
	return b.String()
}

// RequestedField
// Append a field corresponding to the GraphQl query field provided
// into the String builder provided
func RequestedField(ctx context.Context, field graphql.CollectedField, depth int, builder *strings.Builder, predicate Predicate, replaceFunc Replace) {

	// Retrieve the field name
	// and initialize a variable hasChild at true
	fieldName := field.Name
	hasChild := true

	// If a replacement function is provided for this field name
	// replace fieldName and hasChild variables
	if r := replaceFunc(fieldName); r != nil {
		fieldName = r.Name
		hasChild = r.HasChild
	}

	// If an egde tag corresponding to the fieldname is found
	if predicate(fieldName) {

		// Write the fieldname into the provided stringBuilder
		builder.WriteString(strings.Repeat(SPACE, depth))
		builder.WriteString(fieldName)

		// If the current field has fields
		fields := graphql.CollectFields(graphql.GetOperationContext(ctx), field.Selections, nil)

		if len(fields) > 0 && hasChild {
			builder.WriteString(":{\n")

			// Do a recursive call on each sub-fields
			for _, selection := range fields {
				RequestedField(ctx, selection, depth+1, builder, predicate, replaceFunc)
			}

			builder.WriteString(strings.Repeat(SPACE, depth))
			builder.WriteString("}")
		}

		builder.WriteString(",\n")
	}
}

// PredicateFor
// Return a predicate that indicate
// if a field name provide in GraphQl query
// corresponds to a EdgeDB Entity field tag
func PredicateFor(i interface{}) Predicate {

	// name is the name of GraphQl query field
	return func(name string) bool {

		// Use reflection to retrieve type of provided interface
		t := reflect.TypeOf(i)

		// t.NumField() method returns the number of fields of the of t
		for i := 0; i < t.NumField(); i++ {

			// Get the tag `edgedb` on the field at index i
			tag := t.Field(i).Tag.Get(EdgedbTag)

			// If the edgedb tag corresonds to the given name
			// then return true
			if tag == name {
				return true
			}
		}

		return false
	}
}
