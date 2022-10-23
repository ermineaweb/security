package infra

import (
	"fmt"
	"websec/src/domain/repository"
	"websec/src/domain/target"
)

type OtherProcessor struct {
	Target     target.Target
	Repository repository.Repository
}

func NewOtherProcessor(t target.Target, r repository.Repository) OtherProcessor {
	return OtherProcessor{Target: t, Repository: r}
}

func (p OtherProcessor) Process() {
	fmt.Println("OtherProcessor run process")
	p.Repository.Save("another result")
}
