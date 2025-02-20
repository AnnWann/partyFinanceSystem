package terminal

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Promote() {
	if len(op.Arguments) < 2 {
		fmt.Println("Promover quem para o que? Use 'promote <id> + <função>'")
		return
	}

	id := op.Arguments[0]
	role := op.Arguments[1]

	err := executors.Promote(id, role)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Promovido com sucesso")
}