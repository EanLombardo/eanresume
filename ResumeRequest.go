package eanresume

import (
	"appengine"
	"appengine/datastore"
)

type ResumeRequest struct {
	AccessKey string `json:"accessKey"`
}

func (lr *ResumeRequest) validate(context appengine.Context) bool {
	query := datastore.NewQuery("accessKey").Filter("AccessKey =", lr.AccessKey)
	count, _ := query.Count(context)
	return count == 1
}
