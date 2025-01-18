package task

import (
	"go-gravatar/internal/repository"
	"go-gravatar/pkg/jwt"
	"go-gravatar/pkg/log"
	"go-gravatar/pkg/sid"
)

type Task struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT //nolint:golint,unused
	tm     repository.Transaction
}

func NewTask(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Task {
	return &Task{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
