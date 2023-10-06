package runner

import "sync"

type shutdown struct {
	*sync.Once

	err error
}
