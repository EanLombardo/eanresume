// LanguageService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Language struct {
	Name       string    `json:"name"`
	Experience string    `json:"experience"`
	WebsiteUrl string    `json:"url"`
	ImageUrl   string    `json:"image"`
	Text       string    `json:"text"`
	ACL        ACLSystem `json:"-"`
	Ordinal    int8      `json:"ordinal"`
}

type LanguageList struct {
	Items []Language `json:"items"`
}

type LanguageService struct {
}

func (as *LanguageService) List(r *http.Request, req *ListRequest, resp *LanguageList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("language").Order("Ordinal")

		for t := query.Run(context); ; {
			var item Language
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

func InsertLanguageKind(context appengine.Context) {
	item := Language{}
	item.Text = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "language", nil), &item)
}

func RegisterLanguageService() {
	service := &LanguageService{}

	api, err := endpoints.RegisterService(service, "languages", "v1", "Languages API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "languages", "List all languages accessible to the given access code"
}
