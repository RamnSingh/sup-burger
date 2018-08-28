package utils

import(
  "net/http"
  "github.com/gorilla/sessions"
  // "fmt"
)

var store = sessions.NewCookieStore([]byte("Shhhhhhhhhhhhhhhh"))
const sessionName string = "sup-burger"

func SaveToSession(key interface{}, value interface{}, res http.ResponseWriter, req *http.Request) error{
  if session, err := store.Get(req, sessionName); err != nil {
    return err
  }else{
    session.Values[key] = value
    session.Save(req, res)
    return nil
  }
}
func GetFromSession(key interface{}, req *http.Request) (interface{}, error){
  if session, err := store.Get(req, sessionName); err != nil {
    return nil, err
  }else{
    return session.Values[key], nil
  }
}
func RemoveFromSession(key interface{}, res  http.ResponseWriter, req *http.Request) error {
  if session, err := store.Get(req, sessionName); err != nil {
    return  err
  }else{
    session.Values[key] = nil
    session.Save(req, res)
    return nil
  }
}
func DestroySession(res http.ResponseWriter, req *http.Request) error{
  if session, err := store.Get(req, sessionName); err != nil {
    return err
  }else{
    session.Options.MaxAge = -1
    session.Save(req, res)
    return nil
  }
}
