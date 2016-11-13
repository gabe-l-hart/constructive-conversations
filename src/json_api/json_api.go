package json_api

// Internal ////////////////////////////////////////////////////////////////////

type AccountPublic struct {
	FirstName   string   `json:"first_name"`
	LastInitial string   `json:"last_initial"`
	Identities  []string `json:"identities"`
}

type AccountPrivate struct {
	ChatCreds      string `json:"chat_creds"`
	Disabled       bool   `json:"disabled"`
	InConversation bool   `json:"in_conversation"`
}

// API Messages ////////////////////////////////////////////////////////////////

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type AccountRequest struct {
	Key  string        `json:"key"`
	Data AccountPublic `json:"account"`
}

type AccountResponse struct {
	Public  AccountPublic  `json:"public"`
	Private AccountPrivate `json:"private"`
}
