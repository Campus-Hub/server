package model

type (
	Contributor struct {
		// Contributor ID, auto generate.
		ID int64
		// Contributor Name, required.
		Name string
		// Contributor's Organization. Optional.
		Org string

		// The Role of Contributor in Course.
		// including Instructor, TeachingAssistant, CommunityContributor
		Role string
	}

	// ContributorTeam
	ContributorTeam struct {
		ID int64

		Name string

		Contributors []Contributor
	}
)

const (
	// Contributor Roles
	Instructor           = "INS"
	TeachingAssistant    = "TA"
	CommunityContributor = "CC"
)
