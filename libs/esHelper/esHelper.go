package esHelper

import (
	"github.com/olivere/elastic"
	"gorestful/libs/logHelper"
	"gorestful/config"
	"context"
)

type Es struct {
	EsIndex string
	EsAddr string
	Client *elastic.Client
}

func CreateEs() *Es {
	this := new(Es)
	esConfig := Config.EsConfig
	this.EsIndex = esConfig["EsIndex"]
	this.EsAddr = "http://"+esConfig["EsAddr"]
	this.Client = connect(this.EsAddr)
	return this
}

func connect(esAddr string) *elastic.Client {

	client,err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(esAddr))

	if err != nil {
		logHelper.WriteLog("es 连接失败"+esAddr,"essearch/error")
		return nil
	}
	return client
}

func (this *Es) Add(esType string,id string,parentId string,data interface{}) bool {

	client := this.Client

	if client == nil {
		return false
	}
	IndexService := client.Index().Index(this.EsIndex).Type(esType).BodyJson(data).Refresh("true");

	if id != "" {
		IndexService.Id(id)
	}

	if parentId != "" {
		IndexService.Parent(parentId)
	}

	_,err := IndexService.Do(context.Background())

	if err != nil {
		logHelper.WriteLog("es 添加失败"+err.Error(),"essearch/add/error")
		return false
	}

	return true
}

func (this *Es) Edit (esType string ,id string,parentId string, data interface{}) bool {
	client := this.Client
	if client == nil {
		return false
	}

	UpadateService := client.Update().
		Index(this.EsIndex).
			Type(esType).
				DocAsUpsert(true).
					Refresh("true").
						RetryOnConflict(5)

	if id != "" {
		UpadateService.Id(id)
	}

	if parentId != "" {
		UpadateService.Parent(parentId)
	}

	_,err := UpadateService.Doc(data).Do(context.Background())

	if err != nil {
		logHelper.WriteLog("es 更新失败"+err.Error(),"essearch/edit/error")
		return false
	}

	return true
}
