// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tActive struct {}
var Active tActive


func (_ tActive) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Active.Index", args).Url
}

func (_ tActive) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Active.Create", args).Url
}

func (_ tActive) PostCreate(
		active interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "active", active)
	return revel.MainRouter.Reverse("Active.PostCreate", args).Url
}

func (_ tActive) AddActive(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Active.AddActive", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) AddActive(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddActive", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


