package model

type (
	Course struct {
		ID int64
		// Name, required, contains the lecture code and lecture name
		Name string
		// Publisher The contribution Team or person.
		// TODO create struct for it.
		Publisher string
		// Disciplines, Categories of Courses.
		Discipline string
		License    string
		// Version, optional, defined the version of lecture
		// example: `23 fall`, `24 spring`, `2023`, `2024`.
		Version string
		TimestampedModel

		ResourceAddr string

		// Lecture Description.
		// LectureDesc struct {
		//	 Overview        string
		//	 Prerequisites   string
		//	 CourseGoals     string
		//	 FollowupCourses string
		// }

		// TODO Lecture Chapter Index
		// ChapterList []Chapter
		// TODO Lecture Notice
		// NoticeList  []Notice
	}

	LectureDesc struct {
		Overview        string
		Prerequisites   string
		CourseGoals     string
		FollowupCourses string
	}

	Chapter struct {
		ChapterID          int64
		ChapterName        string
		ResourceCollection []Resource
	}
)