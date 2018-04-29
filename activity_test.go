package gdrivecreate

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

 /*func TestGDriveCreateFile_Success(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileFullPath", "C:\\Users\\asutar\\Documents\\CodeOfConduct.pdf")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")

}*/

/* func TestGDriveCreateFile_TokenError(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GluqBdPuqRNrY9cJ3elwFXSu_4Mx5leuMYXmrdvIB_MPMdU0Y-6M9iU83ELUEJo5UAqeQTN8EaDdUkMxXqhk11E9Ed8H3Cgh353dC05fr3efeKiDebR0Gf11fM7t")
	tc.SetInput("fileFullPath", "C:\\Users\\asutar\\Documents\\CodeOfConduct.pdf")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "401")
	tc.SetInput("accessToken", "")

	act.Eval(tc)

	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "105")

} */

/* func TestGDriveCreateFile_InvalidPath(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileFullPath", "")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "106")




	tc.SetInput("fileFullPath", "D://MYGsfdkdsf/jdkd.pdf")
	act.Eval(tc)
	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "101")


	tc.SetInput("fileFullPath", "D:\\Programs")


	act.Eval(tc)
	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "102")

} */


func TestGDriveCreateFile_InvalidUser(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileFullPath", "C:\\Users\\asutar\\Documents\\CodeOfConduct.pdf")
	tc.SetInput("emailAddr", "1234tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")






}
