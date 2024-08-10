package crossfit

type MemberWod struct {
	ID       string `gorm:"primaryKey"`
	MemberId string
	WodId    string
}

func NewMemberWod(id, memberId, wodId string) (MemberWod, error) {
	return MemberWod{
		ID:       id,
		MemberId: memberId,
		WodId:    wodId,
	}, nil
}

type MemberWodsRepository interface {
	Save(memberWod MemberWod) error
	// Get(id string) ([]domain.CreatedWod, error)
}
