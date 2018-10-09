package main

import (
	"net/http"
	"os"
)

func main() {
	//p:="D:/goWork/src/firmMangement/FirmMangement/Fims/wifi"
	os.Mkdir("file",0777)
	http.Handle("/js",http.StripPrefix("/js", http.FileServer(http.Dir("file"))))
	http.ListenAndServe(":8888",nil)
}
