package common

type HarStruct struct {
	Log struct {
	Entries []struct {
	Request struct {
		Url string `json:"url"`
	} `json:"request"`
		
	Response struct {
		Content struct {
			Text string `json:"text"`
		} `json:"content"`
	} `json:"response"`
	} `json:"entries"`
	} `json:"log"`
}

type BumbleOpenChatStruct struct {
	Body []struct {
		ClientOpenChat struct {
			ChatMessages []BumbleMessagesStruct `json:"chat_messages"`
			ChatUser struct {
				UserID string `json:"user_id"`
				Name string `json:"name"`	
			} `json:"chat_user"`
		} `json:"client_open_chat"`
	} `json:"body"`
}

type BumbleGetChatStruct struct {
	Body [] struct {
		ClientChatMessages struct {
			Messages []BumbleMessagesStruct `json:"messages"`
		} `json:"client_chat_messages"`
	}
}

type BumbleMessagesStruct struct {
	UID string `json:"uid"`
	Date int `json:"date_modified"`
	FromPersonID string `json:"from_person_id"`
	ToPersonID string `json:"to_person_id"`
	Message string `json:"mssg"`
	Multimedia struct {
		Photo struct {
			Url string `json:"large_url"`
		} `json:"photo"`
		Audio struct {
			Url string `json:"url"`
		} `json:"audio"`
	} `json:"multimedia"`
}
