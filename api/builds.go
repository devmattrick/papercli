package api

import "fmt"

type BuildsResponse struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

func GetBuilds(project string, version string) (resp *BuildsResponse, err error) {
	url := fmt.Sprintf("https://papermc.io/api/v2/projects/%s/versions/%s", project, version)

	resp = &BuildsResponse{}
	err = GetJson(url, resp)

	return resp, err
}

type BuildResponse struct {
	ProjectId   string                 `json:"project_id"`
	ProjectName string                 `json:"project_name"`
	Version     string                 `json:"version"`
	Build       int                    `json:"build"`
	Time        string                 `json:"time"`
	Channel     string                 `json:"channel"`
	Promoted    bool                   `json:"promoted"`
	Changes     []BuildResponseChange  `json:"changes"`
	Downloads   BuildResponseDownloads `json:"downloads"`
}

type BuildResponseChange struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}

type BuildResponseDownloads struct {
	Application    BuildResponseDownload `json:"application"`
	MojangMappings BuildResponseDownload `json:"mojang-mappings"`
}

type BuildResponseDownload struct {
	Name   string `json:"name"`
	Sha256 string `json:"sha256"`
}

func GetBuild(project string, version string, build int) (resp *BuildResponse, err error) {
	url := fmt.Sprintf("https://papermc.io/api/v2/projects/%s/versions/%s/builds/%d", project, version, build)

	resp = &BuildResponse{}
	err = GetJson(url, resp)

	return resp, err
}
