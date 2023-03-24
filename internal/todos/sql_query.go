package todos

import (
	"fmt"
	"strings"
)

var (
	getManyQuery         = "SELECT L.ID, L.Title, L.Description, L.UserID, U.Username FROM Todos L INNER JOIN Users U on U.ID = L.UserID"
	getByIDQuery         = "SELECT L.ID, L.Title, L.Description, L.UserID, U.Username FROM Todos L INNER JOIN Users U on U.ID = L.UserID WHERE L.ID = ?"
	getManyByUserIDQuery = "SELECT L.ID, L.Title, L.Description, L.UserID, U.Username FROM Todos L INNER JOIN Users U on U.ID = L.UserID WHERE L.UserID = ?"
	insertQuery          = "INSERT INTO Todos(Title, Description, UserID) VALUES (?, ?, ?)"
)

func newGetManyByUserIDWithLimit(opt *LimitOption) string {
	limitQuery := newLimitQuery(opt)
	queryString := fmt.Sprintf("%s %s", getManyByUserIDQuery, limitQuery)
	return queryString
}

func newLimitQuery(opt *LimitOption) string {
	if opt == nil {
		return ""
	} else {
		outStrings := make([]string, 0)
		if opt.AfterTodoID != nil {
			outStrings = append(outStrings, "AND L.ID > ?")
		}
		if opt.Limit != nil {
			outStrings = append(outStrings, "LIMIT ?")

		}
		return strings.Join(outStrings, " ")
	}
}

func appendLimitArgs(args []any, opt *LimitOption) []any {
	if args != nil && opt != nil {
		if opt.AfterTodoID != nil {
			args = append(args, *opt.AfterTodoID)
		}
		if opt.Limit != nil {
			args = append(args, *opt.Limit)
		}
	}
	return args
}
