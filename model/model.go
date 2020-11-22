package model

type(
	Re struct {
		Status string
		Msg string
		Data interface{}
	}

	Node struct {
		Token string
		Name string
		IP string
		Type string
		UserToken string
	}
	User struct {
		ChatID string
		Name string
		NickName string
		Token string
		CreatedTime string
		UpdatedTime string
		Enable bool
		Others map[string]interface{}
	}
)
