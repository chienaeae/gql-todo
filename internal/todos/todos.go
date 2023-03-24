package todos

import (
	"database/sql"
	"fmt"
	database "github.com/chienaeae/gql-todo/internal/pkg/db/migrations/mysql"
	"github.com/chienaeae/gql-todo/internal/users"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	User        *users.UserProfile
}

type LimitOption struct {
	AfterTodoID *string
	Limit       *int
}

func (t Todo) Save() int64 {
	stmt, err := database.Db.Prepare(insertQuery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(t.Title, t.Description, t.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	log.Print("New Todo Row inserted!")
	return id
}

func (t Todo) LinkTag(tagId string) error {
	stmt, err := database.Db.Prepare("INSERT INTO TodoTagLinks(TagID, TodoID) VALUE (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(tagId, t.ID)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			fmt.Println(mysqlErr)
			switch mysqlErr.Number {
			case 1452:
				return LinkToInvalidTagIDError
			}
		}
		log.Fatal(err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("New TodoTagLink Row inserted!")
	return nil
}

func GetAll() []Todo {
	stmt, err := database.Db.Prepare(getManyQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return newTodosFromRows(rows)
}

func GetById(todoID string) (Todo, error) {
	stmt, err := database.Db.Prepare(getByIDQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(todoID)
	if err := row.Err(); err != nil {
		log.Fatal(err)
	}
	var todo Todo
	var username string
	var userID string
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &username, &userID)
	todo.User = &users.UserProfile{
		ID:       userID,
		Username: username,
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, &TodoIdNotExistsError{}
		}
		log.Fatal(err)
	}
	return todo, nil
}

func GetManyByUserID(uID string, opt *LimitOption) []Todo {
	query := newGetManyByUserIDWithLimit(opt)
	stmt, err := database.Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var rows *sql.Rows
	if opt == nil {
		rows, err = stmt.Query(uID)
	} else {
		args := []any{
			uID,
		}
		args = appendLimitArgs(args, opt)
		rows, err = stmt.Query(args...)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return newTodosFromRows(rows)
}

func newTodosFromRows(rows *sql.Rows) []Todo {
	var todos []Todo
	var username string
	var userID string
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &userID, &username)
		if err != nil {
			log.Fatal(err)
		}

		todo.User = &users.UserProfile{
			ID:       userID,
			Username: username,
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return todos
}
