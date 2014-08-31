// AchievementService.go
package main

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Achievement struct {
	Text    string    `json:"text"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type AchievementList struct {
	Items []Achievement `json:"items"`
}

type AcheivementService struct {
}

func (as *AcheivementService) List(r *http.Request, req *ListRequest, resp *AchievementList) error {
	context := appengine.NewContext(r)

	if req.validate(context) {
		query := datastore.NewQuery("achievement").Order("Ordinal")

		for t := query.Run(context); ; {
			var ach Achievement
			_, err := t.Next(&ach)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return err
			}

			if ach.ACL.hasAccess(req.AccessKey) {
				resp.Items = append(resp.Items, ach)
			}
		}
	}
	return nil
}

func InsertAchievementKind(context appengine.Context) {
	ach := Achievement{}
	ach.Text = "kind"
	datastore.Put(context, datastore.NewIncompleteKey(context, "achievement", nil), &ach)
}

func RegisterAchievementService() {
	service := &AcheivementService{}

	api, err := endpoints.RegisterService(service, "achievements", "v1", "Achievements API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc = "list", "GET", "achievements", "List all achievements accessible to the given access code"
}