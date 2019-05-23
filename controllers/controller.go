package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type controllerInfo struct {
	regex          *regexp.Regexp
	params         map[int]string
	controllerType reflect.Type
}

type ControllerRegister struct {
	routers     []*controllerInfo
	Application *beego.App
}

func (p *ControllerRegister) Add(pattern string, c beego.ControllerInterface) {
	parts := strings.Split(pattern,"/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts{
		if strings.HasPrefix(part,":"){
			expr := "([^/]+)"

			if index := strings.Index(part,"("); index != -1{
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j ++
		}
	}

	pattern = strings.Join(parts,"/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {
		panic(regexErr)
		return
	}

	t := reflect.Indirect(reflect.ValueOf(c)).Type()

	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	p.routers = append(p.routers, route)
}

func (p *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	defer func() {
		if err := recover(); err != nil{

		}
	}()
}