package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	Tasks  []Task
	NextID int
}

func NewStore() *Store {
	return &Store{
		Tasks:  []Task{},
		NextID: 1,
	}
}

// SetTasks permet au main de remplir le store après chargement
func (s *Store) SetTasks(tasks []Task) {
	s.Tasks = tasks
	for _, t := range tasks {
		if t.ID >= s.NextID {
			s.NextID = t.ID + 1
		}
	}
}

func (s *Store) CreateTask(title string, desc string) Task {
	newTask := Task{
		ID:        s.NextID,
		Title:     title,
		Desc:      desc,
		CreatedAt: time.Now(),
	}
	s.Tasks = append(s.Tasks, newTask)
	s.NextID++
	return newTask
}

func (s *Store) Show() {
	if len(s.Tasks) == 0 {
		fmt.Println("Aucune tâche enregistrée.")
		return
	}
	fmt.Println("ID  | Titre           | Description          | Date")
	fmt.Println("-----------------------------------------------------------")
	for _, t := range s.Tasks {
		fmt.Printf("%-3d | %-15s | %-20s | %s\n",
			t.ID, t.Title, t.Desc, t.CreatedAt.Format("02/01/2006 15:04"))
	}
}

func (s *Store) UpdateTask(id int, title string, desc string) error {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			s.Tasks[i].Title = title
			s.Tasks[i].Desc = desc
			return nil
		}
	}
	return fmt.Errorf("tâche %d introuvable", id)
}

func (s *Store) DeleteTask(id int) error {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("tâche %d introuvable", id)
}
