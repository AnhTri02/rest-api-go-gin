package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendees struct {
	Id       int `json:"id"`
	User_Id  int `json:"user_id"`
	Event_id int `json:"event_id"`
	
}
