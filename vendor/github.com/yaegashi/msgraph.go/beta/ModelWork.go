// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkPosition undocumented
type WorkPosition struct {
	// ItemFacet is the base model of WorkPosition
	ItemFacet
	// Categories undocumented
	Categories []string `json:"categories,omitempty"`
	// Detail undocumented
	Detail *PositionDetail `json:"detail,omitempty"`
}
