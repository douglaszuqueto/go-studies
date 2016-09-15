package Applications

import "time"

type User struct {
	Id        int       `json:"id`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Users []User
