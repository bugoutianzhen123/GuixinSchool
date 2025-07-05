package dao


// ChatInfo 代表某个学生的某次会话具体对话的内容
type ChatInfo struct {
	SessionID   uint64 `gorm:"column:session_id"`       // 会话ID
	StuID       string `gorm:"column:stu_id"`        // 学生ID
	Question    string `gorm:"type:text;column:question"` // 问题，长文本
	Answer      string `gorm:"type:text;column:answer"`   // 回答，长文本
	CreatedTime string `gorm:"column:created_time"`  // 创建时间
}

// 代表某个学生的某次会话
type Session struct{
	ID uint64 `gorm:"column:id;primaryKey"` //会话ID
	StuID string `gorm:"column:stu_id"` // 学生ID
	CreatedTime string `gorm:"column:created_time"` //创建时间
	Summary     string `gorm:"type:text;column:summary"`   // 会话总结（新增字段）
}