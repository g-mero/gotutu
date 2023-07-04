package api

type Group struct {
	TestApi
	UploadApi
	ConfApi
}

var GroupApp = new(Group)
