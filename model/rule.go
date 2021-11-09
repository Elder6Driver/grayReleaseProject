package model

type Rule struct {
	// rule condition
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
	MinUserDID int `json:"min_user_did"`
	MaxUserDID int `json:"max_user_did"`

	// apk or ipa file link
	GrayLink string `json:"gray_link"`
}

func GetAllRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinVersion: 10,
		MaxVersion: 20,
		MinUserDID: 10,
		MaxUserDID: 20,
		GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
	})

	rules = append(rules, Rule{
		MinVersion: 0,
		MaxVersion: 9,
		MinUserDID: 0,
		MaxUserDID: 9,
		GrayLink:   "https://dldir1.qq.com/weixin/android/weixin8015android2020_arm64.apk",
	})

	return &rules

}
