package models

import (
	"os"
	"strconv"
)

type Course struct {
	ID               uint   `json:"id"`
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	LargeDescription string `json:"large_description" binding:"required"`
	BaseDirectory    string `json:"-"`
}

type ProfileStats struct {
	UserName             string `json:"username"`
	UserEmail            string `json:"email"`
	SolvedExercises      int    `json:"solved_exercises"`
	TotalCourseExercises uint   `json:"total_course_ex"`
}

func (c *Course) Add() error {
	tx := db.Begin()

	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return err
	}

	appDirectory := os.Getenv("APP_DIRECTORY")
	strCID := strconv.FormatUint(uint64(c.ID), 10)
	path := appDirectory + "/" + "course" + "_" + c.Name + "_" + strCID

	if err := os.Mkdir(path, os.FileMode(0777)); err != nil {
		tx.Rollback()
		return err
	}

	tx.Model(c).Update("base_directory", path)

	for i := 1; i <= 5; i++ {
		level := Level{
			CourseID: c.ID,
			Number:   uint(i),
		}
		if err := Create(&level).Error; err != nil {
			tx.Rollback()
			return err
		}
		strID := strconv.FormatUint(uint64(level.Number), 10)
		if err := os.Mkdir(path+"/Level_"+strID, os.FileMode(0777)); err != nil {
			tx.Rollback()
			return err
		}
		if err := db.Model(level).Update("base_directory", path+"/Level_"+strID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (c Course) Update(id string) error {
	return db.Model(&Course{}).Where("id = ?", id).Update(c).Error
}

func GetCourses(last int) ([]Course, error) {
	courses := []Course{}
	return courses, db.Find(&courses, "id > ?", last).Limit(10).Error
}

func (c Course) GetUsersSuscribed() ([]ProfileStats, error) {
	users := []User{}
	profStats := []ProfileStats{}

	db.Table("subscriptions").Select("users.id,name,email").Joins("join users on subscriptions.user_id=users.id and subscriptions.course_id=?", c.ID).Scan(&users)

	for _, u := range users {
		sol := []Solution{}
		db.Find(&sol, "user_id = ? and course_id=? and status=?", u.ID, c.ID, "accepted")
		proStat := ProfileStats{
			UserName:        u.Name,
			UserEmail:       u.Email,
			SolvedExercises: len(sol),
		}
		profStats = append(profStats, proStat)
	}
	return profStats, nil
}
