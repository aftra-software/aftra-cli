package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	openapi "github.com/syndis-software/iris-api/pkg/openapi"
)

// VIKINGR_ACCESS_TOKEN
// iris create -uid <uid> -name <name> -score <score> --details=a=1,b=2,c=3

func stringToMap(str string) map[string]string {
	result := make(map[string]string)

	if str == "" {
		return result
	}
	// split the string into key-value pairs
	pairs := strings.Split(str, ",")
	// loop through each key-value pair
	for _, pair := range pairs {
		// split the pair into key and value
		kv := strings.Split(pair, "=")

		// skip empty key-value pairs
		if len(kv) != 2 || kv[0] == "" || kv[1] == "" {
			continue
		}

		// add the key-value pair to the map
		result[kv[0]] = kv[1]
	}

	return result
}

func getFlags() (string, string, string, map[string]string) {
	var uid, name, score, details_str string
	flag.StringVar(&uid, "uid", "", "Unique identifier for the opportunity")
	flag.StringVar(&name, "name", "", "Name of the opportunity")
	flag.StringVar(&score, "score", string(openapi.Unknown), "Risk score of the opportunity (critical, high, medium, low, info, none, unknown)")
	flag.StringVar(&details_str, "details", "", "Additional details. Comma separated key=value pairs.")
	flag.Parse()

	details := stringToMap(details_str)

	return uid, name, score, details
}

func usage() {
	fmt.Println(`
Create Opportunities
--------------------
iris create -uid=<uid> -name=<name> -score=<score> -details=<k>=<v>,<k2>=<v2>

uid: required
name: required
`)
}

func main() {
	uid, name, score, details := getFlags()

	if (uid == "") || (name == "") {
		usage()
		os.Exit(1)
	}

	createObj := openapi.CreateOpportunity{
		Name:    name,
		Uid:     uid,
		Score:   openapi.OpportunityScore(score),
		Details: details,
	}
	fmt.Printf("%v\n", createObj.Name)
	fmt.Printf("%v\n", createObj.Uid)
	fmt.Printf("%v\n", createObj.Score)
	fmt.Printf("%v\n", createObj.Details)
}
