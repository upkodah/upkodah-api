package facility

import "github.com/upkodah/upkodah-api/pkg/db"

func SelectAll(fs *[]Facility) {
	db.Conn.Find(fs)
}
