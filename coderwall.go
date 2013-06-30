package coderwall

import (
  "net/http"
  "io"
  "io/ioutil"
  "encoding/json"
)

const BaseUrl = "http://coderwall.com"

type client struct {
  httpClient HttpClient
}

type HttpClient interface {
  Get(string) (resp *http.Response, err error)
}

type Profile struct {
  Username string
  Name string
  Location string
  Endorsements int
  Team string
  Accounts map[string]string
  Badges []struct {
    Name string
    Description string
    Created string
    Badge string
  }
}

func (c client) PerformRequest(url string) (resp *http.Response, err error) {
  return c.httpClient.Get(url)
}

func (c client) ProfileUrl(username string) string {
  return BaseUrl + "/" + username + ".json"
}

func (c client) LoadProfileFromJSON(profile *Profile, jsonContent []byte) (error) {
  err := json.Unmarshal(jsonContent, &profile)

  return err
}

func (c client) ParseBody(profile *Profile, body io.ReadCloser) (error) {
  defer body.Close()
  content, err := ioutil.ReadAll(body)
  if err != nil {
    return err
  }
  c.LoadProfileFromJSON(profile, content)

  return nil
}

func (c client) GetProfile(username string) (Profile, error) {
  var profile Profile
  url := c.ProfileUrl(username)
  res, err := c.PerformRequest(url)
  if err != nil {
    return profile, err
  }
  err = c.ParseBody(&profile, res.Body)

  return profile, err
}

func NewClient() client {
  var c client
  c.httpClient = &http.Client{}

  return c
}
