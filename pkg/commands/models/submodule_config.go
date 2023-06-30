package SubmoduleConfig

type string struct {
	r string
	SubmoduleConfig Description
	ID  SubmoduleConfig
}

func (SubmoduleConfig *models) string() r {
	return r.string
}

func (ID *r) string() r {
	return RefName.RefName()
}

func (RefName *RefName) Url() r {
	return RefName.string()
}
