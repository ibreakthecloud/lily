package models

type DataIssue struct {
	TableName     string `json:"table_name"`
	ColumnName    string `json:"column_name"`
	IssueType     string `json:"issue_type"`
	IssueSeverity string `json:"issue_severity"`
}
