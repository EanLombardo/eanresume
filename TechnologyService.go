// TechnologyService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Technology struct {
	Name       string    `json:"name"`
	Experience string    `json:"experience"`
	WebsiteUrl string    `json:"url"`
	ImageUrl   string    `json:"image"`
	Text       string    `json:"text"`
	ACL        ACLSystem `json:"-"`
	Ordinal    int8      `json:"ordinal"`
}

type TechnologyList struct {
	Items []Technology `json:"items"`
}

type TechnologyService struct {
}

func (as *TechnologyService) List(r *http.Request, req *ListRequest, resp *TechnologyList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("technology").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Technology
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

func InsertTechnologyKind(context appengine.Context) {
	item := Technology{}
	item.Text = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "technology", nil), &item)
}

func RegisterTechnologyService() {
	service := &TechnologyService{}

	api, err := endpoints.RegisterService(service, "technologies", "v1", "Technologies API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "technologies", "List all technologys accessible to the given access code"
}