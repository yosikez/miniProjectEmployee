package inputs

type Gender string

const (
    Male   Gender = "male"
    Female Gender = "female"
)

type EmployeeInput struct {
	Name    string `form:"name" json:"name" binding:"required" gorm:"column:name"`
	Email   string `form:"email" json:"email" binding:"required,email" gorm:"column:email"`
	Address string `form:"address" json:"address" binding:"required" gorm:"column:address"`
	Phone   string `form:"phone" json:"phone" binding:"required,max=14" gorm:"column:phone"`
	Gender  Gender `form:"gender" json:"gender" binding:"required,gender" gorm:"column:gender"`
}
