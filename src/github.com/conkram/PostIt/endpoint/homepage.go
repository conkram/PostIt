package endpoints

import (
  "net/http"
  "log"
  "html/template"
)

type PageVariables struct {
  PageTitle string
  Answer string
}

func Homepage(w http.ResponseWriter, r *http.Request) {
  //Display textbox and button

  Title := "Put your input here."

  FinalPageVariables := PageVariables {
    PageTitle: Title,
  }

  t, err := template.ParseFiles("../../static/template/homepage.html") //parse the html file homepage.html
  if err != nil { // if there is an error
    log.Print("template parsing error: ", err) // log it
  }

  err = t.Execute(w, FinalPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
  if err != nil { // if there is an error
    log.Print("template executing error: ", err) //log it
  }
}