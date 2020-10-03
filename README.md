# db
## user
* id
* username
* password

## question

## AnswerResult(一問ごと)
gorm.Model
UserID         uint   `json:"user_id"`
UserAnswer     string `json:"user_answer"`   // userの選んだ答え
AnswerResult   string `json:"answer_result"` // correctかincorrect
MemoLog        string `json:"memo_log"`
StartTime      time.Time
EndTime        time.Time
OtherFocusSeconds uint `json:"other_focus_time"`
QuestionID     uint `json:"question_id"`

## getdata(10問ごと)
* id
* faceMovie
* startTime
* endTime
* nonFocusTime
* corrects 正解数
* resultIds
* selfAssessment



