package facility

import "github.com/jinzhu/gorm"

// stringify Facility Array
type Facility struct {
	gorm.Model
	Name string
}
