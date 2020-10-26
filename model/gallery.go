package model

import (
	"context"
	"encoding/json"
	"lhc.go.crawler/libs/es"
	"lhc.go.crawler/logs"
)

type NetbianImg struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Alt  string `json:"alt"`
	Details string `json:"details"`
	Src  string `json:"src"`
	Sort int `json:"sort"`
	Mark string `json:"mark"`
	Origin string `json:"origin"`
	TickTime string `json:"tick_time"`
	Lists []string `json:"lists"`
}

func NewNetbianImg() *NetbianImg {
	return &NetbianImg{}
}

func (this *NetbianImg) GetList(param Params)(data []*NetbianImg,total int64,err error)  {

	result, err := es.Client.Search("gallery").From(param.Satrt).Size(param.Length).Do(context.Background())
	if err!=nil {
		return nil,0,err
	}
	total = result.Hits.TotalHits.Value
	if total > 0 {
		for _,v := range result.Hits.Hits {
			var tmp NetbianImg
			err := json.Unmarshal(v.Source, &tmp)
			if err != nil {
				logs.Error.Error(err)
				return nil,0,err
			}
			tmp.Id = v.Id
			data = append(data, &tmp)
		}
	}
	return
}