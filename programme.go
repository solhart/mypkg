package mypkg

type TVProgramme struct {
	start       string
	stop        string
	title       string
	description string
}

func (p *TVProgramme) GetStart() string {
	return p.start
}

func (p *TVProgramme) GetStop() string {
	return p.stop
}

func (p *TVProgramme) GetTitle() string {
	return p.title
}

func (p *TVProgramme) GetDescription() string {

}

func new_programme(start string, stop string, title string, desc string) *TVProgramme {
	p := new(TVProgramme)
	p.start = start
	p.stop = stop
	p.title = title
	p.description = desc
	return p
}
