// AccessKeyService
package eanresume

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type AccessKey struct {
	AccessKey string `json:"accessKey"`
	Owner     string
}

type AccessKeyService struct {
}

func (as *AccessKeyService) Get(r *http.Request, req *ResumeRequest, resp *AccessKey) error {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("accessKey").Filter("AccessKey =", req.AccessKey)

	var akeys []AccessKey
	_, err := query.GetAll(context, &akeys)

	if len(akeys) == 1 {
		*resp = akeys[0]
	} else {
		resp = nil
	}

	return err
}

func (key *AccessKey) validate(context appengine.Context) bool {
	query := datastore.NewQuery("accessKey").Filter("AccessKey =", key.AccessKey)
	count, _ := query.Count(context)
	return count == 1
}


func RegisterAccessKeyService() {
	service := &AccessKeyService{}

	api, err := endpoints.RegisterService(service, "accessKeys", "v1", "AccessKeys API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("Get").Info()
	info.Name,info.HTTPMethod, info.Path, info.Desc = "get","GET", "accessKeys", "Get the owner of an access key"
}
