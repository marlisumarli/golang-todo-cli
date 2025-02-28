package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	ID             int
	Task, Category string
	Done           bool
	CreatedAt      time.Time
	CompletedAt    *time.Time
}

type Todos []item

var nexID int

func (t *Todos) Add(task, category string) {
	todo := item{
		ID:          nexID,
		Task:        task,
		Category:    category,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	nexID++

	*t = append(*t, todo)
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 {
		return err
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	if len(*t) > 0 {
		maxID := (*t)[0].ID
		for _, todo := range *t {
			if todo.ID > maxID {
				maxID = todo.ID
			}
		}
		nexID = maxID + 1
	}

	return nil
}

func (t *Todos) Print(status int, category string) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Category"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	requestedTodos := []item{}

	for _, todo := range *t {
		if status == 1 {
			if todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status == 0 {
			if !todo.Done {
				requestedTodos = append(requestedTodos, todo)
			}
		}

		if status != 1 && status != 0 {
			requestedTodos = append(requestedTodos, todo)
		}
	}

	requestedCategory := []item{}

	for _, todo := range requestedTodos {
		// if strings.ToLower(todo.Category) == strings.ToLower(category) || category == "" {

		// Simplenya
		if strings.EqualFold(todo.Category, category) || category == "" {
			requestedCategory = append(requestedCategory, todo)
		}
	}

	for _, item := range requestedCategory {
		task := item.Task
		done := "no"
		completedAt := ""

		if item.Done {
			task = fmt.Sprint(item.Task)
			done = "\u2705"
		}

		if item.CompletedAt != nil {
			completedAt = item.CompletedAt.Format("2006-01-02")
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprint(item.ID)},
			{Text: item.Category},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format("2006-01-02")},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: ""},
		{Align: simpletable.AlignLeft, Span: 5, Text: fmt.Sprintf("Ada %d pending todo bro.", t.CountPending())},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int {
	total := 0
	for _, item := range *t {
		if !item.Done {
			total++
		}
	}
	return total
}

func (t *Todos) Delete(id int) error {
	ls := *t
	index := t.getIndexById(id)

	if index == -1 {
		return errors.New("invalid ID")
	}

	*t = append(ls[:index], ls[index+1:]...)

	return nil
}

func (t *Todos) DeleteAll() error {
	*t = []item{}
	return nil
}

func (t *Todos) getIndexById(id int) int {
	index := -1
	for i, todo := range *t {
		if todo.ID == id {
			index = i
			break
		}
	}

	return index
}

func (t *Todos) Update(id int, task string, cat string, done int) error {
	ls := *t

	index := t.getIndexById(id)
	if index == -1 {
		return errors.New("invalid ID")
	}

	if len(task) != 0 {
		ls[index].Task = task
	}

	if len(cat) != 0 {
		ls[index].Category = cat
	}

	if done == 0 {
		ls[index].Done = true
		ls[index].CompletedAt = nil
	} else if done == 1 {
		ls[index].Done = false
		completedAt := time.Now()
		ls[index].CompletedAt = &completedAt
	}

	return nil
}
