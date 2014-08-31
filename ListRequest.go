package main

import (
	"appengine"
	"appengine/datastore"
)

type ListRequest struct {
	AccessKey string `json:"accessKey"`
}

type AccessKeyDatastore struct {
	AccessKey string
	Owner     string
}

func (lr *ListRequest) validate(context appengine.Context) bool {
	query := datastore.NewQuery("accessKey").Filter("AccessKey =", lr.AccessKey)
	count, _ := query.Count(context)
	return count == 1
}