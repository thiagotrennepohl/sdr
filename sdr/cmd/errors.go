package cmd

import "errors"

var FILE_NOT_FOUND_ERR = errors.New("Sorry the spcified file was not found")

var ERROR_OPENING_FILE = errors.New("Sorry something went wrong while reading the file")
