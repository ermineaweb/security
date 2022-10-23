package main

import (
	"websec/src/domain/target"
	"websec/src/infra"
)

func main() {
	target := target.Target{Name: "www.google.fr", Url: "https://www.google.fr"}

	robottxt := infra.NewRobotsTxt(target, infra.NewFileRepository("robots.txt", target.Name))
	target.Processors = append(target.Processors, robottxt)

	otherProcess := infra.NewOtherProcessor(target, infra.NewFileRepository("other.txt", target.Name))
	target.Processors = append(target.Processors, otherProcess)

	for _, processor := range target.Processors {
		processor.Process()
	}
}
