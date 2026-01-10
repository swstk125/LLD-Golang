package main

import (
	"time"
)

/*
Design a Logger system for a backend application such that:
	-There must be only one Logger instance throughout the application.
	-Logger should support multiple log levels:
		INFO
		WARN
		ERROR
	-Logger should be thread-safe.
	-Logs should be printed with:
		timestamp
		log level
		message
	-The logger should be globally accessible.
*/

/*
// ----------lazy logger-------------
var instance *Logger

func GetLogger() *Logger {
	if instance == nil {
		return &Logger{}
	}
	return instance
}

// -----------------thread safe-------------
var (
	instance *Logger
	mu       sync.Mutex
)

func GetLogger() *Logger {

	if instance == nil {
		return &Logger{}
	}
	return instance
}

// --------------sync.Once--------------
var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}
*/

type LogType string

const (
	INFO  LogType = "INFO"
	WARN  LogType = "WARN"
	ERROR LogType = "ERROR"
)

type Log struct {
	logLevel  LogType
	timestamp time.Time
	message   string
}

type Logger struct{}

var instance = &Logger{}

func GetLogger() *Logger {
	return instance
}
