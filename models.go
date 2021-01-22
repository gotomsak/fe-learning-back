package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GetQuestionIDs 解く問題のIDと解いた問題のIDのstruct
type GetQuestionIDs struct {
	QuestionIDs []uint `json:"question_ids"`
	SolvedIDs   []uint `json:"solved_ids"`
}

// Question questionテーブルのstruct
type Question struct {
	gorm.Model
	Question    string `json:"question"`
	QimgPath    string `json:"qimg_path"`
	Mistake1    string `json:"mistake1"`
	Mistake2    string `json:"mistake2"`
	Mistake3    string `json:"mistake3"`
	Ans         string `json:"ans"`
	MimgPath1   string `json:"mimg_path1"`
	MimgPath2   string `json:"mimg_path2"`
	MimgPath3   string `json:"mimg_path3"`
	AimgPath    string `json:"aimg_path"`
	Season      string `json:"season"`
	QuestionNum string `json:"question_num"`
	Genre       string `json:"genre"`
}

// QuestionSend クライアントに送信する問題のstruct
type QuestionSend struct {
	QuestionID  uint     `json:"question_id"`
	Question    string   `json:"question"`
	QimgPath    []string `json:"qimg_path"`
	AnsList     []string `json:"ans_list"`
	AimgList    []string `json:"aimg_list"`
	Season      string   `json:"season"`
	QuestionNum string   `json:"question_num"`
	Genre       string   `json:"genre"`
}

// AnswerResultSectionIDSend クライアントに送信するsectionID
type AnswerResultSectionIDSend struct {
	AnswerResultSectionID uint `json:"answer_result_section_id"`
}

// AnswerResultSend checkAnswerのレスポンスをまとめたstruct
type AnswerResultSend struct {
	AnswerResultID uint   `json:"answer_result_id"`
	Result         string `json:"result"`
	Answer         string `json:"answer"`
}

// User userテーブルのstruct
type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"size:255"`
	Email          string `json:"email" gorm:"type:varchar(100);unique_index"`
	PasswordDigest string `json:"password_digest" gorm:"size:255"`
}

// UserSignup userのサインアップ時のstruct
type UserSignup struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
}

// UserSignin userのサインイン時のstruct
type UserSignin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSend Signup時に送られるdata
type UserSend struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}

// AnswerResult 解答の結果を保存するテーブルのstruct
type AnswerResult struct {
	gorm.Model
	UserID            uint   `gorm:"not null"`
	UserAnswer        string `gorm:"not null"` // userの選んだ答え
	AnswerResult      string `gorm:"not null"` // correctかincorrect
	ConcentrationData string
	MemoLog           string    `gorm:"type:text;"`
	OtherFocusSecond  uint      `json:"other_focus_second"`
	QuestionID        uint      `gorm:"not null"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
}

// AnswerResultSection 解答の結果のまとめを保存するテーブルのstruct
type AnswerResultSection struct {
	gorm.Model
	UserID              uint      `json:"user_id" gorm:"not null"`
	AnswerResultIDs     string    `json:"answer_result_ids"`
	CorrectAnswerNumber uint      `json:"correct_answer_number" gorm:"not null"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
}

// CheckAnswerSection sectionごとの問題が送信されてきた時のbindのstruct
type CheckAnswerSection struct {
	UserID              uint     `json:"user_id" gorm:"not null"`
	AnswerResultIDs     []uint64 `json:"answer_result_ids" gorm:"type:text;not null"`
	CorrectAnswerNumber uint     `json:"correct_answer_number" gorm:"not null"`
	StartTime           string   `json:"start_time"`
	EndTime             string   `json:"end_time"`
}

// CheckAnswer postされてきたdataのbind
type CheckAnswer struct {
	UserID            uint          `json:"user_id"`
	UserAnswer        string        `json:"user_answer"`
	MemoLog           string        `json:"memo_log"`
	OtherFocusSecond  uint          `json:"other_focus_second"`
	QuestionID        uint          `json:"question_id"`
	ConcentrationData []interface{} `json:"concentration_data"`
	StartTime         string        `json:"start_time"`
	EndTime           string        `json:"end_time"`
}

// Questionnaire アンケート結果を保存するテーブルのstruct
type Questionnaire struct {
	gorm.Model
	AnswerResultSectionID uint `json:"answer_result_section_id"`
	UserID                uint `json:"user_id" gorm:"not null"`
	Concentration         int  `json:"concentration"` // 集中
	WhileDoing            int  `json:"while_doing"`   // しながら
	Cheating              int  `json:"cheating"`      // カンニング
	Nonsense              int  `json:"nonsense"`      // デタラメ
}

// Frequency 最高最低頻度
type Frequency struct {
	gorm.Model
	UserID               uint    `gorm:"not null"`
	MaxBlinkFrequency    float64 `json:"max_blink_frequency"`
	MinBlinkFrequency    float64 `json:"min_blink_frequency"`
	MaxFaceMoveFrequency float64 `json:"max_face_move_frequency"`
	MinFaceMoveFrequency float64 `json:"min_face_move_frequency"`
	MaxBlinkNumber       float64 `json:"max_blink_number"`
	MinBlinkNumber       float64 `json:"min_blink_number"`
	MaxFaceMoveNumber    float64 `json:"max_face_move_number"`
	MinFaceMoveNumber    float64 `json:"min_face_move_number"`
}

// ConcentrationData 集中度の保存
// type ConcentrationData struct {
// 	UserID                uint      `json:"user_id"`
// 	AnswerResultSectionID uint      `json:"answer_result_section_id"`
// 	FaceImagePath         string    `json:"face_image_path"`
// 	Blink                 []float64 `json:"blink"`
// 	FaceMove              []float64 `json:"face_move"`
// 	Angle                 []float64 `json:"angle"`
// 	W                     []float64 `json:"w"`
// 	C1                    []float64 `json:"c1"`
// 	C2                    []float64 `json:"c2"`
// 	C3                    []float64 `json:"c3"`
// }

// ConcentrationData 集中度の保存
type ConcentrationData struct {
	ConcentrationData []interface{} `json:"concentration_data"`
}

// SonConcentrationData 集中度の保存
type SonConcentrationData struct {
	UserID                uint      `gorm:"not null"`
	AnswerResultSectionID uint      `json:"answer_result_section_id"`
	FaceImagePath         string    `gorm:"not null"`
	Concentration         []float64 `json:"concentration"`
}

// Results answer_result_section_idsをnosqlに保存
type Results struct {
	ResultIDs []uint64 `json:"answer_result_ids"`
}

type GetQuestionGymBind struct {
	NowLevel int `json:"now_level"`
}
type GetQuestionGymRes struct {
	QuestionID uint     `json:"question_id"`
	Question   string   `json:"question"`
	AnsList    []string `json:"ans_list"`
}
type QuestionsSub struct {
	ID       uint   `json:"id"`
	Question string `json:"question"`
	Mistake1 string `json:"mistake1"`
	Mistake2 string `json:"mistake2"`
	Mistake3 string `json:"mistake3"`
	Ans      string `json:"ans"`
}
