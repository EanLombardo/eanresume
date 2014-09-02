// DocumentService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Document struct {
	Name        string    `json:"name"`
	DocumentUrl string    `json:"url"`
	ImageUrl    string    `json:"image"`
	ACL         ACLSystem `json:"-"`
	Ordinal     int8      `json:"ordinal"`
}

type DocumentList struct {
	Items []Document `json:"items"`
}

type DocumentService struct {
}

func (as *DocumentService) List(r *http.Request, req *ListRequest, resp *DocumentList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("document").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Document
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

func InsertDocumentKind(context appengine.Context) {
	item := Document{}
	item.Name = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "document", nil), &item)
}

func RegisterDocumentService() {
	service := &DocumentService{}

	api, err := endpoints.RegisterService(service, "documents", "v1", "Documents API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "documents", "List all documents accessible to the given access code"
}
