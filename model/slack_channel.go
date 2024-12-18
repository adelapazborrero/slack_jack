package model

type ChannelList struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	IsChannel               bool     `json:"is_channel"`
	IsGroup                 bool     `json:"is_group"`
	IsIM                    bool     `json:"is_im"`
	IsMPIM                  bool     `json:"is_mpim"`
	IsPrivate               bool     `json:"is_private"`
	Created                 int64    `json:"created"`
	IsArchived              bool     `json:"is_archived"`
	IsGeneral               bool     `json:"is_general"`
	Unlinked                int      `json:"unlinked"`
	NameNormalized          string   `json:"name_normalized"`
	IsShared                bool     `json:"is_shared"`
	IsOrgShared             bool     `json:"is_org_shared"`
	IsPendingExtShared      bool     `json:"is_pending_ext_shared"`
	PendingShared           []string `json:"pending_shared"`
	ContextTeamID           string   `json:"context_team_id"`
	Updated                 int64    `json:"updated"`
	ParentConversation      *string  `json:"parent_conversation"`
	Creator                 string   `json:"creator"`
	IsExtShared             bool     `json:"is_ext_shared"`
	SharedTeamIDs           []string `json:"shared_team_ids"`
	PendingConnectedTeamIDs []string `json:"pending_connected_team_ids"`
	IsMember                bool     `json:"is_member"`
	Topic                   Topic    `json:"topic"`
	Purpose                 Purpose  `json:"purpose"`
	Properties              Property `json:"properties"`
	PreviousNames           []string `json:"previous_names"`
	NumMembers              int      `json:"num_members"`
}

type Topic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
}

type Purpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int64  `json:"last_set"`
}

type Property struct {
	UseCase string `json:"use_case"`
}
