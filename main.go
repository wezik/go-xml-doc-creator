package main

import (
	"bufio"
	"com/config"
	"com/xmlCreator"
	"fmt"
	"os"
	"strings"
)

func readInput() string {
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')

        if err != nil {
                fmt.Println("Error reading input:", err)
        }

        return strings.TrimSpace(input)
}

func main() {
        exitCode := 0
        for exitCode != 1 {
                
                fmt.Println("1. Create doc file", "X. Exit")
                fmt.Print("Input: ")
                userInput := readInput()

                switch userInput {
                case "1":
                        createNewDoc()
                case "x", "X":
                        exitCode = 1
                }

        }
}

func createNewDoc() {
        fmt.Println("Creating doc")
        fmt.Print("Enter name: ")
        nameInput := readInput()
        fmt.Print("Enter Id: ")
        idInput := readInput()

        template := xmlCreator.Template{
        	ID:             idInput,
        	Name:           nameInput,
        	PDP:            config.GetPDP(),
        	XSI:            config.GetXSI(),
        	SchemaLocation: config.GetSchemaLocation(),
        	Concept:        xmlCreator.Concept{},
        }

        modifyTemplate(template)
}

func modifyTemplate(template xmlCreator.Template) {
        fmt.Print("Concept version: ")
        conceptVersion := readInput()
        fmt.Print("Concept pageID: ")
        conceptLink := readInput()
        fmt.Print("Product version: ")
        productVersion := readInput()
        fmt.Print("Project version: ")
        projectVersion := readInput()

        concept := xmlCreator.Concept{
        	Value: conceptVersion,
                ConceptLink: struct{PageID string "xml:\"pageId,attr\""}{ 
                        PageID: conceptLink,
                },
                Tickets: []xmlCreator.Ticket{},
                Processes: []xmlCreator.Process{},
        	Version: struct{Product string "xml:\"product,attr\""; Project string "xml:\"project,attr\""}{
                        Product: productVersion,
                        Project: projectVersion,
                },
        }

        concept = modifyConcept(concept)
        template.Concept = concept
        xmlCreator.BuildFile(template)

}

func modifyConcept(concept xmlCreator.Concept) xmlCreator.Concept {

        exitCode := 0
        for exitCode != 1 {
                
                fmt.Println("Modifying: concept")
                fmt.Println("1. Add requirement ticket", "2. Add development ticket", "3. Add Process", "4. Add ParentProcess", "X. Finish")
                fmt.Print("Input: ")
                userInput := readInput()

                switch userInput {
                case "1":
                        concept = addTicket("Requirement", concept)
                case "2":
                        concept = addTicket("Task Development", concept)
                case "3":
                        concept = addProcess(concept)
                case "4":
                        concept = addParentProcess(concept)
                case "x", "X":
                        exitCode = 1
                }

        }

        return concept
}

func addProcess(concept xmlCreator.Concept) xmlCreator.Concept {
        fmt.Print("Process name: ")
        processName := readInput()
        fmt.Print("Process ID: ")
        processID := readInput()
        fmt.Print("Summary: ")
        summary := readInput()

        process := xmlCreator.Process{
        	ID:      processID,
        	Name:    processName,
        	Type:    "POS",
        	Summary: xmlCreator.Summary {
                        Value: summary,
                },
        }
        concept.Processes = append(concept.Processes, process)
        return concept
}


func addParentProcess(concept xmlCreator.Concept) xmlCreator.Concept {
        fmt.Print("Process name: ")
        processName := readInput()
        fmt.Print("Parent Process ID: ")
        processID := readInput()
        fmt.Print("Summary: ")
        summary := readInput()

        process := xmlCreator.Process{
        	PrntID:      processID,
        	Name:    processName,
        	Type:    "POS",
        	Summary: xmlCreator.Summary {
                        Value: summary,
                },
        }
        concept.Processes = append(concept.Processes, process)
        return concept
}

func addTicket(taskType string, concept xmlCreator.Concept) xmlCreator.Concept {
        fmt.Print("Task ID: ")
        taskID := readInput()
        ticket := xmlCreator.Ticket {
                Type: taskType,
                Value: taskID,
        }

        concept.Tickets = append(concept.Tickets, ticket)

        return concept
}
