package entities

type WalletData struct {
	Data []AllocationData `json:"claimables"`
}

type AllocationData struct {
	Slug   string      `json:"org_slug"`
	Amount string      `json:"amount"`
	Phases []PhaseData `json:"phases"`
}

type PhaseData struct {
	ID          string `json:"id"`
	PhaseNumber int    `json:"phase"`
	Amount      string `json:"amount"`
}
