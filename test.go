package main

import (
	. "SubnetCalSERVER/Models"
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"os/exec"
)

const (
	ShellToUse = "/bin/sh"
	SubnetPath = "/home/gkay/go/src/subnetcal/subnetcal"
)

var (
	IpAddr      string
	SubnetValue string
	Jsn         = Jsonfile{}
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", mux)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	IpAddr = r.Form.Get("IpAddr")
	SubnetValue = r.Form.Get("SubnetValue")

	_, out, _ := Shellout(SubnetPath + " " + IpAddr + " " + SubnetValue)
	json.Unmarshal([]byte(out), &Jsn)

	t, _ := template.ParseFiles("Pages/Test.html")

	if len(r.Form.Get("IpAddr")) == 0 {
		t.Execute(w, nil)

	} else if len(Jsn.NetAddr) == 0 {
		t.Execute(w, nil)
	} else {
		t.Execute(w, Jsn)
	}

}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()

}
