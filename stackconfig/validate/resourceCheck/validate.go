package resourceCheck

import (
	"github.com/deqodelabs/IaaC/stackconfig/pb"

	"errors"
)

func CustomResourceValidation(resource *pb.Resource) error {
	if resource != nil {
		if resource.Cpu != 1 && (resource.Cpu%2) != 0 {
			return errors.New("cpu number not valid")
		}
	}
	return nil
}
