package googlecompute

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepCreateImageMock struct {
}

func (s *StepCreateImageMock) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Mocking: Creating image...")

	state.Put("image_name", config.ImageName)
	return multistep.ActionContinue
}

// Cleanup.
func (s *StepCreateImageMock) Cleanup(state multistep.StateBag) {}
