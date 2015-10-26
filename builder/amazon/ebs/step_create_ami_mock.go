package ebs

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepCreateAMIMock struct {
	AmiRegions []string
	AmiName    string
	Region     string
}

func (s *stepCreateAMIMock) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	// Create the image
	ui.Say(fmt.Sprintf("MOCK-STEP Creating the AMI: %s", config.AMIName))

	amis := map[string]string{
		s.Region: "ami-123456",
	}
	for i, r := range s.AmiRegions {
		amis[r] = fmt.Sprintf("ami-11111%d", i)
	}
	state.Put("amis", amis)

	return multistep.ActionContinue
}

func (s *stepCreateAMIMock) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	ui.Say(fmt.Sprintf("MOCK-STEP Cleanup(): nothing todo ..."))
}
