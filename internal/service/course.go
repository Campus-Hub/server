package service

import (
	"github.com/Campus-Hub/server/internal/pack"
)

type (
	CourseSrv struct {
		//PageSize int `json:"pagesize"`
		//Offset   int `json:"offset"`
	}
)

// List Courses by Conditional.
func (svr CourseSrv) List() []pack.Course {
	//var (
	//	courses []model.Course
	//	err     error
	//)
	//
	//err = model.DB.Model(model.Course{}).Find(&courses).Error
	//if err != nil {
	//	logger.Logger.Error(err)
	//}
	//
	//packedList := pack.PackCourseList(courses)
	//
	//return packedList

	return []pack.Course{
		{
			Name:            "Introduction to Programming",
			ContributorTeam: "Code Masters",
			Discipline:      "Computer Science",
			License:         "MIT",
			Origination:     "University XYZ",
			Version:         "2022",
			ResourceAddr:    "https://example.com/intro-to-programming",
		},
		{
			Name:            "Advanced Data Structures",
			ContributorTeam: "Algo Wizards",
			Discipline:      "Computer Science",
			License:         "Apache 2.0",
			Origination:     "Institute ABC",
			Version:         "Spring 2023",
			ResourceAddr:    "https://example.com/advanced-data-structures",
		},
		{
			Name:            "Artificial Intelligence Basics",
			ContributorTeam: "AI Explorers",
			Discipline:      "Artificial Intelligence",
			License:         "GNU GPL",
			Origination:     "Research Institute DEF",
			Version:         "23 fall",
			ResourceAddr:    "https://example.com/ai-basics",
		},
	}
}
