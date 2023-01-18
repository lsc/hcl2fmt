// Copyright © 2021 Lowe Schmidt
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/mattn/go-zglob"
)

const Version = "0.1.1"

var (
	version    = flag.Bool("v", false, "version information")
	workingDir = flag.String("w", "", "Working directory, defaults to CWD")
)

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Println(Version)
		return nil
	}

	if *workingDir == "" {
		*workingDir, _ = os.Getwd()
	}

	if hclFiles, err := zglob.Glob(joinPath(*workingDir, "**", "*.hcl")); err != nil {
		fmt.Println(err)
		return err
	} else {
		for _, fname := range hclFiles {
			if stat, err := os.Stat(fname); err == nil {
				if err := checkErrors(fname); err != nil {
					fmt.Println(err)
					break
				}
				if contents, err := os.ReadFile(fname); err == nil {
					newContents := hclwrite.Format(contents)
					fmt.Println(fname)
					if err := os.WriteFile(fname, newContents, stat.Mode()); err != nil {
						fmt.Println("Unable to write file %s (%s)", fname, err)
					}
				} else {
					fmt.Println("Unable to read file %s. (%s)", fname, err)
				}
			}
		}
	}
	return nil
}

// checkErrors parsers the given hclFile and looks for syntax errors
func checkErrors(hclFile string) error {
	parser := hclparse.NewParser()
	_, diags := parser.ParseHCLFile(hclFile)
	if diags.HasErrors() {
		return diags
	}
	return nil
}

func joinPath(elem ...string) string {
	return filepath.ToSlash(filepath.Join(elem...))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hcl2fmt\n")
	flag.PrintDefaults()
	os.Exit(2)
}
