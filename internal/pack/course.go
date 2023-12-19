package pack

import (
	"strconv"

	"github.com/Campus-Hub/server/internal/model"
)

type (
	Course struct {
		Name            string
		ContributorTeam string
		Discipline      string
		License         string
		Origination     string
		Version         string
		ResourceAddr    string
	}
)

func Covert() {

}

func PackCourse(course model.Course) Course {
	return Course{
		Name:            course.Name,
		ContributorTeam: strconv.FormatInt(course.ContributorTeamID, 10),
		Discipline:      course.Discipline,
		License:         course.License,
		Origination:     course.Origination,
		Version:         course.Version,
		ResourceAddr:    course.ResourceAddr,
	}
}

func PackCourseList(courses []model.Course) (PackedCourses []Course) {

	for _, cours := range courses {
		PackedCourses = append(PackedCourses, PackCourse(cours))
	}

	return
}
