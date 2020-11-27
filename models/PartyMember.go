package models

// PartyMember 活动成员
type PartyMember struct {
	ID            int
	PartyID       int
	MemberID      int
	JoinPeopleNum int
	CanJoinDate   string
	CreatedAt     string
	UpdatedAt     string
}
