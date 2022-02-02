package api

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func AuthorizeFactory(api string, username string, password string) ([]byte, error){
	postBody, _ := json.Marshal(map[string]string{
		"username":  username,
		"password": password,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(api + "/login", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func CreateKerberosAgent(api string, token string, name string, rtsp string) ([]byte, error){
	postBody, _ := json.Marshal(map[string]string{
		"name":  name,
		"rtsp": rtsp,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	req, err := http.NewRequest("POST", api + "/deployments", responseBody)
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func DeleteKerberosAgent(api string, token string, name string) ([]byte, error){
	req, err := http.NewRequest("DELETE", api + "/deployments/" + name , nil)
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func GetKerberosAgents(api string, token string) ([]byte, error){
	req, err := http.NewRequest("GET", api + "/deployments/services" , nil)
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
