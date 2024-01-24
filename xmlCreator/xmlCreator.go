package xmlCreator

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Summary struct {
        Value string `xml:",chardata"`
}

type Process struct {
        ID      string  `xml:"id,attr,omitempty"`
        PrntID  string  `xml:"parentProcessId,attr,omitempty"`
        Name    string  `xml:"name,attr,omitempty"`
        Type    string  `xml:"type,attr"`
        Summary Summary `xml:"Summary,omitempty"`
}


type Ticket struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type Concept struct {
	Value        string `xml:"value,attr"`
	ConceptLink struct {
		PageID string `xml:"pageId,attr"`
	} `xml:"conceptLink"`
	Tickets []Ticket `xml:"ticket"`
        Processes []Process `xml:"gkProcess"`
	Version struct {
		Product string `xml:"product,attr"`
		Project string `xml:"project,attr"`
	} `xml:"version"`
}

type Template struct {
        XMLName        xml.Name  `xml:"pdp:gap"`
	ID             string    `xml:"id,attr"`
	Name           string    `xml:"name,attr"`
	PDP            string    `xml:"xmlns:pdp,attr"`
	XSI            string    `xml:"xmlns:xsi,attr"`
	SchemaLocation string    `xml:"xsi:schemaLocation,attr"`
	Concept       Concept `xml:"conceptVersion,omitempty"`
}

func BuildFile(templateInstance Template) {
        xmlData, err := xml.MarshalIndent(templateInstance, "", "  ")

        if err != nil {
                fmt.Println("Error masrhaling XML:", err)
                return
        }

        file, err := os.Create(templateInstance.ID + " - " + templateInstance.Name + ".xml")
        if err != nil {
                fmt.Println("Error creating file:", err)
        }

        defer file.Close()

        file.Write(xmlData)
        fmt.Println("XML data written to", file.Name())
}
