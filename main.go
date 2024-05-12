package main

import (
	"fmt"
	"sync"
)

func main() {

	schemas := []BloggerInterface{
		&Antonz{},
		&Akarin{},
		&ArthurChiaoArt{},
		&B303248153{},
		&BogomolovTech{},
		&BouKe{},
		&CoolShell{},
		&DunWu{},
		&EltonMinetto{},
		&Evilsocket{},
		&GoDev{},
		&MeituanTech{},
		&WhyDegree{},
	}

	var wg sync.WaitGroup
	var errLst []error
	for _, blogger := range schemas {
		wg.Add(1)
		go func(blogger BloggerInterface) {
			defer wg.Done()
			if err := Download(blogger); err != nil {
				errLst = append(errLst, fmt.Errorf("failed to download blogger %v: %s", blogger.GetBloggerName(), err.Error()))
			}
			fmt.Println("FINISHED:", blogger.GetBloggerName())
		}(blogger)
	}
	wg.Wait()
	if len(errLst) > 0 {
		for _, e := range errLst {
			fmt.Println(e.Error())
		}
	}
}
