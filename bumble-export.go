package main

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"strings"
	Common "bumble-export/common"
)

func main() {

	// Check inputs
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s /path/to/file.har /path/to/export/dir", os.Args[0])
		return
	}

	var HarFile = os.Args[1]
	var ExportDir = os.Args[2]

	if Common.IsFileOrDir(HarFile) != "file" {
		log.Panicf("Error: %s is not a file", HarFile)
	}

	if Common.IsFileOrDir(ExportDir) != "dir" {
		log.Panicf("Error: %s is not a directory", ExportDir)
	}

	// Open HAR file 
	file, err := os.Open(HarFile)
	if err != nil {
	
		log.Panic("Error while opening file: ", err)
	}
	defer file.Close()

	// Parse HAR file to JSON
	var HarJson Common.HarStruct
	harParser := json.NewDecoder(file)
	err = harParser.Decode(&HarJson)
	if err != nil {
		log.Panic("Error while parsing HAR file into JSON: ", err)
	}

	// Create final export JSON 
	var ExportJson []Common.ExportStruct

	// Set UserID and UserName for all files
	var UserID string
	var UserName string
	
	// Loop over entries array, messsages order will be backwards, but we will reverse it when writing to file
	var HarEntries = HarJson.Log.Entries
	for entriesCounter := 0; entriesCounter < len(HarEntries); entriesCounter++ {

		// Remove backslash escape characters from nested JSON
		var HarBumbleText = HarEntries[entriesCounter].Response.Content.Text
    	var BumbleJsonByte = []byte(strings.ReplaceAll(HarBumbleText, "\\", ""))
    	

		// Check whether API request was for SERVER_GET_CHAT_MESSAGES or SERVER_OPEN_CHAT
		// SERVER_OPEN_CHAT is the first request when the chat is opened, every another request for chat messages is SERVER_GET_CHAT_MESSAGES
		var BumbleOpenChatJson Common.BumbleOpenChatStruct
		var BumbleGetChatJson Common.BumbleGetChatStruct		
		var RequestUrl = HarEntries[entriesCounter].Request.Url

		if strings.Contains(RequestUrl, "SERVER_OPEN_CHAT") {
			err := json.Unmarshal(BumbleJsonByte, &BumbleOpenChatJson)
			if err != nil {
				log.Panic("Error while parsing Bumble JSON into struct: ", err)
			}

			UserID = BumbleOpenChatJson.Body[0].ClientOpenChat.ChatUser.UserID
			UserName = BumbleOpenChatJson.Body[0].ClientOpenChat.ChatUser.Name
			
			// Count backwards so we get messages from the newest to the oldest ones 
			// Messages in SERVER_OPEN_CHAT request are in a reversed order compared to SERVER_GET_CHAT_MESSAGES (the oldest one is index 0)
			var BumbleOpenChatMessages = BumbleOpenChatJson.Body[0].ClientOpenChat.ChatMessages
			for messagesCounter := len(BumbleOpenChatMessages) - 1; messagesCounter >= 0; messagesCounter-- {
				message := Common.BumbleToExportJson(UserID, UserName, BumbleOpenChatMessages[messagesCounter], ExportDir)
				ExportJson = append(ExportJson, message)
			} 
			
		} else if strings.Contains(RequestUrl, "SERVER_GET_CHAT_MESSAGES") {
			err := json.Unmarshal(BumbleJsonByte, &BumbleGetChatJson)
			if err != nil {
				log.Panic("Error while parsing Bumble JSON into struct: ", err)
			}

			// Count normally so we get messages from the newest to the oldest ones 
			// For SERVER_GET_CHAT_MESSAGES the newest message is index 0, so opposite than for SERVER_OPEN_CHAT
			var BumbleGetChatMessages = BumbleGetChatJson.Body[0].ClientChatMessages.Messages
			for messagesCounter := 0; messagesCounter < len(BumbleGetChatMessages); messagesCounter++ {
				message := Common.BumbleToExportJson(UserID, UserName, BumbleGetChatMessages[messagesCounter], ExportDir)
				ExportJson = append(ExportJson, message)
			} 
		} else {
			log.Print("Error: unknown API request: ", RequestUrl)
		}
	}
	// Reverse order in array so messages are sorted from the oldest to the newest
	var ReverseExportJson []Common.ExportStruct
	for counter := len(ExportJson) - 1; counter >= 0; counter-- {
		ReverseExportJson = append(ReverseExportJson, ExportJson[counter])
	}

	// Parse JSON back to text
    ExportJsonString, err := json.MarshalIndent(ReverseExportJson, "", "\t")
    if err != nil {
        log.Print(err)
    }
	
	// Create file
	ExportJsonOutFile, err := os.Create(ExportDir + "/messages.json")
	if err != nil {
		log.Panicf("Error while creating output file: %s", err)
	}	
	defer ExportJsonOutFile.Close()

	// Write data to file
	_, err = ExportJsonOutFile.WriteString(string(ExportJsonString))
	if err != nil {
		log.Panicf("Error while saving to file: %s", err)
	}
}
