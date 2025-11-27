package util

import (
	"fmt"
	"strings"

	"github.com/obscuren/mutan"
	backend "github.com/obscuren/mutan/backends"
)

func Compile(script string, silent bool) (ret []byte, err error) {
	if len(script) > 2 {
		line := strings.Split(script, "\n")[0]
		if len(line) > 1 && line[0:2] == "#!" {
			switch line {
			}
		} else {
			compiler := mutan.NewCompiler(backend.NewEthereumBackend())
			compiler.Silent = silent
			byteCode, errors := compiler.Compile(strings.NewReader(script))
			if len(errors) > 0 {
				var errs string
				for _, er := range errors {
					if er != nil {
						errs += er.Error()
					}
				}
				return nil, fmt.Errorf("%v", errs)
			}
			return byteCode, nil
		}
	}
	return nil, nil
}
