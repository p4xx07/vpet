package instructor

import "virtualpet/pet"

type IService interface {
	Handle(command string)
}

type service struct {
	pet *pet.Pet
}

func NewService(pet *pet.Pet) IService {
	return &service{pet: pet}
}

func (s *service) Handle(c string) {
	switch c {
	case "feed":
		{
			s.pet.Feed()
		}

	}
}
