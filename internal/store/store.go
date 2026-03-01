package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/EvertsRozz/timetracker/project"
)

type Store struct {
	Projects []project.Project `json:"projects"`
}

func Load() (*Store, error) {
	data, err := os.ReadFile("tracker.json")
	if err != nil {
		return &Store{}, nil // Empty on first run
	}
	var s Store
	json.Unmarshal(data, &s)
	return &s, nil
}

func (s *Store) Save() error {
	data, _ := json.MarshalIndent(s, "", "  ")
	return os.WriteFile("tracker.json", data, 0644)
}

func (s *Store) Find(name string, create bool) (*project.Project, error) {
	for i := range s.Projects {
		if s.Projects[i].Name == name {
			return &s.Projects[i], nil
		}
	}

	return nil, errors.New("Project doesn't exist")
}

func (s *Store) Create(name string, wage float32) (*project.Project, error) {
	for i := range s.Projects {
		if s.Projects[i].Name == name {
			return nil, errors.New("Project with name already exists")
		}
	}

	newProj, err := project.NewProject(name, wage)
	if err != nil {
		return nil, err
	}
	s.Projects = append(s.Projects, *newProj)
	return newProj, nil
}
