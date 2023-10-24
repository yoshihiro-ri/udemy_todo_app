package models

import (
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values(?,?,?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	return hundleError(err)
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id,created_at from todos where id =?`
	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo)

	return todo, err

}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id,content,user_id,created_at from todos`
	rows, err := Db.Query(cmd)
	hundleError(err)
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		hundleError(err)
		todos = append(todos, todo)
	}

	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser(user_id int) (todos []Todo, err error) {
	cmd := `select id,content,user_id,created_at from todos where user_id = ?`
	row, err := Db.Query(cmd, user_id)
	hundleError(err)
	for row.Next() {
		var todo Todo
		err = row.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		hundleError(err)
		todos = append(todos, todo)
	}
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set content = ?,user_id = ?
	where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	return hundleError(err)
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id=?`
	_, err = Db.Exec(cmd, t.ID)
	return hundleError(err)
}
