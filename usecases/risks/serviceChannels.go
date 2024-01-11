package risks

type ServiceChannel struct {
	Channel string
}

func NewServiceChannelsUsecases() *ServiceChannel {
	return &ServiceChannel{}
}

func (c *ServiceChannel) ChannelUsecase() int {
	switch c.Channel {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	default:
		return 0
	}
}
