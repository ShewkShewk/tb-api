package tbapi

import "time"

type Tournament struct {
	Id   int
	Date time.Time
	Name string
}

type TournamentData struct {
	Timezone   string `json:"timezone"`
	Categories []struct {
		Id     string `json:"id"`
		Judges []struct {
			Id       string `json:"id"`
			Last     string `json:"last"`
			Email    string `json:"email"`
			Settings []struct {
				Value *string `json:"value"`
				Tag   string  `json:"tag"`
				Meta  string  `json:"meta,omitempty"`
			} `json:"settings,omitempty"`
			Person       int    `json:"person"`
			Phone        string `json:"phone,omitempty"`
			First        string `json:"first"`
			Active       int    `json:"active"`
			Covers       int    `json:"covers,omitempty"`
			Ada          int    `json:"ada"`
			School       string `json:"school,omitempty"`
			ChapterJudge int    `json:"chapter_judge,omitempty"`
			Strikes      []struct {
				Conflict   int    `json:"conflict"`
				Conflictee int    `json:"conflictee"`
				Tag        string `json:"tag"`
				Registrant int    `json:"registrant"`
				Id         int    `json:"id"`
				School     int    `json:"school,omitempty"`
				End        string `json:"end,omitempty"`
				Start      string `json:"start,omitempty"`
			} `json:"strikes,omitempty"`
		} `json:"judges"`
		Abbr     string `json:"abbr"`
		Settings []struct {
			Tag   string `json:"tag"`
			Value string `json:"value"`
			Meta  string `json:"meta"`
		} `json:"settings"`
		Name   string `json:"name"`
		Events []struct {
			Settings []struct {
				Value string `json:"value"`
				Tag   string `json:"tag"`
				Meta  string `json:"meta"`
			} `json:"settings"`
			Fee        string `json:"fee"`
			Id         string `json:"id"`
			ResultSets []struct {
				Bracket int `json:"bracket"`
				Results []struct {
					Rank       int    `json:"rank,omitempty"`
					Percentile string `json:"percentile,omitempty"`
					Values     []struct {
						Value     string `json:"value,omitempty"`
						Priority  int    `json:"priority,omitempty"`
						ResultKey string `json:"result_key,omitempty"`
						Protocol  int    `json:"protocol,omitempty"`
					} `json:"values"`
					Entry   int    `json:"entry,omitempty"`
					Place   string `json:"place,omitempty"`
					Student int    `json:"student,omitempty"`
					Round   int    `json:"round,omitempty"`
				} `json:"results"`
				Coach      interface{} `json:"coach"`
				Generated  string      `json:"generated"`
				ResultKeys []struct {
					Id          *string `json:"id"`
					SortDesc    *int    `json:"sort_desc"`
					Tag         *string `json:"tag"`
					NoSort      *int    `json:"no_sort"`
					Description *string `json:"description"`
				} `json:"result_keys,omitempty"`
				Label     string `json:"label"`
				Published int    `json:"published"`
			} `json:"result_sets"`
			Rounds []struct {
				Published    int    `json:"published"`
				Protocol     int    `json:"protocol"`
				StartTime    string `json:"start_time"`
				Site         int    `json:"site"`
				Timeslot     int    `json:"timeslot"`
				ProtocolName string `json:"protocol_name"`
				Sections     []struct {
					Ballots []struct {
						Side   int `json:"side"`
						Audit  int `json:"audit"`
						Scores []struct {
							Id       int     `json:"id"`
							Content  string  `json:"content,omitempty"`
							Value    float64 `json:"value"`
							Tag      string  `json:"tag"`
							Speaker  int     `json:"speaker,omitempty"`
							Position int     `json:"position,omitempty"`
						} `json:"scores,omitempty"`
						StartedBy    int    `json:"started_by,omitempty"`
						Id           string `json:"id"`
						Chair        int    `json:"chair"`
						JudgeStarted string `json:"judge_started,omitempty"`
						Judge        int    `json:"judge,omitempty"`
						Panel        string `json:"panel"`
						Entry        int    `json:"entry"`
						Bye          int    `json:"bye,omitempty"`
						Forfeit      int    `json:"forfeit,omitempty"`
					} `json:"ballots,omitempty"`
					Room     *int          `json:"room"`
					Flight   string        `json:"flight"`
					Id       string        `json:"id"`
					Round    string        `json:"round"`
					Letter   string        `json:"letter"`
					Settings []interface{} `json:"settings"`
					Bye      int           `json:"bye,omitempty"`
					Bracket  int           `json:"bracket,omitempty"`
				} `json:"sections"`
				Name          int    `json:"name"`
				Type          string `json:"type"`
				Flights       int    `json:"flights"`
				PostFeedback  int    `json:"post_feedback"`
				PairedAt      string `json:"paired_at"`
				PostSecondary int    `json:"post_secondary"`
				Settings      []struct {
					Meta  string `json:"meta"`
					Value string `json:"value"`
					Tag   string `json:"tag"`
				} `json:"settings"`
				PostPrimary int    `json:"post_primary"`
				Id          string `json:"id"`
				Label       string `json:"label,omitempty"`
			} `json:"rounds"`
			Name string `json:"name"`
			Abbr string `json:"abbr"`
			Type string `json:"type"`
		} `json:"events"`
	} `json:"categories"`
	City          string `json:"city"`
	Creator       string `json:"creator"`
	BackupCreated string `json:"backup_created"`
	Schools       []struct {
		Chapter int    `json:"chapter"`
		Code    string `json:"code"`
		Entries []struct {
			Active   int         `json:"active"`
			Ada      interface{} `json:"ada"`
			Id       string      `json:"id"`
			Hybrid   interface{} `json:"hybrid"`
			Settings []struct {
				Meta  string `json:"meta"`
				Tag   string `json:"tag"`
				Value string `json:"value"`
			} `json:"settings"`
			Waitlist     int      `json:"waitlist"`
			Event        int      `json:"event"`
			Unconfirmed  int      `json:"unconfirmed"`
			Code         string   `json:"code"`
			Dropped      int      `json:"dropped"`
			CreatedAt    string   `json:"created_at"`
			RegisteredBy int      `json:"registered_by"`
			School       string   `json:"school"`
			Students     []string `json:"students,omitempty"`
			Name         *string  `json:"name"`
		} `json:"entries"`
		Onsite   int `json:"onsite,omitempty"`
		Settings []struct {
			Value string `json:"value"`
			Tag   string `json:"tag"`
			Meta  string `json:"meta"`
		} `json:"settings"`
		Id       string `json:"id"`
		Nsda     int    `json:"nsda"`
		Name     string `json:"name"`
		Students []struct {
			First    string `json:"first"`
			Nsda     int    `json:"nsda,omitempty"`
			GradYear int    `json:"grad_year"`
			Chapter  int    `json:"chapter"`
			Last     string `json:"last"`
			Id       string `json:"id"`
			Middle   string `json:"middle,omitempty"`
		} `json:"students"`
	} `json:"schools"`
	Country   string `json:"country"`
	Name      string `json:"name"`
	End       string `json:"end"`
	RoomPools []struct {
		Id     string  `json:"id"`
		Rounds []int   `json:"rounds,omitempty"`
		Rooms  []int   `json:"rooms"`
		Name   *string `json:"name"`
	} `json:"room_pools"`
	RegStart  string `json:"reg_start"`
	Hidden    int    `json:"hidden"`
	Timeslots []struct {
		Start string `json:"start"`
		Id    string `json:"id"`
		End   string `json:"end"`
		Name  string `json:"name"`
	} `json:"timeslots"`
	Webname   string `json:"webname"`
	Protocols []struct {
		Tiebreaks []struct {
			Count            string `json:"count"`
			Name             string `json:"name"`
			Priority         int    `json:"priority"`
			Multiplier       int    `json:"multiplier"`
			Highlow          int    `json:"highlow,omitempty"`
			HighlowCount     int    `json:"highlow_count,omitempty"`
			TruncateSmallest int    `json:"truncate_smallest,omitempty"`
			Truncate         int    `json:"truncate,omitempty"`
		} `json:"tiebreaks"`
		Name     string `json:"name"`
		Id       string `json:"id"`
		Settings []struct {
			Value string `json:"value"`
			Tag   string `json:"tag"`
			Meta  string `json:"meta"`
		} `json:"settings"`
	} `json:"protocols"`
	Whole    int    `json:"whole"`
	Start    string `json:"start"`
	RegEnd   string `json:"reg_end"`
	Webpages []struct {
		PageOrder  int         `json:"page_order"`
		LastEditor int         `json:"last_editor"`
		Special    interface{} `json:"special"`
		Id         int         `json:"id"`
		Content    string      `json:"content"`
		Title      string      `json:"title"`
		Parent     interface{} `json:"parent"`
		Published  int         `json:"published"`
	} `json:"webpages"`
	Sites []struct {
		Name  string `json:"name"`
		Rooms []struct {
			Name     string `json:"name"`
			Id       string `json:"id"`
			Inactive int    `json:"inactive,omitempty"`
		} `json:"rooms"`
		Id string `json:"id"`
	} `json:"sites"`
	State    string `json:"state"`
	Settings []struct {
		Value string `json:"value"`
		Tag   string `json:"tag"`
		Meta  string `json:"meta"`
	} `json:"settings"`
	Emails []struct {
		SentTo   string `json:"sent_to"`
		Sender   int    `json:"sender"`
		Content  string `json:"content"`
		SentAt   string `json:"sent_at"`
		Subject  string `json:"subject"`
		Metadata string `json:"metadata"`
		Id       int    `json:"id"`
	} `json:"emails"`
}
