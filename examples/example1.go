package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-errctx"
)

func foo() (int, error) {
	age, err := getAge("foo")
	if err != nil {
		return 0, errctx.Wrap(err, nil, "failed to foo")
	}
	return age, err
}

func getAge(name string) (int, error) {
	return 0, errctx.Wrap(fmt.Errorf("invalid name"), errctx.Fields{
		"name": name,
	}, "failed to get an age")
}

func main() {
	age, err := foo()
	if err != nil {
		if e, ok := err.(errctx.Error); ok {
			logrus.WithFields(logrus.Fields(e.Fields())).Fatal(e.Error())
		}
		logrus.Fatal(err)
	}
	fmt.Printf("age: %d\n", age)
}
