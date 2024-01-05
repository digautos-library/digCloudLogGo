package main

import (
	"fmt"

	digcloudlog "github.com/digautos-library/digCloudLogGo"
)

func main() {

	digcloudlog.DCL_addStdout()
	err := digcloudlog.DCL_addLocalFileDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = digcloudlog.DCL_addLogflare("d0399bff-7fc7-4874-a572-05309021d853", "WoMn49mFQDkh")
	if err != nil {
		fmt.Println(err)
		return
	}
	digcloudlog.DCL_Info("hello info")
	digcloudlog.DCL_Error("hello error")
}
