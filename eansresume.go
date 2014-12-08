// endpoint.go
package eansresume

import (
	"appengine"
	"github.com/crhym3/go-endpoints/endpoints"
)

func InsertAllKinds(context appengine.Context) {
	InsertIntroductionKind(context)
	InsertProjectKind(context)
	InsertReferenceKind(context)
	InsertTechnologyKind(context)
}

func init() {
	RegisterAchievementService()
	RegisterLanguageService()
	RegisterExperienceService()
	RegisterIntroductionService()
	RegisterProjectService()
	RegisterReferenceService()
	RegisterTechnologyService()
	RegisterAccessKeyService()
	RegisterDocumentService()
	endpoints.HandleHTTP()
}
