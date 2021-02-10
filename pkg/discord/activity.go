package discord

// Activity defines the Activity sent with GatewayStatusUpdate.
//
// https://discord.com/developers/docs/topics/gateway#activity-object
type Activity struct {
	Name string 		`json:"name"`
	Type ActivityType 	`json:"type"`
	Url string 			`json:"url, omitempty"`
}

// ActivityType is the type of activity in the Activity struct.
//
// See ActivityType* consts.
//
// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type ActivityType int

// Valid ActivityType values.
const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom = 4
)
