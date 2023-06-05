package core

type System struct {
	Modules []*Module
}

func (s *System) Walk(v Visitor) error {
	for _, m := range s.Modules {
		if err := v.Visit(m); err != nil {
			return err
		}
	}
	return nil
}
