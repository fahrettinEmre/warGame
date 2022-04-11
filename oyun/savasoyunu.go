package oyun

type Savascı struct {
	Name        string
	Can         int
	MermiSayısı int
}

type Dunyali struct {
	Savascı
	Birlik string
}

type Marslı struct {
	Savascı
	UzayGemisi string
}

type ISavascı interface {
	AtesEt()
	HasarAl()
}

func (d *Dunyali) AtesEt() {
	if d.MermiSayısı > 1 {
		d.MermiSayısı--
	}
}

func (m *Marslı) AtesEt() {
	if m.MermiSayısı > 1 {
		m.MermiSayısı--
	}

}
func (d *Dunyali) HasarAl() {
	if d.Can > 0 {
		d.Can -= 10
	}
}

func (m *Marslı) HasarAl() {
	if m.Can > 0 {
		m.Can -= 10
	}
}
