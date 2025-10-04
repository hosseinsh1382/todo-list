package Models

import (
	"time"
)

type Task struct {
	id          int       `json:"id"`
	title       string    `json:"title"`
	description string    `json:"description"`
	startAt     time.Time `json:"start_at"`
	endAt       time.Time `json:"end_at"`
	createdAt   time.Time `json:"created_at"`
}

func NewTask(title string, description string, startAt time.Time, endAt time.Time) *Task {
	t := Task{
		title:       title,
		description: description,
		startAt:     startAt,
		endAt:       endAt,
		createdAt:   time.Now(),
	}
	return &t
}

func (t *Task) Id() int {
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
