package models

type Lesson struct {
	ID            uint
	LevelID       uint
	Name          string
	BaseDirectory string
}

type RequestNewLesson struct {
	Name     string `json:"name" binding:"required"`
	CourseID uint   `json:"course_id" binding:"required"`
	LevelID  uint   `json:"level_id" binding:"required"`
	Markup   string `json:"code" binding:"required"`
}

var htmlBase = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Summernote</title>
  <link href="http://netdna.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.css" rel="stylesheet">
  <link href="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.8/summernote.css" rel="stylesheet">
  <script src="http://cdnjs.cloudflare.com/ajax/libs/summernote/0.8.8/summernote.js"></script>
</head>
<body>`

//var err Error

func (r *RequestNewLesson) Add() error {

	// course := Course{}
	// if db.First(&course, "id = ?", r.CourseID).RecordNotFound() {
	// 	return errors.New("course not found")
	// }
	//
	// if r.LevelID < 1 || r.LevelID > 5 {
	// 	return errors.New("Invalid course")
	// }
	//
	// levelPath := course.BaseDirectory + "/Level_"+r.LevelID
	// lessonPath := levelPath+"/"
	// if _, err := os.Stat(levelPath); err == nil {
	// 	//directory exists
	// } else {
	// 	if err := os.Mkdir(levelPath, os.FileMode(0777)); err != nil {
	// 		return err
	// 	}
	//
	// 	insertOverview(levelPath)
	// }
	//
	return nil
}
