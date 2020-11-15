package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/linkedin"
)

func searchService(ln *linkedin.Linkedin, keywords string) error {
	svcNode, err := ln.SearchService(keywords)
	if err != nil {
		return err
	}

	svcs := make([]linkedin.Service, 0)
	for svcNode.Next() {
		svcs = append(svcs, svcNode.Elements...)
		if len(svcs) >= 20 {
			break
		}
	}

	svcNode.Elements = svcs
	f, err := os.Create("services.json")
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(svcNode)
}