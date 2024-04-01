package main

import (
	"testing"
)

func TestNormalParse(t *testing.T) {
	remoteString := `{"id":"1", "attributes": {"name": "response 1", "description": "Longer text saying more details", "label": "Rensponse"}}`
	remoteParsed := GetRemoteResponse(remoteString)

	parsedAttributes := RemoteAttributes{}
	err := FromMap(remoteParsed.Attributes, &parsedAttributes)
	if err != nil {
		t.Errorf("Ohh no %v", err)
		return
	}
	if parsedAttributes.Name != "response 1" {
		t.Errorf("ParsedName got %s", parsedAttributes.Name)
		return
	}
	if parsedAttributes.Description != "Longer text saying more details" {
		t.Errorf("ParsedDescription got %s", parsedAttributes.Description)
		return
	}
	if parsedAttributes.Label != "Rensponse" {
		t.Errorf("ParsedLabel got %s", parsedAttributes.Label)
		return
	}
}
func TestWithMissingParse(t *testing.T) {
	remoteString := `{"id":"1", "attributes": {"name": "response 1", "description": "Longer text saying more details" }}`
	remoteParsed := GetRemoteResponse(remoteString)

	parsedAttributes := RemoteAttributes{}
	err := FromMap(remoteParsed.Attributes, &parsedAttributes)
	if err != nil {
		t.Errorf("Ohh no %v", err)
		return
	}
	if parsedAttributes.Name != "response 1" {
		t.Errorf("ParsedName got %s", parsedAttributes.Name)
		return
	}
	if parsedAttributes.Description != "Longer text saying more details" {
		t.Errorf("ParsedDescription got %s", parsedAttributes.Description)
		return
	}
	if parsedAttributes.Label != "" {
		t.Errorf("ParsedLabel got %s", parsedAttributes.Label)
		return
	}
}

func TestWithNumbers(t *testing.T) {
	remoteString := `{"id":"1", "attributes": {"name": "response 1", "total": 900, "page": 8 }}`
	remoteParsed := GetRemoteResponse(remoteString)

	parsedAttributes := RemoteAttributesWithNumber{}
	err := FromMap(remoteParsed.Attributes, &parsedAttributes)
	if err != nil {
		t.Errorf("Ohh no %v", err)
		return
	}
	if parsedAttributes.Name != "response 1" {
		t.Errorf("ParsedName got %s", parsedAttributes.Name)
		return
	}
	if parsedAttributes.Total != 900 {
		t.Errorf("ParsedDescription got %d", parsedAttributes.Total)
		return
	}
	if parsedAttributes.Page != 8 {
		t.Errorf("ParsedLabel got %d", parsedAttributes.Page)
		return
	}
	if parsedAttributes.IgnoreMe != 0 {
		t.Errorf("ParsedLabel got %d", parsedAttributes.Page)
		return
	}
}
