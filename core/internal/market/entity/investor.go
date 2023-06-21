package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPossition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPossition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtShares int) {
	assetPosition := i.GetAssetPossition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtShares))
	} else {
		assetPosition.Shares += qtShares
	}
}

func (i *Investor) GetAssetPossition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}
