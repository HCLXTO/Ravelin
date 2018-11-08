package main

import (
   "fmt"
   "log"
   "net/http"
   "encoding/json"
   "github.com/HCLXTO/Ravelin/monitor"
)
/*
// Basic response struct
type Response struct {
   Status  string
   Obs string
}
*/
// Handler for the monitoring events
func monitorHandler(w http.ResponseWriter, r *http.Request) {
   if r.Method == "POST"{
      if r.Body == nil{
         http.Error(w, "Request without a body.", 400)
         return
      }

      event := monitor.NewEvent()
      decoder := json.NewDecoder(r.Body)
      err := decoder.Decode(&event)
      if err != nil {
         http.Error(w, err.Error(), 400)
         return
      }
      
      response, err := event.Process() //Response {Status: "OK", Obs: event.EventType}
      if err != nil {
         fmt.Println(err.Error())
         http.Error(w, err.Error(), 422)
         return
      }

      response.Data.Print()

      json.NewEncoder(w).Encode(response.Status)

   } else{
      http.Error(w, "Wrong request method.", 400)
      return
   }
}
func main() {
   //Monitoring event handlers
   http.HandleFunc("/screenResize", monitorHandler)
   http.HandleFunc("/timeTaken", monitorHandler)
   http.HandleFunc("/copyAndPaste", monitorHandler)
   //Static content handlers   
   js := http.FileServer(http.Dir("static/js"))
   http.Handle("/js/", http.StripPrefix("/js/", js))

   http.Handle("/", http.FileServer(http.Dir("static/html")))
   
   fmt.Println("Server running on: http://localhost:8080/")
   fmt.Println("Press 'ctrl+c' to stop it.")
   log.Fatal(http.ListenAndServe(":8080", nil))

}