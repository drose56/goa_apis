// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the add service.
//
// Command:
// $ goa gen add/design

package server

import (
	"fmt"
)

// AddnumsAddPath returns the URL path to the add service addnums HTTP endpoint.
func AddnumsAddPath(a int, b int) string {
	return fmt.Sprintf("/addnums/%v/%v", a, b)
}
