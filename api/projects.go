package api

type ProjectsResponse struct {
	Projects []string `json:"projects"`
}

func ListProjects() (resp *ProjectsResponse, err error) {
	resp = &ProjectsResponse{}
	err = GetJson("https://papermc.io/api/v2/projects", resp)

	return resp, err
}
