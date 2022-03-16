package main

import "fmt"

type Action func() error

func SafeExec(a Action) Action {
	err := a
	if err != nil {
		return nil, fmt.Errorf("Safe exec : %w",err())
	}
}