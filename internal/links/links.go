package links

import (
	database "github.com/chienaeae/gql-todo/internal/pkg/db/migrations/mysql"
	"github.com/chienaeae/gql-todo/internal/users"
	"log"
)

// Link definition of  struct that represent a link
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.UserProfile
}

// Save function that insert a Link object into database and returns its ID
func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address, UserID) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("New Link Row inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("SELECT L.id, L.title, L.address, L.UserID, U.Username FROM Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var userID string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &userID, &username)
		if err != nil {
			log.Fatal(err)
		}
		link.User = &users.UserProfile{ID: userID, Username: username}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links

}
