package barracudawaf

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const (
	loginEndpoint = "restapi/v3/login"
)

func buildCallingURL(endpoint string) string {
	var callingURL string
	if WAFConfig.AdminPort == "8443" {
		callingURL = "https://" + WAFConfig.IPAddress + ":" + WAFConfig.AdminPort + "/" + endpoint
	} else {
		callingURL = "http://" + WAFConfig.IPAddress + ":" + WAFConfig.AdminPort + "/" + endpoint
	}
	return callingURL
}

func getRestLoginToken() (int, map[string]interface{}) {
	//Build Login REST URL
	loginURL := buildCallingURL(loginEndpoint)
	log.Printf("Login URL : %v", loginURL)

	//Request Body for Login
	requestBody, _ := json.Marshal(map[string]string{
		"username": WAFConfig.Username,
		"password": WAFConfig.Password,
	})

	//RestAPI Client
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, loginURL, bytes.NewBuffer(requestBody))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	//Make RestAPI call to get Login Token
	resp, err1 := client.Do(req)
	if err1 != nil {
		log.Printf("Login Token Get Failed : %v", err1)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	// Read Login Token and return
	// TODO : Error Handling
	Body, _ := ioutil.ReadAll(resp.Body)
	var BodyJSON map[string]interface{}
	err := json.Unmarshal(Body, &BodyJSON)

	if err != nil {
		return -1, nil
	}

	return resp.StatusCode, BodyJSON
}

func doRestAPICall(callData map[string]interface{}) (int, map[string]interface{}) {
	// Logging for errors
	Logfile, err := os.OpenFile("/Users/ashishaswal/go_project/provider1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open Log file %s. err: %s", "File", err)
	} else {
		defer Logfile.Close()
		log.SetOutput(Logfile)
	}

	// Get Login Token For Auth
	loginStatus, loginRespBody := getRestLoginToken()
	log.Printf("Login Token : %v", loginRespBody)

	if loginStatus == 200 || loginStatus == 201 {
		token := loginRespBody["token"].(string) + ":"
		token = "BASIC " + b64.StdEncoding.EncodeToString([]byte(token))

		//Build CREATE / POST RestAPI URL
		callURL := buildCallingURL(callData["endpoint"].(string))
		operation := callData["operation"].(string)
		requestBody, _ := json.Marshal(callData["payload"])
		log.Printf("Call Data with Payload : %v", callData)

		//RestAPI Client for POST
		req := &http.Request{}
		client := &http.Client{}

		//Decide CRUD, what to do ?
		if operation == "DELETE" {
			req, _ = http.NewRequest(http.MethodDelete, callURL, nil)
		} else if operation == "POST" {
			req, _ = http.NewRequest(http.MethodPost, callURL, bytes.NewBuffer(requestBody))
		} else if operation == "PUT" {
			req, _ = http.NewRequest(http.MethodPut, callURL, bytes.NewBuffer(requestBody))
		} else if operation == "GET" {
			req, _ = http.NewRequest(http.MethodGet, callURL, bytes.NewBuffer(requestBody))
		} else {
			return -1, nil
		}

		// Set Headers i.e. Auth, Content-Type
		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")
		log.Printf("Request body : %v", req)

		//Make RestAPI Call to create resource
		resp, err1 := client.Do(req)
		log.Printf("Response from Server : %v", resp)

		if err1 != nil {
			log.Printf("REST API call Failed with error %v", err1)
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
		err := json.Unmarshal(Body, &BodyJSON)

		if err != nil {
			return -1, nil
		}

		return resp.StatusCode, BodyJSON
	}

	return -1, nil
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func uploadCertificateContent(client *http.Client, url string, values map[string]io.Reader) (int, map[string]interface{}) {

	Logfile, err := os.OpenFile("/Users/ashishaswal/go/barracuda_provider.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open Log file %s. err: %s", "File", err)
	} else {
		defer Logfile.Close()
		log.SetOutput(Logfile)
	}

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
			if fw, err = w.CreateFormFile(key, filenameName); err != nil {
				return -1, nil
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return -1, nil
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
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

	loginStatus, loginRespBody := getRestLoginToken()
	if loginStatus == 200 || loginStatus == 201 {
		token := loginRespBody["token"].(string) + ":"
		token = "BASIC " + b64.StdEncoding.EncodeToString([]byte(token))
		req.Header.Set("Authorization", token)
		req.Header.Set("accept", "application/json")
	}

	log.Printf("Request Body : %v", req)

	// Submit the request
	resp, err := client.Do(req)
	log.Printf("Response from Server : %v", resp)

	if err != nil {
		log.Printf("Response Error : %v", err)
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
