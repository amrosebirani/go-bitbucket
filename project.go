package bitbucket

import (
	"encoding/json"
	"os"

	"github.com/k0kubun/pp"
)

type Projects struct {
	c *Client
}

func (r *Projects) buildProjectsTeamBody(ro *ProjectsTeamOptions) string {

	body := map[string]interface{}{}

	if ro.Description != "" {
		body["description"] = ro.Description
	}
	if ro.Key != "" {
		body["key"] = ro.Key
	}
	if ro.Name != "" {
		body["name"] = ro.Name
	}
	body["is_private"] = ro.Private

	data, err := json.Marshal(body)
	if err != nil {
		pp.Println(err)
		os.Exit(9)
	}

	return string(data)
}

func (r *Projects) CreateOnTeam(ro *ProjectsTeamOptions) (interface{}, error) {
	data := r.buildProjectsTeamBody(ro)
	urlStr := r.c.requestUrl("/teams/%s/projects/", ro.Team)
	return r.c.execute("POST", urlStr, data)
}
