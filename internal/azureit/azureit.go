package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func startContainer() (*http.Response, error) {
	azureResourceId := os.Getenv("ACI_CONTAINER_GROUP")
	if azureResourceId == "" {
		return nil, fmt.Errorf("ACI_CONTAINER_GROUP is not specified")
	}
	azureRM := azure.PublicCloud.ResourceManagerEndpoint
	// let's rely on MSI to do all of the passwordless auth
	authorizer, err := auth.MSIConfig{
		// when you use a user-assigned identity, you need to specify the client_id of the MSI
		ClientID: os.Getenv("AZURE_CLIENT_ID"),
		Resource: azureRM,
	}.Authorizer()
	if err != nil {
		return nil, err
	}
	// https://docs.microsoft.com/en-us/rest/api/container-instances/container-groups/start
	managementResourceURL := fmt.Sprintf("%s%s/start?api-version=2021-09-01", azureRM, azureResourceId)
	req, err := http.NewRequest("POST", managementResourceURL, nil)
	if err != nil {
		return nil, err
	}
	// add Bearer token to crafted request
	_, err = autorest.Prepare(req, authorizer.WithAuthorization())
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func triggerStart(w http.ResponseWriter, _ *http.Request) {
	res, err := startContainer()
	if err != nil {
		log.Printf("[Error] %s", err)
		w.WriteHeader(400)
		return
	}
	defer res.Body.Close()
	// this function just proxies the start of ACI group
	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

// entry point for azure function to trigger ACI container with integration tests on time schedule
func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/api/TriggerStart", triggerStart)
	http.HandleFunc("/api/Debug", triggerStart)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
	if err != nil {
		panic(err)
	}
}
