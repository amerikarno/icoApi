package risks

type ChannelRisk struct {
	channel      string
	ChannelPoint int
}

func NewChannelRiskUsecase(channel string) *ChannelRisk {
	return &ChannelRisk{channel: channel}
}

func (r *ChannelRisk) GetChannelRisk() *ChannelRisk {
	if r.channel == ndid {
		return r
	}
	return r
}
