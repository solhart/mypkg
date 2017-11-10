package mypkg

type TVListing struct {
	channels map[string]*TVChannel
}

func (l *TVListing) GetChannels() []*TVChannel {
	return l.channels
}

func (l *TVListing) put_channel(c *TVChannel) {
	l.channels[c.GetName()] = c
}

func (l *TVListing) GetChannel(name string) *TVChannel {
	return l.channels[name]
}

func new_listings() *TVListing {
	l := new(TVListing)
	l.channels = make(map[string]*TVChannel)
	return l
}
