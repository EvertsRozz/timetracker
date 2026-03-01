package project

import (
	"errors"
	"strings"
	"time"
)

type Project struct {
	Name      string  `json:"name"`
	Wage      float32 `json:"wage"`
	TotalTime uint    `json:"total_time"`
	Logs      []Log   `json:"logs"`
}

type Log struct {
	Minutes uint16
	Note    string
	Date    time.Time
}

func (p *Project) AddLog(minutes uint16, note string) {
	p.TotalTime += uint(minutes)

	p.Logs = append(p.Logs, Log{
		Minutes: minutes,
		Note:    strings.TrimSpace(note),
		Date:    time.Now(),
	})
}

func NewProject(name string, wage float32) (*Project, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("project name must be specified")
	}
	if wage < 0 {
		return nil, errors.New("wage must be non-negative")
	}

	return &Project{
		Name:      name,
		Wage:      wage,
		TotalTime: 0,
		Logs:      nil,
	}, nil
}
