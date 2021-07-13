package validate

import (
	"errors"

	"github.com/deqodelabs/IaaC/stackconfig/pb"
	"github.com/deqodelabs/IaaC/stackconfig/validate/forwardedPortCheck"
	"github.com/deqodelabs/IaaC/stackconfig/validate/resourceCheck"
	"github.com/deqodelabs/IaaC/stackconfig/validate/runtimeCheck"
)

func CustomValidate(app *pb.StackConfig) error {
	if app == nil {
		return errors.New("app is nil")
	}
	for _, v := range app.Services {
		if v.Network != nil {
			err := forwardedPortCheck.CustomForwardedPortValidation(v.Network.ForwardedPorts)
			if err != nil {
				return err
			}
		}
		err := resourceCheck.CustomResourceValidation(v.Resources)
		if err != nil {
			return err
		}
		err = runtimeCheck.CustomValidateRuntime(v)
		if err != nil {
			return err
		}
	}
	return nil
}
