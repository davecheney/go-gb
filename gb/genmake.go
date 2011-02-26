package main

import (
	"template"
)

var MakeCmdTemplate = template.MustParse(
`# Makefile generated by gb: http://go-gb.googlecode.com
# gb provides configuration-free building and distributing

include $(GOROOT)/src/Make.inc

TARG={Target}
GOFILES=\
{.repeated section GoFiles}	{@}\
{.end}

# gb: this is the local install
GBROOT={GBROOT}

# gb: compile/link against local install
GC+= -I $(GBROOT)/_obj
LD+= -L $(GBROOT)/_obj

# gb: default target is in GBROOT this way
command:

include $(GOROOT)/src/Make.cmd

# gb: copy to local install
$(GBROOT)/{BuildDirCmd}/$(TARG): $(TARG)
	mkdir -p $(dir $@); cp -f $< $@
command: $(GBROOT)/bin/$(TARG)
{.section LocalDeps}

# gb: local dependencies
{.repeated section LocalDeps}$(TARG): $(GBROOT)/{BuildDirPkg}/{@}.a
{.end}
{.end}
`,
	nil)

var MakePkgTemplate = template.MustParse(
`# Makefile generated by gb: http://go-gb.googlecode.com
# gb provides configuration-free building and distributing

include $(GOROOT)/src/Make.inc

TARG={Target}
GOFILES=\
{.repeated section GoFiles}	{@}\
{.end}
{.section AsmObjs}

OFiles=\
{.repeated section AsmObjs}	{@}\
{.end}
{.end}
{.section CGoFiles}

CGOFILES=\
{.repeated section CGoFiles}	{@}\
{.end}
{.end}
{.section CObjs}

CGO_OFILES=\
{.repeated section CObjs}	{@}\
{.end}
{.end}

# gb: this is the local install
GBROOT={GBROOT}

# gb: compile/link against local install
GC+= -I $(GBROOT)/{BuildDirPkg}
LD+= -L $(GBROOT)/{BuildDirPkg}

{.section CopyLocal}
# gb: copy to local install
$(GBROOT)/{BuildDirPkg}/$(TARG).a: {BuildDirPkg}/$(TARG).a
	mkdir -p $(dir $@); cp -f $< $@
{.end}
package: $(GBROOT)/{BuildDirPkg}/$(TARG).a

include $(GOROOT)/src/Make.pkg
{.section LocalDeps}

# gb: local dependencies
{.repeated section LocalDeps}{BuildDirPkg}/$(TARG).a: $(GBROOT)/{BuildDirPkg}/{@}.a
{.end}
{.end}
`,
	nil)

type MakeData struct {
	Target      string
	GBROOT      string
	GoFiles     []string
	AsmObjs     []string
	CGoFiles    []string
	CObjs       []string
	LocalDeps   []string
	BuildDirPkg string
	BuildDirCmd string
	CopyLocal   bool
}
