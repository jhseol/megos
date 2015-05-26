package megos

import (
  "io/ioutil"
  "net/http"
  "net/url"
  "strings"
  "time"
)

type Query map[string]string

var httpClient = http.Client{
  Timeout: 5 * time.Second,
}

func httpHeadRequest(URL string, query Query) (int, error) {
  if strings.HasPrefix(URL, "http://") == false {
    URL = "http://" + URL
  }
  u, err := url.Parse(URL)
  if err != nil {
    return 500, err
  }
  q := u.Query()
  for k, v := range query {
    q.Set(k, v)
  }
  u.RawQuery = q.Encode()
  URL = u.String()

  resp, err := httpClient.Head(URL)
  if err != nil {
    return 500, err
  }
  defer resp.Body.Close()

  return resp.StatusCode, err
}

func httpGetRequest(URL string, query Query) ([]byte, error) {
  if strings.HasPrefix(URL, "http://") == false {
    URL = "http://" + URL
  }
  u, err := url.Parse(URL)
  if err != nil {
    return nil, err
  }
  q := u.Query()
  for k, v := range query {
    q.Set(k, v)
  }
  u.RawQuery = q.Encode()
  URL = u.String()

  resp, err := httpClient.Get(URL)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return body, err
}
