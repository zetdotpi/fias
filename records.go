package fias

type Record struct {
	ID string
}

const (
	REC_LinkAlive             = "LA"
	REC_LinkStart             = "LS"
	REC_LinkEnd               = "LE"
	REC_LinkDescription       = "LD"
	REC_LinkRecord            = "LR"
	REC_DatabaseResyncRequest = "DR"
	REC_DatabaseResyncStart   = "DS"
	REC_DatabaseResyncEnd     = "DE"
	REC_NightAuditStart       = "NS"
	REC_NightAuditEnd         = "NE"
	REC_GuestCheckin          = "GI"
	REC_GuestCheckout         = "GO"
	REC_GuestDataChange       = "GC"
)
