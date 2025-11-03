package model

import (
	"time"
)

type CheckRun struct {
	ID             int64
	RepoID         int64
	AppID          int64
	CommitID       string
	MergeRequestID *int64
	Branch         *string
	ExternalID     *string
	Name           string
	Description    string // 长度限制 255
	Text           string // 长度限制 65535
	TextHTML       string // 长度限制 65535
	Status         CheckRunStatus
	Conclusion     *CheckRunConclusion
	StartedAt      *time.Time
	CompletedAt    *time.Time
	DetailsURL     *string
	Annotations    []*Annotation
	Operations     []*Operation
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DisplayMode    DisplayMode

	// Assembled by QueryService. TODO: Bypassed *bool
	Required *bool
}

func (c *CheckRun) Completed() bool {
	return c != nil && c.Status == CheckRunStatusCompleted
}

func (c *CheckRun) Passed() bool {
	return c.Completed() && c.Conclusion != nil && c.Conclusion.Passed()
}

//go:generate go run github.com/abice/go-enum -f=$GOFILE

// CheckRunStatus
// ENUM(queued,in_progress,completed)
type CheckRunStatus string

// CheckRunConclusion
// ENUM(operation_required,timed_out,canceled,failed,warning,neutral,succeeded)
type CheckRunConclusion string

func (c CheckRunConclusion) Passed() bool {
	switch c {
	case CheckRunConclusionSucceeded, CheckRunConclusionNeutral, CheckRunConclusionWarning:
		return true
	default:
		return false
	}
}

// DisplayMode
// ENUM(default,hidden,highlighted)
type DisplayMode string

type Annotation struct {
	ID          *string         `json:"id,omitempty"` // 标识符，当 Operations 有值时该字段必须有值，长度限制 64，超过截断
	Path        string          `json:"path"`
	StartLine   int             `json:"start_line"`
	StartColumn *int            `json:"start_column,omitempty"`
	EndLine     int             `json:"end_line"`
	EndColumn   *int            `json:"end_column,omitempty"`
	Level       AnnotationLevel `json:"level"`
	Message     string          `json:"message,omitempty"`    // 长度限制 2048，超过截断
	DetailsURL  *string         `json:"details_url"`          // 长度限制 1024，超过截断
	Operations  []*Operation    `json:"operations,omitempty"` // 长度限制 8，超过截断
	Folded      *bool           `json:"folded,omitempty"`     // 展示控制，表示是否默认折叠，true 表示默认折叠，false 表示默认展开，不指定由页面自行控制
}

func (a *Annotation) GetID() string {
	if a == nil || a.ID == nil {
		return ""
	}
	return *a.ID
}

// AnnotationLevel
// ENUM(info,warning,error,critical)
type AnnotationLevel string

type Operation struct {
	ID          string  `json:"id"`                    // 标识符，长度限制 64，超过截断
	Label       string  `json:"label"`                 // 按钮显示文字，长度限制 32，超过截断
	Description string  `json:"description,omitempty"` // 作用简短描述，长度限制 255，超过截断
	TargetURL   *string `json:"target_url,omitempty"`  // 点击按钮后跳转的目标链接，长度限制 1024，超过截断
}

// UnfinalizedApp is to represent the unfinalized app, consisted of a part of app.
