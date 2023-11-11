package main

import "time"

type Task struct {
	ID        int
	Title     string
	CreatedAt time.Time
}
