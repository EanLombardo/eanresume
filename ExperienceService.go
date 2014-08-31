// ExperienceService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Experience struct {
	CompanyName string    `json:"companyName"`
	WorkPeriod  string    `json:"workPeriod"`
	Content     string    `json:"content" datastore:"Content,noindex"`
	CompanyUrl  string    `json:"companyUrl"`
	ACL         ACLSystem `json:"-"`
	Ordinal     int8      `json:"ordinal"`
}

type ExperienceList struct {
	Items []Experience `json:"items"`
}

type ExperienceService struct {
}

func (as *ExperienceService) List(r *http.Request, req *ListRequest, resp *ExperienceList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("experience").Order("Ordinal")
		for t := query.Run(context); ; {
			var item Experience
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

func InsertExperienceKind(context appengine.Context) {
	item := Experience{}
	item.Content = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "experience", nil), &item)
}

func RegisterExperienceService() {
	service := &ExperienceService{}

	api, err := endpoints.RegisterService(service, "experiences", "v1", "Experiences API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "experiences", "List all experiences accessible to the given access code"
}
