package types

// distribution module event types
const (
	EventTypeIssue             = "issue"
	EventTypeApprove           = "approve"
	EventTypeIncreaseAllowance = "increase_allowance"
	EventTypeDecreaseAllowance = "decrease_allowance"
	EventTypeTransfer          = "transfer"
	EventTypeTransferFrom      = "transfer_from"

	AttributeKeyIssuer    = "issuer"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeySpender   = "spender"
	AttributeKeyIssueId   = "issueId"

	AttributeValueCategory = ModuleName
)
