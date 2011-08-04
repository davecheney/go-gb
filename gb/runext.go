/* 
   Copyright 2011 John Asmuth

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package main

import (
	"os"
	"exec"
	"fmt"
	"strings"
)

var MakeCMD,
	CompileCMD,
	CCMD,
	AsmCMD,
	LinkCMD,
	PackCMD,
	CopyCMD,
	GoInstallCMD,
	GoFMTCMD,
	CGoCMD,
	GCCCMD string

func FindExternals() (err os.Error) {

	CompileCMD, err = exec.LookPath(GetCompilerName())
	if err != nil {
		fmt.Printf("Could not find '%s' in path\n", GetCompilerName())
		return
	}
	AsmCMD, err = exec.LookPath(GetAssemblerName())
	if err != nil {
		fmt.Printf("Could not find '%s' in path\n", GetAssemblerName())
		return
	}
	LinkCMD, err = exec.LookPath(GetLinkerName())
	if err != nil {
		fmt.Printf("Could not find '%s' in path\n", GetLinkerName())
		return
	}
	PackCMD, err = exec.LookPath("gopack")
	if err != nil {
		fmt.Printf("Could not find 'gopack' in path\n")
		return
	}
	CGoCMD, err = exec.LookPath("cgo")
	if err != nil {
		fmt.Printf("Could not find 'cgo' in path\n")
		return
	}

	var err2 os.Error
	MakeCMD, err2 = exec.LookPath("gomake")
	if err2 != nil {
		fmt.Printf("Could not find 'gomake' in path\n")
	}
	GoInstallCMD, err2 = exec.LookPath("goinstall")
	if err2 != nil {
		fmt.Printf("Could not find 'goinstall' in path\n")
	}
	GoFMTCMD, err2 = exec.LookPath("gofmt")
	if err2 != nil {
		fmt.Printf("Could not find 'gofmt' in path\n")
	}
	GCCCMD, err2 = exec.LookPath("gcc")
	if err2 != nil {
		fmt.Printf("Could not find 'gcc' in path\n")
	}
	CCMD, err2 = exec.LookPath(GetCCompilerName())
	if err2 != nil {
		fmt.Printf("Could not find '%' in path\n", GetCCompilerName())
	}

	CopyCMD, _ = exec.LookPath("cp")

	return
}

func SplitArgs(args []string) (sargs []string) {
	for _, arg := range args {
		sarg := strings.Split(arg, " ")
		sargs = append(sargs, sarg...)
	}
	return
}

func RunExternalDump(cmd, wd string, argv []string, dump *os.File) (err os.Error) {
	argv = SplitArgs(argv)

	c := exec.Command(cmd, argv[1:]...)
	c.Dir = wd
	c.Env = os.Environ()

	c.Stdout = dump
	c.Stderr = os.Stderr

	err = c.Run()

	if wmsg, ok := err.(*os.Waitmsg); ok {
		if wmsg.ExitStatus() != 0 {
			err = os.NewError(fmt.Sprintf("%v: %s\n", argv, wmsg.String()))
		} else {
			err = nil
		}
	}
	return
}
func RunExternal(cmd, wd string, argv []string) (err os.Error) {
	return RunExternalDump(cmd, wd, argv, os.Stdout)
}
