package kubectl

import "os/exec"

type Kubectl struct {
	bin string
	args []string
}

func New(b string, a []string) *Kubectl {
	return &Kubectl{bin: b, args: a}
}

func (k *Kubectl) Exec() (string, error) {
	c := exec.Command(k.bin, k.args...)
	o, err := c.Output()
	return string(o), err
}

func WhereIs() (k string, err error) {
	k, err = exec.LookPath("kubectl")
	return
}
