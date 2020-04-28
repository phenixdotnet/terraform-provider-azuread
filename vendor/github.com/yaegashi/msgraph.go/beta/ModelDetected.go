// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DetectedApp A managed or unmanaged app that is installed on a managed device. Unmanaged apps will only appear for devices marked as corporate owned.
type DetectedApp struct {
	// Entity is the base model of DetectedApp
	Entity
	// DisplayName Name of the discovered application. Read-only
	DisplayName *string `json:"displayName,omitempty"`
	// Version Version of the discovered application. Read-only
	Version *string `json:"version,omitempty"`
	// SizeInByte Discovered application size in bytes. Read-only
	SizeInByte *int `json:"sizeInByte,omitempty"`
	// DeviceCount The number of devices that have installed this application
	DeviceCount *int `json:"deviceCount,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []ManagedDevice `json:"managedDevices,omitempty"`
}

// DetectedSensitiveContent undocumented
type DetectedSensitiveContent struct {
	// Object is the base model of DetectedSensitiveContent
	Object
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UniqueCount undocumented
	UniqueCount *int `json:"uniqueCount,omitempty"`
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// RecommendedConfidence undocumented
	RecommendedConfidence *int `json:"recommendedConfidence,omitempty"`
	// Matches undocumented
	Matches []SensitiveContentLocation `json:"matches,omitempty"`
}

// DetectedSensitiveContentWrapper undocumented
type DetectedSensitiveContentWrapper struct {
	// Object is the base model of DetectedSensitiveContentWrapper
	Object
	// Classification undocumented
	Classification []DetectedSensitiveContent `json:"classification,omitempty"`
}
