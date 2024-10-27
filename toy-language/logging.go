package main

import (
    "log"
    // "io"
    "os"
)
var Info *log.Logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
var Error *log.Logger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags)
var Output *log.Logger = log.New(os.Stdout, "", 0)
var Assembly *log.Logger = log.New(os.Stdout, "", 0)
