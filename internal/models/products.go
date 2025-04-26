package models

type Product struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
    Name       string `json:"name"`
    Price      string `json:"price"`
    ImageURL   string `json:"image_url"`
    Brand      string `json:"brand"`
    CategoryID uint   `json:"category_id"`
}

type Category struct {
    ID         uint   `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Img      string    `json:"img"`
	Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
// CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
type Catalog struct {
	Categories []Category `json:"categories"`
}