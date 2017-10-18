package models

import (
	"os"
	"strconv"
)

type Course struct {
	ID            uint
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	BaseDirectory string
}

func (c *Course) Add() error {
	tx := db.Begin()

	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return err
	}

	appDirectory := os.Getenv("APP_DIRECTORY")
	strCID := strconv.FormatUint(uint64(c.ID), 10)
	path := appDirectory + "_" + "course" + "_" + c.Name + "_" + strCID

	if err := os.Mkdir(path, os.FileMode(0522)); err != nil {
		tx.Rollback()
		return err
	}

	c.BaseDirectory = path

	tx.Commit()
	return nil
}
