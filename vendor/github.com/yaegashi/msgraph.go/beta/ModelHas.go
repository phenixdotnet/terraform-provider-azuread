// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// HasPayloadLinkResultItem undocumented
type HasPayloadLinkResultItem struct {
	// Object is the base model of HasPayloadLinkResultItem
	Object
	// PayloadID Key of the Payload, In the format of Guid.
	PayloadID *string `json:"payloadId,omitempty"`
	// HasLink Indicate whether a payload has any link or not.
	HasLink *bool `json:"hasLink,omitempty"`
	// Error Exception information indicates if check for this item was successful or not.Empty string for no error.
	Error *string `json:"error,omitempty"`
	// Sources The reason where the link comes from.
	Sources []DeviceAndAppManagementAssignmentSource `json:"sources,omitempty"`
}
