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
	Name       string  `json:"name"`
	Collection string  `json:"collectionAddr"`
	Minter     string  `json:"minterAddr"`
	Image      string  `json:"image"`
	MintedAt   *string `json:"mintedAt"`
	HasMember  bool    `json:"hasMember"`
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

type Health struct {
	Status string `json:"status"`
}
