package types

// distribution module event types
const (
	EventTypeIssue    = "issue"
	EventTypeApprove  = "approve"
	EventTypeTransfer = "transfer"

	AttributeKeyIssuer    = "issuer"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeySpender   = "spender"
	AttributeKeyIssueId   = "issueId"

	AttributeValueCategory = ModuleName
)
