// LanguageService.go
package eanresume

import (
	"appengine"
	"appengine/datastore"
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

type Achievement struct {
	Text    string    `json:"text"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type Document struct {
	Name        string    `json:"name"`
	DocumentUrl string    `json:"url"`
	ImageUrl    string    `json:"image"`
	ACL         ACLSystem `json:"-"`
	Ordinal     int8      `json:"ordinal"`
}

type Experience struct {
	CompanyName string    `json:"companyName"`
	WorkPeriod  string    `json:"workPeriod"`
	Content     string    `json:"content" datastore:"Content,noindex"`
	CompanyUrl  string    `json:"companyUrl"`
	ACL         ACLSystem `json:"-"`
	Ordinal     int8      `json:"ordinal"`
}

type Introduction struct {
	Content string    `json:"content"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type Project struct {
	Name         string    `json:"name"`
	RepoUrl      string    `json:"repoUrl"`
	Technologies string    `json:"technologies"`
	Content      string    `json:"content"`
	ACL          ACLSystem `json:"-"`
	Ordinal      int8      `json:"ordinal"`
}

type Reference struct {
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Text    string    `json:"text"`
	ACL     ACLSystem `json:"-"`
	Ordinal int8      `json:"ordinal"`
}

type Technology struct {
	Name       string    `json:"name"`
	Experience string    `json:"experience"`
	WebsiteUrl string    `json:"url"`
	ImageUrl   string    `json:"image"`
	Text       string    `json:"text"`
	ACL        ACLSystem `json:"-"`
	Ordinal    int8      `json:"ordinal"`
}

func InsertType(context appengine.Context, kind string,dst interface{}){
	key:= datastore.NewIncompleteKey(context,kind,nil)
	datastore.Put(context,key,dst)
}

func InsertTypes(context appengine.Context){
	acl:= ACLSystem{true,"google3926"}
	InsertType(context,"achievement",&Achievement{"I did stuff",acl,2})
	InsertType(context,"document",&Document{"Name","http://f.com/d.pdf","null",acl,2})
	InsertType(context,"experience",&Experience{"Name co","now","sdvsdvsd","http://flarg.com",acl,2})
	InsertType(context,"introduction",&Introduction{"defwefwefwefwefwefwefwf",acl,2})
	InsertType(context,"project",&Project{"This","http://goof.com","android,stuff,windows","yapyapyap",acl,2})
	InsertType(context,"reference",&Reference{"Person","235235234235","sefZ@sd.net","A person",acl,2})
	InsertType(context,"technology",&Technology{"Android","Alot n stuff","http://android.com","null","Android is a platform",acl,2})
}

func GetLanguages(key AccessKey, context appengine.Context) ([]Language, error) {
	var items []Language
	if key.validate(context) {
		query := datastore.NewQuery("language").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Language
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetAchievements(key AccessKey, context appengine.Context) ([]Achievement, error) {
	var items []Achievement
	if key.validate(context) {
		query := datastore.NewQuery("achievement").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Achievement
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetDocuments(key AccessKey, context appengine.Context) ([]Document, error) {
	var items []Document
	if key.validate(context) {
		query := datastore.NewQuery("document").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Document
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetExperiences(key AccessKey, context appengine.Context) ([]Experience, error) {
	var items []Experience
	if key.validate(context) {
		query := datastore.NewQuery("experience").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Experience
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetIntroductions(key AccessKey, context appengine.Context) ([]Introduction, error) {
	var items []Introduction
	if key.validate(context) {
		query := datastore.NewQuery("introduction").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Introduction
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetProjects(key AccessKey, context appengine.Context) ([]Project, error) {
	var items []Project
	if key.validate(context) {
		query := datastore.NewQuery("project").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Project
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetReferences(key AccessKey, context appengine.Context) ([]Reference, error) {
	var items []Reference
	if key.validate(context) {
		query := datastore.NewQuery("reference").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Reference
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func GetTechnologies(key AccessKey, context appengine.Context) ([]Technology, error) {
	var items []Technology
	if key.validate(context) {
		query := datastore.NewQuery("technology").Order("Ordinal")
		

		for t := query.Run(context); ; {
			var item Technology
			_, err := t.Next(&item)

			if err == datastore.Done {
				break
			}

			if err != nil {
				return nil,err
			}

			if item.ACL.hasAccess(key.AccessKey) {
				items = append(items, item)
			}
		}
	}
	return items, nil
}