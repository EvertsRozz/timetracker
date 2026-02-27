package project

import (
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
	Minutes uint
	Note    string
	Date    time.Time
}

func (p *Project) AddLog(minutes uint, note string) {
	p.TotalTime += minutes

	p.Logs = append(p.Logs, Log{
		Minutes: minutes,
		Note:    strings.TrimSpace(note),
		Date:    time.Now(),
	})
}
