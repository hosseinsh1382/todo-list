package Models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primary_key;autoincrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	CreatedAt   time.Time `json:"created_at"`
	Deleted     bool      `json:"deleted"`
}

func NewTask(title string, description string, startAt time.Time, endAt time.Time) *Task {
	t := Task{
		Title:       title,
		Description: description,
		StartAt:     startAt,
		EndAt:       endAt,
		CreatedAt:   time.Now(),
		Deleted:     false,
	}
	return &t
}

/*func (t *Task) Id() int {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}
func (t *Task) Description() string {
	return t.description
}
func (t *Task) StartAt() time.Time {
	return t.startAt
}
func (t *Task) EndAt() time.Time {
	return t.endAt
}
func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) SetTitle(title string) {
	t.title = title
}
func (t *Task) SetDescription(description string) {
	t.description = description
}
func (t *Task) SetStartAt(startAt time.Time) {
	t.startAt = startAt
}
func (t *Task) SetEndAt(endAt time.Time) {
	t.endAt = endAt
}
*/
