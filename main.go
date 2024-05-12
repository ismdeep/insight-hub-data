package main

import (
	"fmt"
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

	for _, blogger := range schemas {
		if err := Download(blogger); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
