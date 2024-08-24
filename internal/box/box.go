package box

var admUsrPtfm = "0190e21a-adf8-75c5-9ed5-d60bce84686c"

type Box struct {
	ID        string `json:"id" gorm:"primary_key"`
	OwnerId   string `json:"owner_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Logo      string `json:"logo"`
	CreatedBy string `json:"created_by"`
}

func NewBox(id, owner_id, name, address, phone, logo string) (Box, error) {
	return Box{
		ID:        id,
		OwnerId:   owner_id,
		Name:      name,
		Address:   address,
		Phone:     phone,
		Logo:      logo,
		CreatedBy: admUsrPtfm,
	}, nil
}

type BoxRepository interface {
	Save(box Box) error
}
