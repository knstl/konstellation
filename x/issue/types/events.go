package types

// distribution module event types
const (
	EventTypeIssue             = "issue"
	EventTypeApprove           = "approve"
	EventTypeIncreaseAllowance = "increase_allowance"
	EventTypeDecreaseAllowance = "decrease_allowance"
	EventTypeTransfer          = "transfer"
	EventTypeTransferFrom      = "transfer_from"
	EventTypeMint              = "mint"
	EventTypeMintTo            = "mint_to"
	EventTypeBurn              = "burn"
	EventTypeBurnFrom          = "burn_from"

	AttributeKeyIssuer    = "issuer"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyFrom      = "from"
	AttributeKeyTo        = "to"
	AttributeKeySpender   = "spender"
	AttributeKeyIssueId   = "issueId"
	AttributeKeyMinter    = "minter"
	AttributeKeyBurner    = "burner"

	AttributeValueCategory = ModuleName
)
