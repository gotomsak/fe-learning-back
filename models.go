package main

import "github.com/jinzhu/gorm"

// GetQuestionIDs 解く問題のIDと解いた問題のIDのstruct
type GetQuestionIDs struct {
	QuestionIDs []int `json:"question_ids"`
	SolvedIDs   []int `json:"solved_ids"`
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

// User userテーブルのstruct
type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `gorm:"type:varchar(100);unique_index"`
	PasswordDigest string `json:"password_digest"`
}
