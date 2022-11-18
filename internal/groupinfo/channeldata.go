package groupinfo

import "fmt"

type MeetingOn string

const (
	MeetingOnUnsure  = "unsure"
	MeetingOnZoom    = "zoom"
	MeetingOnDiscord = "discord"
)

type ChannelData struct {
	ChannelId      string
	Name           string
	Leader         string
	MeetingEvery   string
	MeetingDay     string
	MeetingTimeGmt string
	ReadingWhat    string
	ResourcesLink  string
	ZoomLink       string
	MeetingOn      MeetingOn
}

func (cd ChannelData) EmbedDescription() string {
	template := `**Info about the *%s* reading group**
We meet every %s on **%s** at **%s** GMT
We are currently reading **%s**
Press the title to open the link with the books
The reading group leader is %s for any further questions`
	return fmt.Sprintf(template, cd.Name, cd.MeetingEvery, cd.MeetingDay, cd.MeetingTimeGmt, cd.ReadingWhat, cd.Leader)
}

var ChannelDataLookup = map[string]ChannelData{
	"bookworms": {
		Name:           "Bookworms",
		Leader:         "Luciano",
		MeetingEvery:   "month",
		MeetingDay:     "last Saturday of the Month",
		MeetingTimeGmt: "6.30PM",
		ReadingWhat:    "See channel",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
		ZoomLink:       "https://zoom.us/j/3564836909?pwd=eHN0N09MMmdxRXNkejQ3azVpVjFLdz09",
		MeetingOn:      MeetingOnZoom,
	},
	"nietzsche": {
		Name:           "Nietzsche",
		Leader:         "Andrew",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "2PM",
		ReadingWhat:    "Beyond Good And Evil",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		MeetingOn:      MeetingOnDiscord,
	},
	"plato": {
		Name:           "Plato",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Saturday",
		MeetingTimeGmt: "1PM",
		ReadingWhat:    "The Republic, book 7",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		MeetingOn:      MeetingOnUnsure,
	},
	"stirner": {
		Name:           "Stirner",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Wednesday",
		MeetingTimeGmt: "7.30 GMT",
		ReadingWhat:    "The Unique and Its Property",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		MeetingOn:      MeetingOnDiscord,
	},
	"stoicism": {
		Name:           "Stoicism",
		Leader:         "Luciano",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "4PM",
		ReadingWhat:    "Marcus Aurelius' Meditations",
		ResourcesLink:  "https://www.dropbox.com/sh/8hehcb8oda7gc1k/AAC1YE5jwQ7VZK3_mEwkzbCDa?dl=0",
		ZoomLink:       "https://us06web.zoom.us/j/97681357129?pwd=VUs3N1FuV2Z4ckx3UDJRUDhJVE94Zz09",
		MeetingOn:      MeetingOnZoom,
	},
	"taoism": {
		Name:           "Taoism",
		Leader:         "Taolex",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "11PM",
		ReadingWhat:    "Chuang Tzu",
		ResourcesLink:  "https://terebess.hu/english/tao/ChuangTzu-palmer.pdf",
		MeetingOn:      MeetingOnDiscord,
	},
}
