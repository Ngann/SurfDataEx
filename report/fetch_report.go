package report

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// An interface is a collection of methods that an object can implement(behaviours of the object)
// We defined a Fetch interface which has a moth GetData which outputs a string and error

type Fetch interface {
	GetData() (string, error)
}

// We defined a UrlFetch struct with feild url which is a string data type.
type UrlFetcher struct {
	url string
}

// the NewUrlFetcher function takes the a url which is a string datatype and returns a new UrlFetcher struct
// Pointers : In this case * will modify the original return value
// the * will give us access to the value the pointer points to. When we write *UrlFetcher, we are saying "store the url string in location/memory from &UrlFetcher{url: url}
func NewUrlFetcher(url string) *UrlFetcher {
	return &UrlFetcher{url: url}
}

// fetcher represents a UrlFetcher struct calls the GetData method and returns the a string and error
func (fetcher UrlFetcher) GetData() (string, error) {
	//the resp is the result from the the httt.Get request which passes the url from the struct
	resp, err := http.Get(fetcher.url)
	//if there is an error , return the error message
	if err != nil {
		return "", err
	}
	//if the reponse statuscode is not 200, returns "NOAA did not return 200 response"
	// error.New returns an error that formats as the given text.
	if resp.StatusCode != 200 {
		return "", errors.New("NOAA did not return 200 response")
	}
	// defer statment defers the execution of closing the response until all previous functions are returned.
	// resp.Body.Close() tells the client to close the response body when the we are done with the response object
	// if the connection remains opens, then there is a resource leak.
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// if there are no error return the string of b which is the body of the response.
	return string(b), nil
}

// we defined the MockFetcher struct with feild data  which is a string datatype.
type MockFetcher struct {
	data string
}

// fetcher represents the MockFetcher struct cna calls the GetData method and returns a string and error
// string is the data feild
func (fetcher MockFetcher) GetData() (string, error) {
	return fetcher.data, nil
}
