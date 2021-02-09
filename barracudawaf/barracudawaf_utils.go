package barracudawaf

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	//Auth Token Endpoint for cudaWAF resource's CRUD operations
	cudaWAFResourceTokenEndpoint = "restapi/v3/login"
)

//getCudaWAFResourceURI : returns cudaWAF resource REST URI
func getCudaWAFResourceURI(resourceEndpoint string) string {
	cudaWAFResourceURL := fmt.Sprintf("://%s:%s/%s", WAFConfig.IPAddress, WAFConfig.AdminPort, resourceEndpoint)
	if WAFConfig.AdminPort == "8443" {
		cudaWAFResourceURL = fmt.Sprintf("%s%s", "https", cudaWAFResourceURL)
	} else {
		cudaWAFResourceURL = fmt.Sprintf("%s%s", "http", cudaWAFResourceURL)
	}
	return cudaWAFResourceURL
}

//getCudaWAFResourceToken : returns cudaWAF REST auth token
func getCudaWAFResourceToken() (int, map[string]interface{}) {

	//get cudaWAF resource auth token
	loginURL := getCudaWAFResourceURI(cudaWAFResourceTokenEndpoint)

	//auth token credential parameters
	requestBody, _ := json.Marshal(map[string]string{
		"username": WAFConfig.Username,
		"password": WAFConfig.Password,
	})

	//cudaWAF resource update client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, loginURL, bytes.NewBuffer(requestBody))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	//Make RestAPI call to get Login Token
	loginResponse, loginError := client.Do(req)
	if loginError != nil {
		return -1, map[string]interface{}{"msg": loginError}
	}

	if loginResponse != nil {
		defer loginResponse.Body.Close()
	}

	// Read Login Token and return
	// TODO : Error Handling
	loginResponseBody, _ := ioutil.ReadAll(loginResponse.Body)
	var loginResponseBodyJSON map[string]interface{}

	unmarshalError := json.Unmarshal(loginResponseBody, &loginResponseBodyJSON)

	if unmarshalError != nil {
		return -1, map[string]interface{}{"msg": unmarshalError}
	}

	return loginResponse.StatusCode, loginResponseBodyJSON
}

//updateCudaWAFResourceObject : REST calls to update cudaWAF resource objects
func updateCudaWAFResourceObject(cudaWAFResourceData map[string]interface{}) (int, map[string]interface{}) {

	//Get auth Token for cudaWAF API calls
	loginStatus, loginResponseBody := getCudaWAFResourceToken()

	if loginStatus == 200 || loginStatus == 201 {
		authToken := loginResponseBody["token"].(string) + ":"
		authToken = "BASIC " + b64.StdEncoding.EncodeToString([]byte(authToken))

		callURL := getCudaWAFResourceURI(cudaWAFResourceData["endpoint"].(string))
		operation := cudaWAFResourceData["operation"].(string)
		requestBody, marshalError := json.Marshal(cudaWAFResourceData["payload"])

		if marshalError != nil {
			return -1, map[string]interface{}{"msg": marshalError}
		}

		//RestAPI Client for POST
		req := &http.Request{}
		client := &http.Client{}

		//CRUD
		if operation == "DELETE" {
			req, _ = http.NewRequest(http.MethodDelete, callURL, nil)
		} else if operation == "POST" {
			req, _ = http.NewRequest(http.MethodPost, callURL, bytes.NewBuffer(requestBody))
		} else if operation == "PUT" {
			req, _ = http.NewRequest(http.MethodPut, callURL, bytes.NewBuffer(requestBody))
		} else if operation == "GET" {
			req, _ = http.NewRequest(http.MethodGet, callURL, bytes.NewBuffer(requestBody))
		} else {
			return -1, map[string]interface{}{"msg": "Operation not supported"}
		}

		// Set Headers i.e. Auth, Content-Type
		req.Header.Set("Authorization", authToken)
		req.Header.Set("Content-Type", "application/json")

		//REST call to update cudaWAF resource
		updateResponse, updateError := client.Do(req)

		if updateError != nil {
			return -1, map[string]interface{}{"msg": updateError}
		}

		if updateResponse != nil {
			defer updateResponse.Body.Close()
		}

		updateResponseBody, updateResponseBodyError := ioutil.ReadAll(updateResponse.Body)
		if updateResponseBodyError != nil {
			return -1, map[string]interface{}{"msg": updateResponseBodyError}
		}

		var updateResponseBodyJSON map[string]interface{}
		unmarshalError := json.Unmarshal(updateResponseBody, &updateResponseBodyJSON)

		if unmarshalError != nil {
			return -1, map[string]interface{}{"msg": unmarshalError}
		}

		return updateResponse.StatusCode, updateResponseBodyJSON
	}

	return -1, loginResponseBody
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func uploadCertificateContent(client *http.Client, url string, values map[string]io.Reader) (int, map[string]interface{}) {

	// Prepare a form : will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an certificate file
		if x, ok := r.(*os.File); ok {
			name := strings.Split(x.Name(), "/")
			filenameName := name[len(name)-1]
			if _, err := w.CreateFormFile(key, filenameName); err != nil {
				return -1, nil
			}
		} else {
			// Add other fields
			if _, err := w.CreateFormField(key); err != nil {
				return -1, nil
			}
		}
		if _, err := io.Copy(fw, r); err != nil {
			return -1, nil
		}
	}

	// If not closed : request will be missing the terminating boundary.
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return -1, nil
	}
	// Set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	loginStatus, loginResponseBody := getCudaWAFResourceToken()
	if loginStatus == 200 || loginStatus == 201 {
		token := loginResponseBody["token"].(string) + ":"
		token = "BASIC " + b64.StdEncoding.EncodeToString([]byte(token))
		req.Header.Set("Authorization", token)
		req.Header.Set("accept", "application/json")
	}

	// Submit the request
	resp, err := client.Do(req)

	if err != nil {
		return -1, nil
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	Body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return -1, nil
	}

	var BodyJSON map[string]interface{}
	err3 := json.Unmarshal(Body, &BodyJSON)

	if err3 != nil {
		return -1, nil
	}

	return resp.StatusCode, BodyJSON
}
