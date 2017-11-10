package mypkg

type TVChannel struct {
	id         string
	name       string
	programmes []*TVProgramme
}

func (c *TVChannel) GetId() string {
	return c.id
}

func (c *TVChannel) GetName() string {
	return c.name
}

func (c *TVChannel) GetProgrammes() []*TVProgramme {
	return c.programmes
}

func (c *TVChannel) add_programme(p *TVProgramme) {
	c.programmes = append(c.programmes, p)
}

func new_channel(id string, name string) *TVChannel {
	c := new(TVChannel)
	c.id = id
	c.name = name
	c.programmes = make(*TVProgramme, 0)
	return c
}
