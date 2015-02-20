// endpoint.go
package eanresume

import
(
	"github.com/crhym3/go-endpoints/endpoints"
)

func init() {
	RegisterAccessKeyService()
	RegisterResumeService()
	SetupAppcacheHandler()
	endpoints.HandleHTTP()
}
