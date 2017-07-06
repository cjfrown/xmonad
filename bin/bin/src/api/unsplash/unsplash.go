package unsplash

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Unsplash struct {
  Urls map[string]string
  Links map[string]string
  Id string `json:"id"`
  Color string `json:"color"`
}

func New(k string)(unsplash * Unsplash) {
  data := GetJson(k)
  err := json.Unmarshal(data, &unsplash)
  if err != nil {
  }
  return
}

func GetJson(k string) (b []byte) {
//   url := "https://api.unsplash.com/" + "photos/random/" + "?client_id=" + k
    url := "https://api.unsplash.com/" + "photos/random?snow;" + "client_id=" + k
    r, err := http.Get(url)
    defer r.Body.Close()
    b, err = ioutil.ReadAll(r.Body)
    if err != nil {
    }
    return
}
