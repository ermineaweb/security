package infra

import (
	"fmt"
	"io"
	"net/http"
	"websec/src/domain/repository"
	"websec/src/domain/target"
)

type RobotsTxt struct {
	Target     target.Target
	Repository repository.Repository
}

func NewRobotsTxt(t target.Target, r repository.Repository) RobotsTxt {
	return RobotsTxt{Target: t, Repository: r}
}

func (p RobotsTxt) Process() {
	fmt.Println("RobotTxt run process", p.Target.Url+"/robots.txt")

	res, err := http.Get(p.Target.Url + "/robots.txt")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	p.Repository.Save(string(body))
}
