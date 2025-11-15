package bot

type Bot struct {
	Token string
}

func NewBot(token string) *Bot {
	return &Bot{
		Token: token,
	}
}
