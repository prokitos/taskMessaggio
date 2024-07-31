package models

type Grades struct {
	Id          int   `json:"grade_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	RusLanguage int64 `json:"language"`
	Mathematics int64 `json:"math"`
}

type Statistic struct {
	Id      int    `json:"statistic_id" example:"" gorm:"unique;primaryKey;autoIncrement"`
	Metrics string `json:"metric"`
	Value   int64  `json:"value"`
}
