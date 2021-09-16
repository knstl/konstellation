package types

// distribution module event types
const (
	EventTypeIssue             = "issue"
	EventTypeApprove           = "approve"
	EventTypeIncreaseAllowance = "increase_allowance"
	EventTypeDecreaseAllowance = "decrease_allowance"
	EventTypeTransfer          = "transfer"
	EventTypeTransferFrom      = "transfer_from"
	EventTypeTransferOwnership = "transfer_ownership"
	EventTypeMint              = "mint"
	EventTypeBurn              = "burn"
	EventTypeBurnFrom          = "burn_from"
	EventTypeFreeze            = "freeze"
	EventTypeUnfreeze          = "unfreeze"
	EventTypeChangeFeatures    = "features"
	EventTypeChangeDescription = "description"

	AttributeKeyIssuer      = "issuer"
	AttributeKeyRecipient   = "recipient"
	AttributeKeyOwner       = "owner"
	AttributeKeyFreezer     = "freezer"
	AttributeKeyHolder      = "holder"
	AttributeKeyMinter      = "minter"
	AttributeKeyBurner      = "burner"
	AttributeKeySpender     = "spender"
	AttributeKeyFrom        = "from"
	AttributeKeyTo          = "to"
	AttributeKeyOp          = "op"
	AttributeKeyDenom       = "denom"
	AttributeKeyFeatures    = "features"
	AttributeKeyDescription = "description"

	AttributeValueCategory = ModuleName
)
