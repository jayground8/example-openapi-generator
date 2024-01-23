package main

import (
	"context"
	"fmt"
	"os"

	openapi "github.com/jayground8/example-openapi-generator/client"
)

func getDomainId(client *openapi.APIClient, domainName string) *int64 {
	req := client.DefaultAPI.GetDomain(context.Background()).
		Page(0).
		Size(10).
		DomainName(domainName)
	value, res, err := client.DefaultAPI.GetDomainExecute(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	if res.StatusCode != 200 || len(value.GetContent()) <= 0 {
		return nil
	}

	content := value.GetContent()[0]
	println("Id:", *content.Id)
	println("Name:", *content.Name)
	println("Status:", *content.Status)
	println("CompleteYn:", *content.CompleteYn)

	return content.Id
}

func applyRecordChange(client *openapi.APIClient, domainId *int64) {
	req := client.DefaultAPI.ApplyRecordChange(context.Background(), *domainId)
	res, err := client.DefaultAPI.ApplyRecordChangeExecute(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	println(res.StatusCode)
}

func getRecord(client *openapi.APIClient, domainId *int64, recordType string, recordName string) *openapi.GetRecord200ResponseContentInner {
	req := client.DefaultAPI.GetRecord(context.Background(), *domainId).
		Page(0).
		Size(10).
		RecordType(recordType)

	value, res, err := client.DefaultAPI.GetRecordExecute(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	for _, c := range value.GetContent() {
		println("Id:", *c.Id)
		println("Name:", *c.Name)
		println("Host:", *c.Host)
		println("Type:", *c.Type)
		println("Content:", *c.Content)
		if *c.Name == recordName {
			return &c
		}
	}

	return nil
}

func createRecord(client *openapi.APIClient, domainId *int64, host string, recordType string, content string, ttl int64) {
	body := []openapi.CreateRecordRequestInner{{Host: host, Type: recordType, Content: content, Ttl: ttl}}
	req := client.DefaultAPI.CreateRecord(context.Background(), *domainId).
		CreateRecordRequestInner(body)
	res, err := client.DefaultAPI.CreateRecordExecute(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	println(res.StatusCode)
}

func updateRecord(client *openapi.APIClient, domainId *int64, id int64, host string, recordType string, content string, ttl int64) {
	body := []openapi.UpdateRecordRequestInner{{Id: id, Host: host, Type: recordType, Content: content, Ttl: ttl}}
	req := client.DefaultAPI.UpdateRecord(context.Background(), *domainId).
		UpdateRecordRequestInner(body)
	res, err := client.DefaultAPI.UpdateRecordExecute(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	println(res.StatusCode)
}

func deleteRecord(client *openapi.APIClient, domainId *int64, recordId int64) {
	req := client.DefaultAPI.DeleteRecord(context.Background(), *domainId).RequestBody([]int64{recordId})
	res, err := client.DefaultAPI.DeleteRecordExecute(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}

	println(res.StatusCode)
}

func main() {
	config := openapi.NewConfiguration()
	client := openapi.NewAPIClient(config)
	domainId := getDomainId(client, "yourdomain.com")
	println(*domainId)

	createRecord(client, domainId, "test", "A", "8.8.8.8", 600)
	applyRecordChange(client, domainId)
	record := getRecord(client, domainId, "A", "sub.yourdomain.com")
	println(*record.Name, *record.Id)
	updateRecord(client, domainId, *record.Id, *record.Host, *record.Type, "7.7.7.7", *record.Ttl)
	applyRecordChange(client, domainId)
	deleteRecord(client, domainId, *record.Id)
	applyRecordChange(client, domainId)
}
