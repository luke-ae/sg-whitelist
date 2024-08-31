package main

type Error struct {
	Error string
}

type Queries struct {
	Minter *MinterQuery `json:"minter,omitempty"`
	Config *ConfigQuery `json:"config,omitempty"`
}

type MinterAddr struct {
	MinterAddr string `json:"minterAddr,omitempty"`
}

type ConfigQuery struct {
}

type MinterQuery struct {
}

type MemberResponse struct {
	Minter     string `json:"minterAddr"`
	Collection string `json:"collectionAddr"`
	HasMember  bool   `json:"hasMember"`
	Name       string `json:"name"`
	Image      string `json:"image"`
}

type WhitelistQueries struct {
	HasMember *WhitelistMemberQuery `json:"has_member,omitempty"`
}

type WhitelistMemberQuery struct {
	Member string `json:"member"`
}

type WhitelistHasMemberQueryResponse struct {
	Data *WhitelistHasMember `json:"data,omitempty"`
}

type WhitelistHasMember struct {
	HasMember bool `json:"has_member"`
}
