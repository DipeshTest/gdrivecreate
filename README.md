---
title: Upload a file to GDrive
weight: 1
---

# Upload a file to GDrive
This activity allows you to upload a file to your Grdive account and optionally add a user to share the file with. An email will be sent to the user with whom the file is being shared if the email address is valid and sendNotification is set to true.

## Installation
### Flogo Web
This activity is built by Team AllStars
### Flogo CLI
```bash
flogo install github.com/DipeshTest/gdrivecreate
```

### Third-party libraries used
- #### package drive - "google.golang.org/api/drive/v3":
	Package drive provides access to the Drive API. For more details, check https://developers.google.com/drive/
- #### package googleapi - "google.golang.org/api/googleapi":
	Package googleapi contains the common code shared by all Google API libraries.

## Schema
Inputs and Outputs:

```json
 "inputs":[
    {
		"name": "accessToken",
		"type": "string",
		"required": true
	},
	{
		"name": "fileFullPath",
		"type": "string"

	},
	{
		"name": "emailAddr",
		"type": "string"

	},
	{
		"name": "sendNotification",
		"type": "boolean"

	},
	{
		"name": "role",
		"type": "string",
		"allowed": [
			"reader",
			"writer",
			"organizer",
			"owner",
			"recency",
			"commenter"
		  ],
		"value": "writer",
		"required": true
	},
	{
		"name": "timeout",
		"type": "string",
		"value": "120"
	}
  ],
  "outputs": [
    {
      "name": "statusCode",
      "type": "string"
    },
    {
      "name": "message",
      "type": "any"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| accessToken | True     | The access token for your account |         
| fileFullPath   | True    | Fullpath of the file you want to upload to GDrive|
| emailAddr       | False    | Optional Email address to share the file with, for now we have given the option to share file with 1 user at a time|
| sendNotification   | False    | Set to true if you want to send an email to the user menitoned in the emailAddr field once the file is uploaded and shared|
| role   | True    | The permissions you want to give the user you are sharing file with|
| timeout   | False    | Timeout for the activity, default is 120 seconds|
## Examples
Please refer activity_test.go 


## Response Codes
### Google Drive Create
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request was successful and the response body contains the representation requested. *If the email is specified and role is not selected, then the file will be uploaded & no email would be sent|
|101 |INVALID INPUT| The file path specified is invalid.|
|102 |INVALID INPUT| The specified destination is a folder.|
|105 |INVALID INPUT| Access token is not specified.|
|106 |INVALID INPUT| File path is not specified.|
|401 |AUTHENTICATION ERROR| Invalid Access Token.|
