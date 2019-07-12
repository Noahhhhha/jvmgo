package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	//声明一个结构体指针
	cmd := &Cmd{}
	//我认为是用户指南的存在
	flag.Usage = printUsage
	//用指定的名称、默认值、使用信息注册一个bool类型flag，并将flag的值保存到&cmd.helpFlag指向的地址里
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//解析命令到相应的已注册的变量上
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	//返回解析结果的指针
	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
