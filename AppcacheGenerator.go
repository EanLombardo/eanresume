//This is an awesome little bit of code. This usses text templates to automatically generate the applications appcache file to include all of the images fromt the content entered into the system
//This does reveal all of the images to do with all languages and technologies in my resume, so if I hide one in a resume that person could look here to see it, but I don't anticipate a need for that
//Best of all, if I make a change to a file that doesn't require a manifest change the manifest updates automatically because it contains the appengine project version which increments auomaticaly with
//ever upload

package eanresume

import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"text/template"
)


type Appcache struct
{
	Version string
	Languages []Language
	Technologies []Technology
}

func SetupAppcacheHandler(){
	http.HandleFunc("/cache.manifest", HandleAppcahce)
	http.HandleFunc("/cache.appcache", HandleAppcahce)
}

func HandleAppcahce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/cache-manifest")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	context := appengine.NewContext(r)

	var appcache = Appcache{appengine.VersionID(context),[]Language{},[]Technology{}};

	query := datastore.NewQuery("language")
	query.GetAll(context, &appcache.Languages)

	query = datastore.NewQuery("technology")
	query.GetAll(context, &appcache.Technologies)

    tpl:= template.Must(template.ParseFiles("appcache.tt"))

    err := tpl.Execute(w,appcache)

    if err != nil{
    	context.Errorf("There was an error generating the appcache %#v",err)
    }
}