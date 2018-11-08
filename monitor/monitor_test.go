// Unit tests for the monitor package
package monitor

import "testing"

func TestBasicValidation(t *testing.T) {
   event := NewEvent()
   event.EventType = "screenResize"
   event.WebsiteUrl = "www.test.com"
   event.SessionId = "123-132-123"
   
   err := event.basicValidation()
   if err != nil {
      t.Errorf("Fail basicValidation test: 1")
   }

   event.EventType = ""
   err = event.basicValidation()
   if err == nil {
      t.Errorf("Fail basicValidation test: 2")
   }
   event.EventType = "screenResize"

   event.WebsiteUrl = ""
   err = event.basicValidation()
   if err == nil {
      t.Errorf("Fail basicValidation test: 3")
   }
   event.WebsiteUrl = "www.test.com"

   event.SessionId = ""
   err = event.basicValidation()
   if err == nil {
      t.Errorf("Fail basicValidation test: 4")
   }
   event.SessionId = "123-132-123"
}

func TestResizeValidation(t *testing.T) {
   event := NewEvent()
   event.ResizeFrom.Width = "10"
   event.ResizeFrom.Height = "20"
   event.ResizeTo.Width = "50"
   event.ResizeTo.Height = "70"
   
   err := event.resizeValidation()
   if err != nil {
      t.Errorf("Fail resizeValidation test: 1")
   }

   event.ResizeTo.Height = "b"
   err = event.resizeValidation()
   if err == nil {
      t.Errorf("Fail resizeValidation test: 2")
   }
   event.ResizeTo.Height = "70"

   event.ResizeFrom.Height = "-10"
   err = event.resizeValidation()
   if err == nil {
      t.Errorf("Fail resizeValidation test: 3")
   }
   event.ResizeFrom.Height = "70"

   event.ResizeFrom.Width = "0"
   err = event.resizeValidation()
   if err != nil {
      t.Errorf("Fail resizeValidation test: 4")
   }
   event.ResizeFrom.Width = "10"
}

func TestTimeTakenValidation(t *testing.T) {
   event := NewEvent()
   event.Time = 10
   
   err := event.timeTakenValidation()
   if err != nil {
      t.Errorf("Fail timeTakenValidation test: 1")
   }

   event.Time = 0
   err = event.timeTakenValidation()
   if err != nil {
      t.Errorf("Fail timeTakenValidation test: 2")
   }

   event.Time = -10
   err = event.timeTakenValidation()
   if err == nil {
      t.Errorf("Fail timeTakenValidation test: 3")
   }
}

func TestCopyAndPasteValidation(t *testing.T) {
   event := NewEvent()
   event.FormId = "TestFrom"
   event.Pasted = true
   
   err := event.copyAndPasteValidation()
   if err != nil {
      t.Errorf("Fail copyAndPasteValidation test: 1")
   }

   event.Pasted = false
   err = event.copyAndPasteValidation()
   if err != nil {
      t.Errorf("Fail copyAndPasteValidation test: 2")
   }

   event.FormId = ""
   err = event.copyAndPasteValidation()
   if err == nil {
      t.Errorf("Fail copyAndPasteValidation test: 3")
   }
}

func TestCopyProcess(t *testing.T) {
   event := NewEvent()
   event.WebsiteUrl = "www.test.com"
   event.SessionId = "123-132-123"

   // Testing copyAndPaste event process
   event.EventType = "copyAndPaste"
   event.FormId = "TestFrom"
   event.Pasted = true
   
   res, err := event.Process()
   if err != nil {
      t.Errorf("Fail CopyProcess test: 1")
   }
   if res.Data.CopyAndPaste[event.FormId] != event.Pasted{
      t.Errorf("Fail CopyProcess test: 2")
   }

   event.FormId = "TestFrom2"
   event.Pasted = true
   
   res, err = event.Process()
   if err != nil {
      t.Errorf("Fail CopyProcess test: 1")
   }
   if res.Data.CopyAndPaste[event.FormId] != event.Pasted{
      t.Errorf("Fail CopyProcess test: 2")
   }

   event.FormId = "TestFrom"
   event.Pasted = false
   
   res, err = event.Process()
   if err != nil {
      t.Errorf("Fail CopyProcess test: 1")
   }
   if res.Data.CopyAndPaste[event.FormId] != event.Pasted{
      t.Errorf("Fail CopyProcess test: 2")
   }
}

func TestResizeProcess(t *testing.T) {
   event := NewEvent()
   event.WebsiteUrl = "www.test.com"
   event.SessionId = "123-132-123"

   // Testing copyAndPaste event process
   event.EventType = "screenResize"
   event.ResizeFrom.Height = "10"
   event.ResizeFrom.Width = "20"
   event.ResizeTo.Height = "50"
   event.ResizeTo.Width = "90"
   
   res, err := event.Process()
   if err != nil {
      t.Errorf("Fail ResizeProcess test: 1")
   }
   if res.Data.ResizeFrom.Height != event.ResizeFrom.Height{
      t.Errorf("Fail ResizeProcess test: 2")
   }
   if res.Data.ResizeFrom.Width != event.ResizeFrom.Width{
      t.Errorf("Fail ResizeProcess test: 3")
   }
   if res.Data.ResizeTo.Height != event.ResizeTo.Height{
      t.Errorf("Fail ResizeProcess test: 4")
   }
   if res.Data.ResizeTo.Width != event.ResizeTo.Width{
      t.Errorf("Fail ResizeProcess test: 5")
   }
}

func TestTimeProcess(t *testing.T) {
   event := NewEvent()
   event.WebsiteUrl = "www.test.com"
   event.SessionId = "123-132-123"

   // Testing copyAndPaste event process
   event.EventType = "timeTaken"
   event.Time = 10
   res, err := event.Process()
   if err != nil {
      t.Errorf("Fail TimeProcess test: 1")
   }
   if res.Data.FormCompletionTime != event.Time{
      t.Errorf("Fail TimeProcess test: 2")
   }

   event.Time = 50
   res, err = event.Process()
   if err != nil {
      t.Errorf("Fail TimeProcess test: 3")
   }
   if res.Data.ResizeFrom.Width != event.ResizeFrom.Width{
      t.Errorf("Fail TimeProcess test: 4")
   }

   event.Time = -50
   res, err = event.Process()
   if err == nil {
      t.Errorf("Fail TimeProcess test: 5")
   }
}
