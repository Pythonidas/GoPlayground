package main

import (
  "fmt"
  "github.com/sendgrid/sendgrid-go"
  )

func main () {
  sg := sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")
  message := sendgrid.NewMail()
  message.AddTo("cn@hk.rs")
  message.AddToName("Yamil Asusta")
  message.SetSubject("SendGrid Testing")
  message.SetText("WIN")
  message.SetFrom("cn@hk.rs")
  if r:= sg.Send(message); r == nil {
      fmt.Println("Email sent!")
  } else {
    fmt.Println(r)
  }
}
