// GENERATED BY goruda
// This file was generated automatically at
// 2020-07-17 07:53:27.966507 +0700 WIB m=+0.013592107

package articletest

type ChildOfTag struct {
	String string `json:"string"`
	Topic  `json:""`
}

func (p ChildOfTag) Tostring() string {
	return p.String
}

func (p ChildOfTag) ToTopic() Topic {
	return p.Topic
}
