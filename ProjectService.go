// ProjectService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Project struct {
	Name         string    `json:"name"`
	RepoUrl      string    `json:"repoUrl"`
	Technologies string    `json:"technologies"`
	Content      string    `json:"content"`
	ACL          ACLSystem `json:"-"`
	Ordinal      int8      `json:"ordinal"`
}

type ProjectList struct {
	Items []Project `json:"items"`
}

type ProjectService struct {
}

func (as *ProjectService) List(r *http.Request, req *ListRequest, resp *ProjectList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("project").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Project
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return err
			}

			if item.ACL.hasAccess(req.AccessKey) {
				resp.Items = append(resp.Items, item)
			}
		}
	}
	return nil
}

func InsertProjectKind(context appengine.Context) {
	item := Project{}
	item.Content = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "project", nil), &item)
}

func RegisterProjectService() {
	service := &ProjectService{}

	api, err := endpoints.RegisterService(service, "projects", "v1", "Projects API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "projects", "List all projects accessible to the given access code"
}
