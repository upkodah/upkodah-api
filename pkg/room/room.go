package room

import "github.com/upkodah/upkodah-api/pkg/db"

func SelectRooms(rs *[]Room) {
	db.Conn.Find(rs)
}

func (i *Info) First() error {
	if err := db.Conn.First(i).Error; err != nil {
		return err
	}
	return nil
}
