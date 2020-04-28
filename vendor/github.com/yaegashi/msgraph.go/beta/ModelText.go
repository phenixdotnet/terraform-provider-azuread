// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TextClassificationRequestObject undocumented
type TextClassificationRequestObject struct {
	// Entity is the base model of TextClassificationRequestObject
	Entity
	// Text undocumented
	Text *string `json:"text,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}

// TextColumn undocumented
type TextColumn struct {
	// Object is the base model of TextColumn
	Object
	// AllowMultipleLines undocumented
	AllowMultipleLines *bool `json:"allowMultipleLines,omitempty"`
	// AppendChangesToExistingText undocumented
	AppendChangesToExistingText *bool `json:"appendChangesToExistingText,omitempty"`
	// LinesForEditing undocumented
	LinesForEditing *int `json:"linesForEditing,omitempty"`
	// MaxLength undocumented
	MaxLength *int `json:"maxLength,omitempty"`
	// TextType undocumented
	TextType *string `json:"textType,omitempty"`
}
