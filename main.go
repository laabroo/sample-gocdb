package main

import (
	"log"

	"github.com/laabroo/gocdb"
)

//Don't forget to change this value by your own setup
const (
	CdbSchema         = "http"
	CdbHost           = "127.0.0.1"
	CdbPort           = 5984
	CdbUser           = "YOUR_USER"
	CdbPassword       = "YOUR_PASSWORD"
	CdbDocName        = "YOUR_DOC_NAME"
	CdbIndexViewName  = "YOUR_INDEX_VIEW_NAME"
	CdbSingleViewName = "YOUR_SINGLE_VIEW_NAME"
	CdbKey            = "YOUR_KEY"
	CdbDbNameArticle  = "YOUR_DATABASE_NAME"
	CdbDocId          = "YOUR_DOC_ID"
)

//Declaration couchdb
var couchDB = gocdb.CouchDB{
	Schema:       CdbSchema,
	Host:         CdbHost,
	Port:         CdbPort,
	User:         CdbUser,
	Password:     CdbPassword,
	DatabaseName: CdbDbNameArticle,
}

func getDataById(docId string) {
	data, err := gocdb.FindById(docId)
	if err != nil {
		log.Panic(err)
	}

	log.Println(data)
}

func getDataByView(docName string, viewName string) {
	data, err := gocdb.FindByView(docName, viewName, 2, 0)
	if err != nil {
		log.Panic(err)
	}

	log.Println(data)
}

func getDataByViewWithKey(docName string, viewName string, key string) {
	data, err := gocdb.FindByViewWithKey(docName, viewName, 1, 0, key)
	if err != nil {
		log.Panic(err)
	}

	log.Println(data)
}

func main() {
	//initialize couchdb
	gocdb.Init(&couchDB)
	//find data by document id
	log.Println("===== GET DATA BY ID =====")
	getDataById(CdbDocId)
	log.Println("")

	//find data by view name
	log.Println("===== GET DATA BY VIEW =====")
	getDataByView(CdbDocName, CdbIndexViewName)
	log.Println("")

	//find data by key
	log.Println("===== GET DATA BY KEY =====")
	getDataByViewWithKey(CdbDocName, CdbSingleViewName, CdbKey)
	log.Println("")

}
