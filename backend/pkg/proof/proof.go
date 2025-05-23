package proof

import (
	"fmt"
	"molo-cli/backend/pkg/conf"

	"github.com/go-omnibus/proof"
)

func MustInitProof() {
	p := proof.New()
	p.SetDivision(proof.TimeDivision)
	p.SetTimeUnit(proof.Day)
	p.SetEncoding(proof.ConsoleEncoder)
	//p.SetCaller(true)
	p.SetCapitalColor(true)
	p.SetInfoFile(conf.Conf.Proof.InfoLog)
	p.SetErrorFile(conf.Conf.Proof.ErrLog)

	if !conf.Conf.Base.IsDebug {
		p.CloseConsoleDisplay()
	}
	fmt.Println("is_debug:", conf.Conf.Base.IsDebug)

	p.Run()
	defer func() {
		err := proof.Sync()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("proof initialized")
}
