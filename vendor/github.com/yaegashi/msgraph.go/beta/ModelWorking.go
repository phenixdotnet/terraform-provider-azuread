// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkingHours undocumented
type WorkingHours struct {
	// Object is the base model of WorkingHours
	Object
	// DaysOfWeek undocumented
	DaysOfWeek []DayOfWeek `json:"daysOfWeek,omitempty"`
	// StartTime undocumented
	StartTime *TimeOfDay `json:"startTime,omitempty"`
	// EndTime undocumented
	EndTime *TimeOfDay `json:"endTime,omitempty"`
	// TimeZone undocumented
	TimeZone *TimeZoneBase `json:"timeZone,omitempty"`
}
