package members

var admUsrPtfm = "0190e21a-adf8-75c5-9ed5-d60bce84686c"

type Members struct {
	ID             string `json:"id" gorm:"primary_key"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email" gorm:"unique"`
	Age            uint8  `json:"age"`
	Gender         string `json:"gender"`
	City           string `json:"city"`
	Country        string `json:"country"`
	ProfilePicture string `json:"profile_picture"`
	FrontPageImg   string `json:"front_page_img"`
	CreatedBy      string `json:"created_by"`
}

func NewMember(id, firstName, LastName, email, gender, city, country, profilePicture, frontPageImg string, age uint8) (Members, error) {
	return Members{
		ID:             id,
		FirstName:      firstName,
		LastName:       LastName,
		Email:          email,
		Age:            age,
		Gender:         gender,
		City:           city,
		Country:        country,
		ProfilePicture: profilePicture,
		FrontPageImg:   frontPageImg,
		CreatedBy:      admUsrPtfm,
	}, nil
}

type MemberRepository interface {
	Save(member Members) error
	Find(id string) (Members, error)
	FindByEmail(email string) (Members, error)
	Update(member Members) error
}
