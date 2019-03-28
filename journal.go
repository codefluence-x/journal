package journal

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	uuid "github.com/satori/go.uuid"
)

var trackId, _ = uuid.NewV4()

type Journal interface {
	// Sometimes log need tags to be grouped or just easier search in log architecture
	SetTags(tags ...string) Journal
	// Add custom field
	AddField(field string, value interface{}) Journal

	// Print the log
	Log()
}

type journalLogger struct {
	msg    string
	level  string
	errRaw error
	tags   []string
	fields map[string]interface{}
}

func baseJournal(msg, level string) *journalLogger {
	return &journalLogger{msg: msg, level: level, fields: map[string]interface{}{}}
}

// New Journal info interface
func Info(msg string) Journal {
	return baseJournal(msg, "info")
}

// New Journal warning interface
func Warning(msg string) Journal {
	return baseJournal(msg, "warning")
}

// New Journal error interface
func Error(msg string, err error) Journal {
	base := baseJournal(msg, "error")
	base.errRaw = err
	return base
}

func (j *journalLogger) SetTags(tags ...string) Journal {
	j.tags = append(j.tags, tags...)
	return j
}

func (j *journalLogger) AddField(field string, value interface{}) Journal {
	j.fields[field] = value
	return j
}

func (j *journalLogger) Log() {
	j.appendAll()

	if j.level == "error" {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			j.fields["caller"] = fmt.Sprintf("%s:%d", file, no)
		}

		j.fields["error"] = j.errRaw
	}

	toBePrint, _ := json.Marshal(j.fields)
	fmt.Printf("%s\n", toBePrint)
}

func (j *journalLogger) appendAll() {
	j.fields["track_id"] = trackId.String()
	j.fields["message"] = j.msg
	j.fields["level"] = j.level

	if len(j.tags) > 0 {
		j.fields["tags"] = j.tags
	}

	j.fields["timestamp"] = time.Now()
}
