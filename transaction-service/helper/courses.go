package helper

import (
	"os"
	"strconv"
)

type FormatCourse struct {
	Status string `json:"status"`
	Data   Course `json:"data"`
}

type Course struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Type     string `json:"profession"`
	MentorID int    `json:"mentor_id"`
}

type FormatMentor struct {
	Status string `json:"status"`
	Data   Mentor `json:"data"`
}

type Mentor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCourseByID(id int) (FormatCourse, error) {

	urlService := os.Getenv("URL_SERVICE_COURSES")
	url := urlService + "/api/v1/courses/" + strconv.Itoa(id)

	var result FormatCourse

	err := GetJson(url, &result)

	if err != nil {

		return result, err

	}

	return result, nil

}

func GetMentorByID(id int) (FormatMentor, error) {

	urlService := os.Getenv("URL_SERVICE_COURSES")
	url := urlService + "/api/v1/mentors/" + strconv.Itoa(id)

	var result FormatMentor

	err := GetJson(url, &result)

	if err != nil {

		return result, err

	}

	return result, nil

}
