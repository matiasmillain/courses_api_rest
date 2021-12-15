package entity

//course struct represents courses table in database
type Course struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserOwner   uint64 `gorm:"not null" json:"-"`
}
