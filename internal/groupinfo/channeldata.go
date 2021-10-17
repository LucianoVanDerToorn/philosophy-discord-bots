package groupinfo

import "fmt"

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
		ReadingWhat:    "King Lear by Shakespeare (October), Notes From Underground by Dostoevsky (November)",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
		ZoomLink:       "https://zoom.us/j/3564836909?pwd=eHN0N09MMmdxRXNkejQ3azVpVjFLdz09",
	},
	"camus": {
		Name:           "Camus",
		Leader:         "Salman",
		MeetingEvery:   "week",
		MeetingDay:     "Saturdays",
		MeetingTimeGmt: "3PM",
		ReadingWhat:    "The Myth of Sisyphus",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
		ZoomLink:       "https://us02web.zoom.us/j/86807037183",
	},
	// TODO: Add correct data and support
	//"film-discussions": {
	//	Name:           "Cinephile film discussions",
	//	Leader:         "Yorgo",
	//	MeetingEvery:   "month",
	//	MeetingDay:     "first Saturday of the Month",
	//	MeetingTimeGmt: "6.30PM",
	//	ReadingWhat:    "To Live (September)",
	//	ResourcesLink:  "https://www.youtube.com/watch?v=wNTd0dydfE4",
	//},
	"epistemology": {
		Name:           "Epistemology",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Thursday",
		MeetingTimeGmt: "7PM",
		ReadingWhat:    "changing every week",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
	},
	"marxism": {
		Name:           "Marxism",
		Leader:         "Chris",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "1.30PM",
		ReadingWhat:    "The German Ideology vol. 1",
		ResourcesLink:  "https://www.marxists.org/archive/marx/works/1845/german-ideology/ch01.htm",
		ZoomLink:       "https://zoom.us/j/97912504249?pwd=VmdvZ0pvNjcyamtVOGZCOEUyc1FQZz09",
	},
	"nietzsche": {
		Name:           "Nietzsche",
		Leader:         "Andrew",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "2PM",
		ReadingWhat:    "On the Use and Abuse of History",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		ZoomLink:       "https://huntercollege.zoom.us/j/4278613680",
	},
	"plato": {
		Name:           "Plato",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Saturday",
		MeetingTimeGmt: "1PM",
		ReadingWhat:    "The Republic",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		ZoomLink:       "https://us02web.zoom.us/j/87596642077?pwd=M2NPSVgxUUhzVVVRNUxFczFlb2cwZz09",
	},
	"stirner": {
		Name:           "Stirner",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Wednesday",
		MeetingTimeGmt: "7.30 GMT",
		ReadingWhat:    "The Ego and Its Own",
		ResourcesLink:  "https://www.dropbox.com/sh/gcmixj7nad7btsh/AAD0OEfua3SqWtDSAJnC3xnaa?dl=0",
		ZoomLink:       "https://us02web.zoom.us/j/87559118162",
	},
	"stoicism": {
		Name:           "Stoicism",
		Leader:         "Luciano",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "4PM",
		ReadingWhat:    "Epictetus' Enchiridion",
		ResourcesLink:  "https://www.dropbox.com/sh/8hehcb8oda7gc1k/AAC1YE5jwQ7VZK3_mEwkzbCDa?dl=0",
		ZoomLink:       "https://zoom.us/j/96810408257?pwd=VEdwMWpHQmxEaFpHNGNiL2l2Y0p4QT09",
	},
	"taoism": {
		Name:           "Taoism",
		Leader:         "Taolex",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "11PM",
		ReadingWhat:    "Chuang Tzu",
		ResourcesLink:  "https://terebess.hu/english/tao/ChuangTzu-palmer.pdf",
		ZoomLink:       "not available, we meet on Discord in the Reading Groups voice channel",
	},
}
