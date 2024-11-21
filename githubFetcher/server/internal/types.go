package internal

// RawRepo is the fetched data of the repository
type RawRepo struct {
	Name string `json:"name"`
}

// RawCollaborator is the fetched data of the repository collaborator
type RawCollaborator struct {
	Login       string `json:"login"`
	SiteAdmin   bool   `json:"site_admin"`
	Permissions struct {
		Pull     bool `json:"pull"`
		Triage   bool `json:"triage"`
		Push     bool `json:"push"`
		Maintain bool `json:"maintain"`
		Admin    bool `json:"admin"`
	} `json:"permissions"`
	RoleName string `json:"role_name"`
}

// Repo is the normalized repository
type Repo struct {
	Name          string            `json:"name"`
	Collaborators []RawCollaborator `json:"collaborators"`
}
