package taskrunner

import (
	"errors"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// Определение тестовой задачи, которая завершается с ошибкой
	errTask := func() error {
		return errors.New("error")
	}

	// Определение тестовой задачи, которая завершается успешно
	successTask := func() error {
		time.Sleep(50 * time.Millisecond) // Имитация работы
		return nil
	}

	tests := []struct {
		name    string
		tasks   []Task
		n       int
		m       int
		wantErr bool
	}{
		{
			name:    "No tasks",
			tasks:   []Task{},
			n:       5,
			m:       2,
			wantErr: false,
		},
		{
			name:    "Tasks without errors",
			tasks:   []Task{successTask, successTask, successTask},
			n:       2,
			m:       1,
			wantErr: false,
		},
		{
			name:    "Tasks with errors below limit",
			tasks:   []Task{errTask, successTask, errTask},
			n:       3,
			m:       3,
			wantErr: false,
		},
		{
			name:    "Tasks with errors reached limit",
			tasks:   []Task{errTask, errTask, successTask},
			n:       2,
			m:       2,
			wantErr: true,
		},
		{
			name:    "Ignore errors",
			tasks:   []Task{errTask, errTask, errTask},
			n:       1,
			m:       0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.tasks, tt.n, tt.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err != ErrErrorsLimitExceeded {
				t.Errorf("Run() error = %v, want %v", err, ErrErrorsLimitExceeded)
			}
		})
	}
}
