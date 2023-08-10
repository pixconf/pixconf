package protos

type SecretAlias struct {
	ACLCreate bool `json:"acl_create"`
	ACLUpdate bool `json:"acl_update"`
	ACLDelete bool `json:"acl_delete"`
}
