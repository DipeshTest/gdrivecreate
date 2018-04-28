package gdrivecreate

import (
	"strings"

	"github.com/DipeshTest/allstarsshared/GDrive"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-gdrivecreate")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	accessToken := context.GetInput("accessToken").(string)
	fileFullPath := context.GetInput("fileFullPath").(string)
	emailAddr := context.GetInput("emailAddr").(string)
	sendNotification := context.GetInput("sendNotification").(bool)
	role := context.GetInput("role").(string)

	if len(strings.TrimSpace(accessToken)) == 0 {

		context.SetOutput("statusCode", "105")

		context.SetOutput("message", "Access Token field is blank")
		//context.SetOutput("failedNumbers", to)

		//respond with this
	} else if len(strings.TrimSpace(fileFullPath)) == 0 {

		context.SetOutput("statusCode", "106")

		context.SetOutput("message", "File Path field is blank")

	} else {

		code, msg := GDrive.CreateFile(accessToken, fileFullPath, emailAddr, role, sendNotification)
		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
