package tags

import (
	database "github.com/chienaeae/gql-todo/internal/pkg/db/migrations/mysql"
	"log"
)

type Tag struct {
	ID          string
	Name        string
	Description string
}

func (t Tag) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Tags(Name, Description) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(t.Name, t.Description)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("New Tag Row inserted!")
	return id
}

func GetAll() []Tag {
	stmt, err := database.Db.Prepare("SELECT ID, Name, Description FROM Tags")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name, &tag.Description)
		if err != nil {
			log.Fatal(err)
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tags
}
