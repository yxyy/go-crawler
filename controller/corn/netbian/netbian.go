package netbian

import "fmt"

type Netbian struct {
	Url string

}

func (this *Netbian) Run()  {
	fmt.Println("开始爬虫")
}