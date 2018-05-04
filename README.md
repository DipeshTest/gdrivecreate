---
title: Upload a file to GDrive
weight: 1
---

# Counter
This activity allows you to upload a file to your Grdive account and optionally add a user to share the file with, an email will be sent to the user with whom the file is being shared if the user account the file is being shared with is a valid one and sendNotification is set as true

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo add activity github.com/DipeshTest/gdrivecreate
```

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
| timeout   | False    | Timeout for the activity, default is set to 120 seconds|
## Examples
### Increment
The below example for a sample create:

```json

```