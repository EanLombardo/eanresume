// AccessKeyService
package eanresume

import (
	"appengine"
	"appengine/datastore"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

const clientId = "REDACTED"

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.APIExplorerClientID}
	audiences = []string{clientId}
)

type AccessKey struct {
	AccessKey string `json:"accessKey"`
	Owner     string
}

type AccessKeyList struct {
	Keys []AccessKey `json:"keys"`
}

type EmptyRequest struct {
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

func (as *AccessKeyService) List(r *http.Request, req *EmptyRequest, resp *AccessKeyList) error {
	context := appengine.NewContext(r)

	u, err := endpoints.CurrentUser(endpoints.NewContext(r), scopes, audiences, clientIds)

	if err != nil {
		return err
	}

	context.Errorf("User %v", u)

	if u.String() != "el3h6@mst.edu" {
		return nil
	}

	query := datastore.NewQuery("accessKey")
	_,err = query.GetAll(context, &resp.Keys)

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
	info.Name,info.HTTPMethod, info.Path, info.Desc = "get","GET", "get", "Get the owner of an access key"

	info = api.MethodByName("List").Info()
	info.Name,info.HTTPMethod, info.Path, info.Desc = "list","GET", "list", "Lists all access keys. Requires authentication"
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences
}
