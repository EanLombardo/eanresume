// acl.go
package main

import (
	"strings"
)

type ACLSystem struct {
	IsWhiteList bool
	AccessKeys  string
}

func (acl *ACLSystem) hasAccess(accessKey string) bool {
	var isInKeys bool
	var usableKey = accessKey

	if usableKey == "" {
		usableKey = "nokey"
	}

	//The proper way would be to parse the CSV here, but this is faster as it doesn't spend time creating an array, and still only loops through the string once
	//AccessKeys should be an array instead of CSV, but I don't want to write an editor and the built in
	index := strings.LastIndex(acl.AccessKeys, usableKey)
	isInKeys = (index == 0 || (index != -1 && acl.AccessKeys[index] == ','))

	return ((!isInKeys && !acl.IsWhiteList) || (isInKeys && acl.IsWhiteList))
}
