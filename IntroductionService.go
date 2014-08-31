// IntroductionService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Introduction struct {
	Content string    `json:"content"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type IntroductionList struct {
	Items []Introduction `json:"items"`
}

type IntroductionService struct {
}

func (as *IntroductionService) List(r *http.Request, req *ListRequest, resp *IntroductionList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("introduction").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Introduction
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

func InsertIntroductionKind(context appengine.Context) {
	item := Introduction{}
	item.Content = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "introduction", nil), &item)
}

func RegisterIntroductionService() {
	service := &IntroductionService{}

	api, err := endpoints.RegisterService(service, "introductions", "v1", "Introductions API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "introductions", "List all introductions accessible to the given access code"
}
