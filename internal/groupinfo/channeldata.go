package groupinfo

import "fmt"

type MeetingOn string

const (
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
	"aristotle": {
		Name:           "Aristotle",
		Leader:         "Humphrey",
		MeetingEvery:   "week",
		MeetingDay:     "Saturday",
		MeetingTimeGmt: "6PM",
		ReadingWhat:    "Nicomachean Ethics",
		ResourcesLink:  "https://www.dropbox.com/sh/sgu71nyo4uesy0g/AAAW8R7edfH2z8-k9JfAKQZoa?dl=0",
		ZoomLink:       "https://zoom.us/j/97316252427?pwd=MmNUQUdDc0w3R0hxOWxPRGdtZ0g4Zz09 ",
		MeetingOn:      MeetingOnZoom,
	},
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
	"existentialism": {
		Name:           "Existentialism",
		Leader:         "Salman",
		MeetingEvery:   "week",
		MeetingDay:     "Saturdays",
		MeetingTimeGmt: "4PM",
		ReadingWhat:    "Fear and Trembling (Kierkegaard)",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
		ZoomLink:       "https://us02web.zoom.us/j/86807037183",
		MeetingOn:      MeetingOnZoom,
	},
	"ontology": {
		Name:           "Ontology",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Thursday",
		MeetingTimeGmt: "8PM",
		ReadingWhat:    "changing every week",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		MeetingOn:      MeetingOnZoom,
	},
	"marxism": {
		Name:           "Marxism",
		Leader:         "Marta",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "7.30PM",
		ReadingWhat:    "The Communist Manifesto",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		ZoomLink:       "https://zoom.us/j/97912504249?pwd=VmdvZ0pvNjcyamtVOGZCOEUyc1FQZz09",
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
		ZoomLink:       "https://us02web.zoom.us/j/87596642077?pwd=M2NPSVgxUUhzVVVRNUxFczFlb2cwZz09",
		MeetingOn:      MeetingOnZoom,
	},
	"stirner": {
		Name:           "Stirner",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Wednesday",
		MeetingTimeGmt: "7.30 GMT",
		ReadingWhat:    "The Unique and Its Property",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		ZoomLink:       "https://us02web.zoom.us/j/87559118162",
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
		ZoomLink:       "https://zoom.us/j/96810408257?pwd=VEdwMWpHQmxEaFpHNGNiL2l2Y0p4QT09",
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
