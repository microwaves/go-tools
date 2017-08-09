package kubectl

import "os/exec"

type Kubectl struct {
	bin  string
	args []string
}

func NewKubectl(args []string) (kubectl *Kubectl, err error) {
	kubectlBinPath, err := WhereIs()
	kubectl = &Kubectl{bin: kubectlBinPath, args: args}
	return
}

func WhereIs() (kubectlBinPath string, err error) {
	kubectlBinPath, err = exec.LookPath("kubectl")
	return
}

func (k *Kubectl) Exec() (string, error) {
	command := exec.Command(k.bin, k.args...)
	commandOutput, err := command.Output()
	return string(commandOutput), err
}
