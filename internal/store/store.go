package store

import (
	"encoding/json"
	"fmt"
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

func (s *Store) FindOrCreate(name string, create bool) (*project.Project, error) {
	for i := range s.Projects {
		if s.Projects[i].Name == name {
			return &s.Projects[i], nil
		}
	}

	fmt.Printf("❌ Project '%s' not found. Use -c to create.\n", name)

	// proj := &project.Project{Name: name} // Default wage=0
	// s.Projects = append(s.Projects, *proj)
	// fmt.Printf("Created new project '%s'\n Configure it later using config", name)
	// return proj
}
