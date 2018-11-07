package monitor

type Dimension struct {
   Width  string
   Height string
}

type MonitoringEvent interface {}

type monitoringEvent struct {
   EventType  string
   WebsiteUrl string
   SessionId string
   ResizeFrom Dimension
   ResizeTo Dimension
   Time int
   Pasted bool
   FormId string
}

// Return a standard monitoringEvent
func NewEvent() monitoringEvent{
   return monitoringEvent{"",                  //EventType
                          "",                  //WebsiteUrl
                          "",                  //SessionId
                          Dimension{"0","0"},  //ResizeFrom
                          Dimension{"0","0"},  //ResizeTo
                          0,                   //Time
                          false,               //Pasted
                          ""}                  //FormId
}