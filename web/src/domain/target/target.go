package target

import "websec/src/domain/processors"

type Target struct {
	Name       string
	Url        string
	Processors []processors.Processor
}
