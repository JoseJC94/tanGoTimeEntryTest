//Service

package main

import (
    "net/http"
    "os"

    //"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    var err error
    readData("entries.csv")
    switch request.Method {
    case "GET":
        err = handleGet(writer, request)
    case "POST":
        err = handlePost(writer, request)
    case "PUT":
        err = handlePut(writer, request)
    case "DELETE":
        err = handleDelete(writer, request)
    }
    writeData("entries.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/entries/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    //arrayTest()
}