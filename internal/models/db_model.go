package models

// This package contains data models used throughout the application.

// Bulls represents a data structure for earthquake data.
type Bulls struct {
	Id        int32   `json:"id,omitempty"`
	Event_id  string  `json:"event_id,omitempty"`
	Epicenter string  `json:"epicenter,omitempty"`
	Mag       float32 `json:"mag,omitempty"`
	Ev_date   string  `json:"ev_date,omitempty"`
	Ev_time   string  `json:"ev_time,omitempty"`
}
