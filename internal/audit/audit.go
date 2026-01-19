package audit

import (
	"encoding/json"
	"log"
	"time"
)

// Logger handles audit logging.
type Logger struct {
	logger  *log.Logger
	enabled bool
}

// New creates a new audit logger.
func New(logger *log.Logger, enabled bool) *Logger {
	return &Logger{
		logger:  logger,
		enabled: enabled,
	}
}

// Log writes an audit log entry as JSON.
func (l *Logger) Log(entry AuditLog) error {
	if !l.enabled {
		return nil
	}

	if entry.Timestamp.IsZero() {
		entry.Timestamp = time.Now().UTC()
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	l.logger.Printf("AUDIT %s", string(data))
	return nil
}

// Entry creates a new audit log entry with the given parameters.
func Entry(userID *int, username, userIP, action, resourceType string) AuditLog {
	return AuditLog{
		Timestamp:    time.Now().UTC(),
		UserID:       userID,
		Username:     username,
		UserIP:       userIP,
		Action:       action,
		ResourceType: resourceType,
		Status:       "success",
	}
}
