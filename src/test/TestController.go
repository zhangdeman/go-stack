package test

import (
	"net/http"
	"fmt"
)

type TestController struct {

}

func (test *TestController) TestMethod(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello World")
}

func TestMethod(w http.ResponseWriter, r *http.Request)  {
	test := TestController{}
	test.TestMethod(w,r)
}