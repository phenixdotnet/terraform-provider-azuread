// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OnlineMeeting undocumented
type OnlineMeeting struct {
	// Entity is the base model of OnlineMeeting
	Entity
	// CreationDateTime undocumented
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// CanceledDateTime undocumented
	CanceledDateTime *time.Time `json:"canceledDateTime,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// EntryExitAnnouncement undocumented
	EntryExitAnnouncement *bool `json:"entryExitAnnouncement,omitempty"`
	// JoinURL undocumented
	JoinURL *string `json:"joinUrl,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
	// IsCancelled undocumented
	IsCancelled *bool `json:"isCancelled,omitempty"`
	// Participants undocumented
	Participants *MeetingParticipants `json:"participants,omitempty"`
	// IsBroadcast undocumented
	IsBroadcast *bool `json:"isBroadcast,omitempty"`
	// AccessLevel undocumented
	AccessLevel *AccessLevel `json:"accessLevel,omitempty"`
	// Capabilities undocumented
	Capabilities []MeetingCapabilities `json:"capabilities,omitempty"`
	// AudioConferencing undocumented
	AudioConferencing *AudioConferencing `json:"audioConferencing,omitempty"`
	// ChatInfo undocumented
	ChatInfo *ChatInfo `json:"chatInfo,omitempty"`
	// VideoTeleconferenceID undocumented
	VideoTeleconferenceID *string `json:"videoTeleconferenceId,omitempty"`
}

// OnlineMeetingInfo undocumented
type OnlineMeetingInfo struct {
	// Object is the base model of OnlineMeetingInfo
	Object
	// JoinURL undocumented
	JoinURL *string `json:"joinUrl,omitempty"`
	// ConferenceID undocumented
	ConferenceID *string `json:"conferenceId,omitempty"`
	// TollNumber undocumented
	TollNumber *string `json:"tollNumber,omitempty"`
	// TollFreeNumbers undocumented
	TollFreeNumbers []string `json:"tollFreeNumbers,omitempty"`
	// QuickDial undocumented
	QuickDial *string `json:"quickDial,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
}
