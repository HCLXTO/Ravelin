package monitor

import (
  "fmt"
  "strconv"
)

type Dimension struct {
  Width  string
  Height string
}

type Data struct {
  WebsiteUrl         string
  SessionId          string
  ResizeFrom         Dimension
  ResizeTo           Dimension
  CopyAndPaste       map[string]bool // map[fieldId]true
  FormCompletionTime int // Seconds
}

// General monitoring event (has all the monitored fields)
type event struct {
  EventType  string
  WebsiteUrl string
  SessionId string
  ResizeFrom Dimension
  ResizeTo Dimension
  Time int
  Pasted bool
  FormId string
}

// Basic response for the event process, status and a copy of the Data
type resp struct {
   Status  bool
   Data Data
}

// Map to store the user's session data, map[SessionId] Data
var store map[string] Data

/* CONSTRUCTORS */

// Return a standard event, that can acept any monitoring event change
func NewEvent() event{
return event{"",                  //EventType
             "",                  //WebsiteUrl
             "",                  //SessionId
             Dimension{"0","0"},  //ResizeFrom
             Dimension{"0","0"},  //ResizeTo
             0,                   //Time
             false,               //Pasted
             ""}                  //FormId
}

// Return a new Data struct with event's basic info
func newData(e* event) Data{
  return Data{e.WebsiteUrl,           //WebsiteUrl
              e.SessionId,            //SessionId
              Dimension{"0","0"},     //ResizeFrom
              Dimension{"0","0"},     //ResizeTo
              make(map[string]bool),  //CopyAndPaste
              0}                      //FormCompletionTime
}

// Return a resposnse struct
func newResponse(status bool, data Data) resp {
  return resp{status, data}
}

/* HASH FUNCTION */

// Implementation of the djb2 hash http://www.cse.yorku.ca/~oz/hash.html
func hash(s string) uint64{
  var hash uint64
  hash = 5381//18446744073709551614
  
  for _, c := range(s){
    hash = ((hash << 5) + hash) + uint64(c)
  }
  return hash
}

/* STORE FUNCTIONS */

// Returns the Data stored based on the event's SessionId
func getData(e* event) Data{
  if store == nil {
    store = make(map[string]Data)
  }
  d, ok := store[e.SessionId]
  // If don't have this session data stored, make a new one
  if ok == false {
    d = newData(e)
  }
  return d
}

// Save the Data Struct on store under its SessionId
func setData(d Data) {
  if store == nil {
    store = make(map[string]Data)
  }
  store[d.SessionId] = d
}

// remove the Data Struct of the store
func removeData(d Data) {
  if store == nil {
    store = make(map[string]Data)
  }
  delete(store, d.SessionId)
}

// Here I would send the Data to another part of the system 
// but for this challenge I will only delete it from store
func finalizeSession(d Data) {
  fmt.Println("Form submited - Ending session.")
  removeData(d)
}

/* DATA FUNCTIONS */

func (d* Data) Print() {
  fmt.Println("Session Data:")
  fmt.Println("WebsiteUrl: ", d.WebsiteUrl)
  fmt.Println("WebsiteUrl HASH: ", hash(d.WebsiteUrl))
  fmt.Println("SessionId: ", d.SessionId)
  fmt.Println("ResizeFrom: ", d.ResizeFrom)
  fmt.Println("ResizeTo: ", d.ResizeTo)
  fmt.Println("CopyAndPaste: ", d.CopyAndPaste)
  fmt.Println("FormCompletionTime: ", d.FormCompletionTime)
  fmt.Println("----------------------------")
}

/* EVENT PROCESSING FUNCTIONS */

// Process the event's request according to EventType
func (e* event) Process() (resp, error){
  err := e.basicValidation()
  if err != nil {
    return newResponse(false,*new(Data)), err 
  }
  switch {
    case e.EventType == "screenResize":
      return e.resize()
    case e.EventType == "timeTaken":
      return e.timeTaken()
    case e.EventType == "copyAndPaste":
      return e.copyAndPaste()
    default:
      msg := "Event with wrong EventType."
      res := newResponse(false,*new(Data))
      return res, fmt.Errorf(msg)
  }
}

// Performs the screenResize update
func (e* event) resize() (resp, error){
  err := e.resizeValidation()
  if err != nil {
    return newResponse(false,*new(Data)), err 
  }
  fmt.Println("Resize Action!.")
  data := getData(e)
  data.ResizeFrom = e.ResizeFrom
  data.ResizeTo = e.ResizeTo
  setData(data)
  return newResponse(true,data), nil
}

// Performs the timeTaken update and finalize the session 
// this action happens when the form is submited
func (e* event) timeTaken() (resp, error){
  err := e.timeTakenValidation()
  if err != nil {
    return newResponse(false,*new(Data)), err 
  }
  fmt.Println("TimeTaken Action!.")
  data := getData(e)
  data.FormCompletionTime = e.Time
  setData(data)
  finalizeSession(data)
  return newResponse(true,data), nil
}

// Performs the timeTaken update
func (e* event) copyAndPaste() (resp, error){
  err := e.copyAndPasteValidation()
  if err != nil {
    return newResponse(false,*new(Data)), err 
  }
  fmt.Println("CopyAndPaste Action!.")
  data := getData(e)
  data.CopyAndPaste[e.FormId] = e.Pasted
  setData(data)
  return newResponse(true,data), nil
}

/* VALIDATION FUNCTIONS */

// Does the basic validation for all the common fields
func (e* event) basicValidation() (error){
  if e.EventType == "" {
    return fmt.Errorf("Empty EventType")
  }
  if e.WebsiteUrl == "" {
    return fmt.Errorf("Empty WebsiteUrl")
  }
  if e.SessionId == "" {
    return fmt.Errorf("Empty SessionId")
  }
  return nil
}

// Does the validation for the resize fields - No negative Dimension
func (e* event) resizeValidation() (error){
  fromW, err := strconv.Atoi(e.ResizeFrom.Width)
  if err != nil {
    return err
  }
  fromH, err := strconv.Atoi(e.ResizeFrom.Height)
  if err != nil {
    return err
  }
  toW, err := strconv.Atoi(e.ResizeTo.Width)
  if err != nil {
    return err
  }
  toH, err := strconv.Atoi(e.ResizeTo.Height)
  if err != nil {
    return err
  }
  if fromW < 0 || fromH < 0 {
    return fmt.Errorf("Bad values for ResizeFrom")
  }
  if toW < 0 || toH < 0 {
    return fmt.Errorf("Bad values for ResizeTo")
  }
  return nil
}

// Does the validation for the timeTaken fields, no negative time
func (e* event) timeTakenValidation() (error){
  if e.Time < 0 {
    return fmt.Errorf("Invalid Time")
  }
  return nil
}

// Does the validation for the copyAndPaste fields, no empty FormId
func (e* event) copyAndPasteValidation() (error){
  if e.FormId == "" {
    return fmt.Errorf("Empty FormId")
  }
  return nil
}


