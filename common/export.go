package common

import "log"
import "strings"

type ExportStruct struct {
	UID string `json:"uid"`
	Date int `json:"date"`
	From string `json:"from"`
	To string `json:"to"`
	Type string `json:"type"`
	Content string `json:"content"`
}

var MyName = "Me"

// Parse from Bumble to Export JSON
func BumbleToExportJson(userID string, userName string, bumble BumbleMessagesStruct, path string) (export ExportStruct) {

	message := ExportStruct {
		UID: bumble.UID,
		Date: bumble.Date,
	}

	// Change IDs to user friendly names
	if bumble.FromPersonID == userID {
		message.From = userName
		message.To = MyName
	} else if bumble.ToPersonID == userID {
		message.From = MyName
		message.To = userName
	}
	
	if bumble.Multimedia.Photo.Url != "" {
		var filename = bumble.UID + ".jpg"
		var filepath = path + "/" + filename

		// Here we add https: to URL because API response lacks it for some reason
		err := DownloadFile(filepath, "https:" + bumble.Multimedia.Photo.Url)
		if err != nil {
			log.Print(err)
		}
		message.Type = "photo"
		message.Content = filename
		
	} else if bumble.Multimedia.Audio.Url != "" {
		var filename = bumble.UID + ".mp4"
		var filepath = path + "/" + filename
		
		err := DownloadFile(filepath, bumble.Multimedia.Audio.Url)
		if err != nil {
			log.Print(err)
		}
		message.Type = "audio"
		message.Content = filename
		
	} else {
		// Fix rendering apostrophe as \u0026#039;
		var apostropheFix = strings.ReplaceAll(bumble.Message, "\u0026#039;", "'")
		message.Type = "text"
		message.Content = apostropheFix
	}

	return message
}
