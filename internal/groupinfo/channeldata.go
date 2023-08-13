package groupinfo

type MeetingOn string

const (
	MeetingOnUnsure  = "unsure"
	MeetingOnZoom    = "zoom"
	MeetingOnDiscord = "discord"
)

type ChannelData struct {
	ChannelId string
	Name      string
	Leader    string
	ZoomLink  string
	MeetingOn MeetingOn
}

var ChannelDataLookup = map[string]ChannelData{
	"bookworms": {
		Name:      "Bookworms",
		Leader:    "Luciano",
		ZoomLink:  "https://zoom.us/j/3564836909?pwd=eHN0N09MMmdxRXNkejQ3azVpVjFLdz09",
		MeetingOn: MeetingOnZoom,
	},
	"nietzsche": {
		Name:      "Nietzsche",
		Leader:    "Andrew",
		MeetingOn: MeetingOnDiscord,
	},
	"plato": {
		Name:      "Plato",
		Leader:    "Constantine",
		MeetingOn: MeetingOnUnsure,
	},
	"stirner": {
		Name:      "Stirner",
		Leader:    "Timothy",
		ZoomLink:  "https://weber.zoom.us/j/96400307602",
		MeetingOn: MeetingOnZoom,
	},
	"stoicism": {
		Name:      "Stoicism",
		Leader:    "Luciano",
		ZoomLink:  "https://us06web.zoom.us/j/97681357129?pwd=VUs3N1FuV2Z4ckx3UDJRUDhJVE94Zz09",
		MeetingOn: MeetingOnZoom,
	},
	"taoism": {
		Name:      "Taoism",
		Leader:    "Taolex",
		MeetingOn: MeetingOnDiscord,
	},
}
