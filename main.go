// BSD 3-Clause License
//
// Copyright (c) 2022, kaba kaba-tech@outlook.com
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	aw "github.com/deanishe/awgo"

	str2duration "github.com/xhit/go-str2duration/v2"
)

var (
	alfredMode bool
	tsArgs     []string
	wf         *aw.Workflow
)

const Usage = `
Usage: ts <duration>
  duration:
	+1w: add 1 weeks
	+2d: add 2 days
	-1h: sub 1 hours
	-2m: sub 2 minutes
	-3s: sub 3 seconds
	+4ms: add 4 milliseconds
	+5us: add 5 microseconds
	+5µs: add 5 microseconds
	+6ns: add 6 nanoseconds
	+1d12h30m: add 1 days, 12 hours and 30 minutes
`

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "alfred" {
		// alfred workflow mode
		alfredMode = true
		tsArgs = args[1:]
		wf = aw.New()
		wf.Run(run)
	} else if len(args) == 1 {
		// show usage
		fmt.Print(Usage)
	} else {
		// command line mode
		tsArgs = args
		run()
	}
}

////////////////////////////////////////////////////////////////

func run() {
	now := time.Now()

	switch len(tsArgs) {
	case 1:
		defaultFeedback(now)
	case 2:
		operator := tsArgs[1][0]
		if operator == '-' || operator == '+' {
			d, err := str2duration.ParseDuration(tsArgs[1][1:])
			if err != nil {
				defaultFeedback(now)
			} else {
				if operator == '+' {
					defaultFeedback(now.Add(d))
				} else {
					defaultFeedback(now.Add(-d))
				}
			}
		} else {
			defaultFeedback(now)
		}
	default:
		defaultFeedback(now)
	}

	if alfredMode {
		wf.SendFeedback()
	}
}

func defaultFeedback(now time.Time) {
	timeString := []feedback{
		{
			time:     strconv.FormatInt(now.Unix(), 10),
			subtitle: "int64",
			arg:      strconv.FormatInt(now.Unix(), 10),
		},
		{
			time:     now.Format("2006-01-02"),
			subtitle: "2006-01-02",
			arg:      now.Format("2006-01-02"),
		},
		{
			time:     now.Format("2006-01-02 15:04:05"),
			subtitle: "2006-01-02 15:04:05",
			arg:      now.Format("2006-01-02 15:04:05"),
		},
		{
			time:     now.Format(time.RFC3339),
			subtitle: time.RFC3339,
			arg:      now.Format(time.RFC3339),
		},
		{
			time:     now.Format("Mon, 02 Jan 2006 15:04:05"),
			subtitle: "Mon, 02 Jan 2006 15:04:05",
			arg:      now.Format("Mon, 02 Jan 2006 15:04:05"),
		},
	}

	for i := range timeString {
		if alfredMode {
			wf.NewItem(timeString[i].time).
				Subtitle(timeString[i].subtitle).
				Arg(timeString[i].arg).
				Valid(true).
				Icon(aw.IconWorkflow)
		} else {
			fmt.Println(timeString[i].time)
		}
	}
	if alfredMode {
		wf.NewWarningItem("ts +1d12h is add 1 days, 12 hours", "Available options w:week d:day h:hour m:minute s:second ms:millisecond us:microsecond ns:nanosecond").
			Arg("ts +1d12h is add 1 days, 12 hours\nAvailable options w:week d:day h:hour m:minute s:second ms:millisecond us:microsecond ns:nanosecond").
			Valid(true).
			Icon(aw.IconHelp)
	}
}

type feedback struct {
	time     string
	subtitle string
	arg      string
}
