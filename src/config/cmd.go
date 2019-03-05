package config

import (
	"flag"
	"fmt"
)

type Cmd struct {
	StaticRootDir string
	Help bool
}

var Args *Cmd

func init(){
	Args = &Cmd{}
}

func ParseCmd() *Cmd{
	flag.StringVar(&Args.StaticRootDir ,"d" ,"html" ,"set static files root dir")
	flag.BoolVar(&Args.Help ,"h" ,false ,"get help")
	flag.Usage = Help
	flag.Parse()
	return Args
}

func Help(){
	fmt.Println("Usage : nyn_static [OPTIONS]")
	flag.PrintDefaults()
}
