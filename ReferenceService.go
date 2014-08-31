// ReferenceService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Reference struct {
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Text    string    `json:"text"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type ReferenceList struct {
	Items []Reference `json:"items"`
}

type ReferenceService struct {
}

func (as *ReferenceService) List(r *http.Request, req *ListRequest, resp *ReferenceList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("reference").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Reference
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

func InsertReferenceKind(context appengine.Context) {
	item := Reference{}
	item.Text = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "reference", nil), &item)
}

func RegisterReferenceService() {
	service := &ReferenceService{}

	api, err := endpoints.RegisterService(service, "references", "v1", "References API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "references", "List all references accessible to the given access code"
}
