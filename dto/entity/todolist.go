package entity

type Todolist struct {
	ID        string // random identity generated uniquely
	UserID    string // id user
	Title     string // title of todolist
	Desc      string // description of todolist
	Done      bool   // status od todolist
	DueDate   string // 18-09-2022 20:00:00 // deadline of todolist
	Category  string // category of todolist
	Priority  int    // 0=low, 1=med, 2=high
	CreatedAt int64  // database creation date
	UpdatedAt int64  // database modificatiion date

}
