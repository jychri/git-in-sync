// Package atp (a test package) sets up testing environments
// for packages conf, gis and repos
package atp

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

var jmap = map[string][]byte{
	"recipes": []byte(`
		{
			"bundles": [{
				"path": "~/tmpgis",
				"zones": [{
						"user": "hendricius",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"pizza-dough",
							"the-bread-code"
						]
					},
					{
						"user": "cocktails-for-programmers",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"cocktails-for-programmers"
						]
					},
					{
						"user": "rochacbruno",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"vegan_recipes"
						]
					},
					{
						"user": "niw",
						"remote": "github",
						"workspace": "recipes",
						"repositories": [
							"ramen"
						]
					}
				]
			}]
		}
`),
}

// Setup creates a test environment at ~/tmpgis/$pkg/.
// ~/tmpgis/$pkg/gisrc.json is created, written with data from jmap[k]
// and its path is returned along with a cleanup function that removes
// ~/tmpgis/$pkg and all of its contents.
func Setup(pkg string, k string) (string, func()) {
	if pkg == "" {
		log.Fatalf("pkg is empty")
	}

	if _, ok := jmap[k]; ok != true {
		log.Fatalf("%v not found in jmap", k)
	}

	var u *user.User

	u, err := user.Current()

	if err != nil {
		log.Fatalf("Unable to identify current user (%v)", err.Error())
	}

	tb := path.Join(u.HomeDir, "tmpgis") // test base

	td := path.Join(tb, pkg) // test dir

	if err = os.MkdirAll(td, 0777); err != nil {
		log.Fatalf("Unable to create %v", td)
	}

	tg := path.Join(td, "gisrc.json") // test gisrc

	if err = ioutil.WriteFile(tg, jmap[k], 0777); err != nil {
		log.Fatalf("Unable to write to %v (%v)", tg, err.Error())
	}

	return tg, func() { os.RemoveAll(tb) }
}

// Result holds the expected values for a zone.
type Result struct {
	User, Remote, Workspace string
	Repos                   []string
}

// Results is a collection of Result structs.
type Results []Result

var rmap = map[string]Results{
	"recipes": {
		{"hendricius", "github", "recipes", []string{"pizza-dough", "the-bread-code"}},
		{"cocktails-for-programmers", "github", "recipes", []string{"cocktails-for-programmers"}},
		{"rochacbruno", "github", "recipes", []string{"vegan_recipes"}},
		{"niw", "github", "recipes", []string{"ramen"}},
	},
}

// Resulter returns expected results for testing.
func Resulter(k string) Results {
	if _, ok := rmap[k]; ok != true {
		log.Fatalf("%v not found in rmap", k)
	}

	return rmap[k]
}
