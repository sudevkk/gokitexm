package main
 

import (
        "context"
        "errors"
        "strings"
)


type StringService interface{
        Uppercase(string) (string, error)
        Count(string) (int)
}


type stringService struct{}

func (srv stringService) Uppercase(s string) (string, err error) {

                if s == "" {
                        return "", errors.NewError("Empty String")
                }

                return string.ToUpper(s), nil
}

 

func (srv stringService) Uppercase(s string) (string) {
        return len(s)
}