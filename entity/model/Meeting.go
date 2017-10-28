package model

type Meeting struct {
	Title string
	Sponsor string
	Participators string
	StartDate string
	EndDate string
}

func (meeting Meeting) GetTitle() string {
	return meeting.Title
}

func (meeting *Meeting) SetTitle(title string) {
	meeting.Title = title
}

func (meeting Meeting) GetSponsor() string {
	return meeting.Sponsor
}

func (meeting *Meeting) SetSponsor(sponsor string) {
	meeting.Sponsor = sponsor
}

func (meeting Meeting) GetParticipators() string {
	return meeting.Participators
}

func (meeting *Meeting) SetParticipators(participators string) {
	meeting.Participators = participators
}

func (meeting Meeting) GetStartDate() string {
	return meeting.StartDate
}

func (meeting *Meeting) Set(startDate string) {
	meeting.StartDate = startDate
}

func (meeting Meeting) GetEndDate() string {
	return meeting.EndDate
}

func (meeting *Meeting) SetEndDate(endDate string) {
	meeting.EndDate = endDate
}
