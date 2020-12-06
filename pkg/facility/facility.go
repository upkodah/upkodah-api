package facility

import "github.com/upkodah/upkodah-api/pkg/db"

func SelectAll(fs *[]Facility) error {
	return db.Conn.Find(fs).Error
}
