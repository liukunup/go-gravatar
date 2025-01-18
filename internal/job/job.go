package job

import (
	"go-gravatar/internal/repository"
	"go-gravatar/pkg/jwt"
	"go-gravatar/pkg/log"
	"go-gravatar/pkg/sid"
)

type Job struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT //nolint:golint,unused
	tm     repository.Transaction
}

func NewJob(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
) *Job {
	return &Job{
		logger: logger,
		sid:    sid,
		tm:     tm,
	}
}
