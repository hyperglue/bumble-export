# Bumble export tool
## Usage
1. Open Bumble web app
2. Open a chat you want to export
3. Open developer tools then go to Network tab
4. Search for "CHAT" and select "XHR" type
5. Refresh page so the SERVER_OPEN_CHAT request is sent again
6. Scroll up until you reach the beginning of the chat (
    * (in Firefox you can set "general.autoScroll" to "true" in about:config, then press middle mouse button and move cursor to the upper edge of screen)
8. Verify that you only see SERVER_OPEN_CHAT and SERVER_GET_CHAT_MESSAGES requests in the list
9. Save all requests to HAR file and then use it as input to this script:
    * In Firefox click on small gear icon in the top-right corner of developer tools pane, then select "Save All As HAR" 
    * In Chromium click on small download icon in the upper bar of developer tools pane
10. Run ```bumble-export /path/to/exported.har /path/to/export/dir```

## How to compile
1. Grab this repo: ```git clone```
2. Run ```go build bumble-export.go```

## Output format
```
[
	{
		"uid": "123",
		"date": 1695600000,
		"from": "Someone",
		"to": "Me",
		"type": "text",
		"content": "Sample message"
	},
	{
		"uid": "456",
		"date": 1695600010,
		"from": "Me",
		"to": "Someone",
		"type": "photo",
		"content": "456.jpg"
	},
	{
		"uid": "789",
		"date": 1695600020,
		"from": "Someone",
		"to": "Me",
		"type": "audio",
		"content": "789.mp4"
	}
]
```
