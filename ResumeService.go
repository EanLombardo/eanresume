package eanresume

import (
	"appengine"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type Resume struct {
	Introduction    []Introduction		`json:"introductions"`
	Achievement 	[]Achievement		`json:"achievements"`
	Experience	[]Experience		`json:"experiences"`
	Language	[]Language		`json:"languages"`
	Document	[]Document		`json:"documents"`
	Reference	[]Reference		`json:"references"`
	Technology	[]Technology		`json:"technologies"`
	Project		[]Project		`json:"projects"`
}

type Response struct{
	Text string
}
type ResumeService struct {
	
}

func (as *ResumeService) Get(r *http.Request, req *ResumeRequest, resp *Resume) error {
	context := appengine.NewContext(r)
	
	access := AccessKey{req.AccessKey,""}
	if req.validate(context) {
		resp.Introduction,_ = GetIntroductions(access,context)
		resp.Achievement,_ = GetAchievements(access,context)
		resp.Experience,_ = GetExperiences(access,context)
		resp.Language,_ = GetLanguages(access,context)
		resp.Document,_ = GetDocuments(access,context)
		resp.Reference,_ = GetReferences(access,context)
		resp.Technology,_ = GetTechnologies(access,context)
		resp.Project,_ = GetProjects(access,context)
	}
	
	return nil
}

func (as *ResumeService) Awaken(r *http.Request, req *Response, resp *Response) error {
	resp = req
	
	return nil
}

func RegisterResumeService() {
	service := &ResumeService{}

	api, err := endpoints.RegisterService(service, "resume", "v1", "Resume API", true)

	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("Awaken").Info()
	info.Name,info.HTTPMethod, info.Path, info.Desc = "awaken","GET", "api", "Awakens the resume service"
	
	info2 := api.MethodByName("Get").Info()
	info2.Name,info2.HTTPMethod, info2.Path, info2.Desc = "get","GET", "resume", "Fetches a resume"
}
