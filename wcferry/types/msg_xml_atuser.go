package types

// 群里At用户消息
// @msg.type 1

type MsgXmlAtUser struct {
	AtUserList  string `xml:"atuserlist"`
	Silence     int32  `xml:"silence"`
	MemberCount int32  `xml:"membercount"`
	Signature   string `xml:"signature"`
	TmpNode     struct {
		PublisherID string `xml:",chardata"`
	} `xml:"tmp_node"`
}
type AvatarPayload struct {
	// 用户 id
	UsrName string `json:"usr_name,omitempty"`
	// 大头像 url
	BigHeadImgUrl string `json:"big_head_img_url,omitempty"`
	// 小头像 url
	SmallHeadImgUrl string `json:"small_head_img_url,omitempty"`
}
