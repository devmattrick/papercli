package api

import "fmt"

type VersionsResponse struct {
	ProjectId     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

func GetVersions(project string) (resp *VersionsResponse, err error) {
	url := fmt.Sprintf("https://papermc.io/api/v2/projects/%s", project)

	resp = &VersionsResponse{}
	err = GetJson(url, resp)

	return resp, err
}
