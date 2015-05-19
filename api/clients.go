package main

type Client struct {
	Host string
}

func (c *Client) CreateBuild(project string, git GitInfo, service string) (Build, error) {
	var respBuild Build
	build := Build{
		ProjectName: project,
		Git:         git,
		Service:     service,
	}

	r, err := makeRequest(c, "POST", "/builds/", build)
	if err != nil {
		return respBuild, err
	}
	err = processResponseEntity(r, &respBuild, 201)
	return respBuild, err
}

func (c *Client) ListBuilds(project string) ([]Build, error) {
	var respBuilds []Build

	r, err := makeRequest(c, "GET", "/builds?project_name="+project, nil)
	if err != nil {
		return respBuilds, err
	}
	err = processResponseEntity(r, &respBuilds, 200)
	return respBuilds, err
}
