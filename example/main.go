package main

import (
	"fmt"

	"github.com/ieee0824/formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	f := formatter.SysLogFormatter{}
	e := logrus.Entry{}

	bin, _ := f.Format(&e)
	fmt.Println(string(bin))
}
