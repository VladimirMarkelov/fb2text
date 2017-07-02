package main

import (
	"flag"
	"fmt"
	"github.com/VladimirMarkelov/fb2text"
	"os"
)

func main() {
	textWidth := flag.Int("w", 70, "width(30-400) width of the text for output file")
	justify := flag.Bool("j", false, "justify(0/1 or true/false) - add spaces between words to make a all text lines to be the same width (except the last line of every paragraph)")
	flag.Parse()

	dstName := flag.Arg(1)
	srcName := flag.Arg(0)

	if *textWidth < 30 || *textWidth > 400 || srcName == "" {
		fmt.Printf("Usage:\n\n    fb2txt [-j=1] [-w N] <source-filename> [<output-filename>]\n\n")
		fmt.Printf("-w   [30-400] maximum width of the text for converter, defaut is 70\n\n")
		fmt.Printf("-j   [0/1] or [true/false] - add spaces between words to make a all text lines to be the same width (except the last line of every paragraph), disabled by default\n")
		return
	}

	if dstName != "" {
		fmt.Printf("Opening file %s to convert...\n", srcName)
	}
	binfo, lines := fb2text.ParseBook(srcName, true)
	lines = fb2text.FormatBook(lines, *textWidth, *justify)

	if dstName != "" {
		f, _ := os.Create(dstName)
		f.WriteString(fmt.Sprintf("%s %s\n", binfo.FirstName, binfo.LastName))
		if binfo.Sequence != "" {
			f.WriteString(fmt.Sprintf("%s (%s)\n\n", binfo.Title, binfo.Sequence))
		} else {
			f.WriteString(fmt.Sprintf("%s\n\n", binfo.Title))
		}
		for _, ll := range lines {
			f.WriteString(ll)
			f.WriteString("\n")
		}
		f.Close()
	} else {
		fmt.Printf("%s %s\n", binfo.FirstName, binfo.LastName)
		if binfo.Sequence != "" {
			fmt.Printf(fmt.Sprintf("%s (%s)\n\n", binfo.Title, binfo.Sequence))
		} else {
			fmt.Printf(fmt.Sprintf("%s\n\n", binfo.Title))
		}
		for _, ll := range lines {
			fmt.Println(ll)
		}
	}

	if dstName != "" {
		fmt.Printf("Book lines: %d\n", len(lines))
		fmt.Printf("   Author: %s %s\n", binfo.FirstName, binfo.LastName)
		fmt.Printf("   Title: [%s] %s (%s)\n", binfo.Genre, binfo.Title, binfo.Sequence)
		fmt.Printf("Converted and saved to %s\n", dstName)
	}
}
