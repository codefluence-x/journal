package journal

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Journal interface {
	// Sometimes log need tags to be grouped or just easier search in log architecture
	SetTags(tags ...string) Journal
	// Add custom field
	AddField(field string, value interface{}) Journal

	// Set track id of a log
	SetTrackId(trackId interface{}) Journal

	// Print the log
	Log()

	// Get raw json string to be logged
	Raw() string
}

type journalLogger struct {
	msg     string
	level   string
	errRaw  error
	tags    []string
	trackId interface{}
	fields  map[string]interface{}
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

func (j *journalLogger) SetTrackId(trackId interface{}) Journal {
	j.trackId = trackId
	return j
}

func (j *journalLogger) AddField(field string, value interface{}) Journal {
	j.fields[field] = value
	return j
}

func (j *journalLogger) Raw() string {
	return string(j.compileLog())
}

func (j *journalLogger) Log() {
	fmt.Println(j.Raw())
}

func (j *journalLogger) compileLog() []byte {
	j.appendAll()

	if j.level == "error" {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			j.fields["caller"] = fmt.Sprintf("%s:%d", file, no)
		}

		switch j.errRaw.(type) {
		case error:
			j.fields["error"] = j.errRaw.Error()
		default:
			j.fields["error"] = j.errRaw
		}

	}

	jsonEncodedString, _ := json.Marshal(j.fields)
	return jsonEncodedString
}

func (j *journalLogger) appendAll() {
	if j.trackId == nil {
		var trackId, _ = uuid.NewV4()
		j.trackId = trackId.String()
	}

	j.fields["track_id"] = j.trackId
	j.fields["message"] = j.msg
	j.fields["level"] = j.level

	if len(j.tags) > 0 {
		j.fields["tags"] = j.tags
	}

	j.fields["timestamp"] = time.Now()
}
