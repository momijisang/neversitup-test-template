package persistence

import (
	"neversitup-test-template/internal/pkg/models"
)

type LogRepository struct{}

var logRepository *LogRepository

func Log() *LogRepository {
	if logRepository == nil {
		logRepository = &LogRepository{}
	}
	return logRepository
}

func (r *LogRepository) AddApiLog(URL string, Status int, Params string, Response string, IsSuccess bool) {
	log := models.ApiLog{
		URL:          URL,
		RequestData:  Params,
		Status:       Status,
		ResponseData: Response,
		IsSuccess:    IsSuccess,
		CreatedAt:    models.IncNowTime(),
	}

	log.Insert()
}

func (r *LogRepository) AddLog(functionName string, err error, param ...interface{}) {
	log := models.Log{
		FunctionName: functionName,
		Param:        param,
		IsError:      err != nil,
		Error:        err,
		CreatedAt:    models.IncNowTime(),
	}
	log.Insert()
}
