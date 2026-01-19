package audit

import "time"

// AuditLog represents a single audit log entry.
type AuditLog struct {
	Timestamp    time.Time              `json:"timestamp"`
	UserID       *int                   `json:"user_id,omitempty"`
	Username     string                 `json:"username"`
	UserIP       string                 `json:"user_ip"`
	Action       string                 `json:"action"`
	ResourceType string                 `json:"resource_type"`
	ResourceID   *int                   `json:"resource_id,omitempty"`
	ResourceName string                 `json:"resource_name,omitempty"`
	Changes      map[string]interface{} `json:"changes,omitempty"`
	Status       string                 `json:"status"` // "success" or "failure"
	ErrorMessage string                 `json:"error_message,omitempty"`
	Context      map[string]interface{} `json:"context,omitempty"`
}
