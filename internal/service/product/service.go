package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() ([]Product, error) {
	return allProduct, nil
}

func (s *Service) GetByOffset(offset int) (*Product, error) {
	// if offset < 0 || offset > len(allProduct) {
	// 	return nil, errors.New("Invalid offset")
	// }

	return &allProduct[offset], nil
}
