package star

type Star struct {
	PositionCode int
	Size         string
	Type         string
	SubType      string
	IsMapped     bool
	Age          *float64 //Gyrs
	Comment      *string
}

func New(knownData ...StarOption) *Star {
	return &Star{}
}

type StarOption func(*Star)

func Size(size string) StarOption {
	return func(s *Star) {
		s.Size = size
	}
}
