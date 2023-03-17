package model

type MailAttachment struct {
	RewardType  int32  `json:"type"`
	RewardValue string `json:"value"`
	RewardCount int32  `json:"count"`
}
