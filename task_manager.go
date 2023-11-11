package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type TaskManager struct {
	Tasks []Task
}

func (tm *TaskManager) AddTask(title string) {
	task := Task{
		ID:        len(tm.Tasks) + 1,
		Title:     title,
		CreatedAt: time.Now(),
	}
	tm.Tasks = append(tm.Tasks, task)
}

func (tm *TaskManager) RemoveTask(taskID int) {
	for i, task := range tm.Tasks {
		if task.ID == taskID {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			break
		}
	}
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}

func (tm *TaskManager) loadTasksFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tm.Tasks)

	if err != nil {
		return err
	}

	return nil
}

func (tm *TaskManager) saveTasksToFile(filename string) error {
	data, err := json.MarshalIndent(tm.Tasks, "", "  ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}

	return nil
}
