package coderwall

import (
  "testing"
  "net/http"
)

type FakeHttpClient struct {
  getCount int
  lastUrl string
}

func (f *FakeHttpClient) Get(url string) (resp *http.Response, err error){
  f.getCount++
  f.lastUrl = url
  return nil, nil
}

func AssertEqual(t *testing.T, expected, actual interface{}) {
  if expected != actual {
    t.Errorf("%s is not equal to %s", actual, expected)
  }
}

func TestPerformRequest(t *testing.T) {
  var c client
  httpClient := new(FakeHttpClient)
  c.httpClient = httpClient
  url := "http://example.com"
  c.PerformRequest(url)
  AssertEqual(t, 1, httpClient.getCount)
  AssertEqual(t, url, httpClient.lastUrl)
}

func TestProfileUrl(t *testing.T) {
  expectedUrl := "http://coderwall.com/foo.json"
  var c client
  url := c.ProfileUrl("foo")
  AssertEqual(t, expectedUrl, url)
}

var ResponseBodyFixture = []byte(`{
  "username":"foo",
  "name": "Foo Bar",
  "location": "World",
  "endorsements": 100,
  "team": "12345",
  "accounts": {
    "github": "foo-bar-github",
    "linkedin": "foo-bar-linkedin"
  },
  "badges": [
    {
      "name": "Badge 1",
      "description": "description 1",
      "created": "2013-06-09T20:22:10Z",
      "badge": "badge1.png"
    }
  ]
}`)

func TestLoadProfile(t *testing.T) {
  var c client
  var profile Profile
  c.LoadProfileFromJSON(&profile, ResponseBodyFixture)
  AssertEqual(t, "foo", profile.Username)
  AssertEqual(t, "Foo Bar", profile.Name)
  AssertEqual(t, "World", profile.Location)
  AssertEqual(t, 100, profile.Endorsements)
  AssertEqual(t, "12345", profile.Team)
}

func TestLoadAcounts(t *testing.T) {
  var c client
  var profile Profile
  c.LoadProfileFromJSON(&profile, ResponseBodyFixture)
  AssertEqual(t, "foo-bar-github", profile.Accounts["github"])
  AssertEqual(t, "foo-bar-linkedin", profile.Accounts["linkedin"])
}

func TestLoadBadges(t *testing.T) {
  var c client
  var profile Profile
  c.LoadProfileFromJSON(&profile, ResponseBodyFixture)
  AssertEqual(t, 1, len(profile.Badges))
  AssertEqual(t, "Badge 1", profile.Badges[0].Name)
  AssertEqual(t, "description 1", profile.Badges[0].Description)
  AssertEqual(t, "2013-06-09T20:22:10Z", profile.Badges[0].Created)
  AssertEqual(t, "badge1.png", profile.Badges[0].Badge)
}
