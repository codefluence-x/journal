package journall

type Journal interface {
	// Info(msg string) JournalInfo
	// InfoWithTags(msg string, tags ...string) JournalInfo
	// InfoWithName(name, msg string) JournalInfo

	// Warning(msg string) JournalWarning
	// WarningWithTags(msg string, tags ...string) JournalWarning
	// WarningWithName(name, msg string) JournalWarning

	// Error(msg string, err error) JournalError
	// ErrorWithTags(msg string, err error, tags ...string) JournalError
	// ErrorWithName(name, msg string, err error) JournalError
}

type JournalBase interface {
	// Name() string
	// Tags() []string
	// Field(field string, value interface{}) JournalBase

	// LogType() string

	// SetName(name string) JournalBase
	// SetTags(tags ...string) JournalBase

	// Log()
}

type JournalInfo interface {
	JournalBase
}

type JournalWarning interface {
	JournalBase
}

type JournalError interface {
	JournalBase
}

type journal struct {
	msg    string
	err    string
	errRaw error
	name   string
	tags   []string
	log    string
}
