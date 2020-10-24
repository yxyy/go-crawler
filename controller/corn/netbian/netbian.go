package netbian

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"lhc.go.crawler/logs"
	"net/http"
	"github.com/axgle/mahonia"
)

type Netbian struct {
	Url string

}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func (this Netbian) Run()  {
	url := "http://pic.netbian.com/4kdongman/"
	resp, err := http.Get(url)
	if err!=nil {
		logs.Error.Info(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		fmt.Println("爬取完成")
		return
	}
	//bytes, err := ioutil.ReadAll(resp.Body)
	//if err!=nil {
	//	logs.Error.Info(err)
	//}
	//
	//fmt.Println(string(bytes))

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		fmt.Println(err)
	}
	//text := document.Find(".col4-t2>.item-img-cont").Text()
	//fmt.Println(text)
	document.Find(".slist>ul>li").Each(func(i int, selection *goquery.Selection) {
		img_alt, _ := selection.Find("a>img").Attr("alt")
		img_url, _ := selection.Find("a>img").Attr("src")
		fmt.Println(gbktoutf8(img_alt),img_url)
	})


}

func gbktoutf8(text string) string {
	return mahonia.NewDecoder("gbk").ConvertString(text)
}

func NewNetbian() *Netbian {
	return &Netbian{}
}
//func (this *Netbian) Run()  {
//	fmt.Println("开始爬虫")
//
//	category := model.NewCategory()
//
//	category.Mark=viper.GetString("web.netbian")
//	if err := category.GetOneCategoryByMark();err!=nil{
//		logs.Corn.WithField("web","netbian").Info(err)
//	}
//	categories, err := category.GetChilendCategoryById()
//	if err!=nil{
//		logs.Corn.WithField("web","netbian").Info(err)
//	}
//
//	for _,v := range categories {
//		if v.Mark=="" {
//			continue
//		}
//
//		var i = 1	//页数
//		for  {
//			if i==1 {
//				this.Url = v.Mark
//			}else {
//				this.Url =v.Mark + "/index_"+strconv.Itoa(i)+ ".html"
//			}
//			resp, err := http.Get(this.Url)
//			if err!=nil {
//				logs.Error.Info(err)
//			}
//			defer resp.Body.Close()
//			if resp.StatusCode != 200 {
//				break
//			}
//			document, err := goquery.NewDocumentFromReader(resp.Body)
//			if err!=nil {
//				logs.Error.Info(err)
//			}
//
//			i++
//		}
//
//	}
//
//
//}

