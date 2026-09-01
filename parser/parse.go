package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"todo/errdef"
	"todo/parser/tokens"
)

/*
folder is a working directory, ext is the extension for the search
returns list of handles and a formatted error
IMPORTANT - call this with a FULL path from root or from homedir
homedir path starts with '~' (a tilda)
*/
func OpenFiles(folder string, ext string) ([]*os.File, error) {
	filelist := make([]*os.File, 0)
	folderSep := ([]rune(folder))[0]

	if folderSep == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			err = errors.New(errdef.ErrBase + "cannot find user home dir")
			return filelist, err
		}
		folder = homeDir + folder[1:]
	}
	err := filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == ".git" {
			return fs.SkipDir
		}

		file, err := os.Open(path)
		if err != nil {
			fmt.Print(errdef.WarnBase + "unable to open file: " + path)
		}
		if filepath.Ext(path) == ext {
			filelist = append(filelist, file)
		}
		return nil
	})
	if err != nil {
		err = errors.New(errdef.ErrBase + "error in walking the named dir")
		return filelist, err
	}
	return filelist, nil
}

func parseTodos(files []*os.File, folder string) {
	compositePath := folder + "./tasksummary.txt"
	logfile, err := os.OpenFile(compositePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0640)
	var str string
	for _, file := range files {
		scanner := bufio.NewScanner(file)
		logfile.WriteString("[" + file.Name() + "]\n")
		var line = 0
		for scanner.Scan() {
			line++
			str = scanner.Text()

		}
	}
}
// [ -> blabla -> ]
func checkTokens(line string) error {
	line = strings.TrimSpace(line)
	lastToken := 0
	lineToRunes := []rune(line)

	if !strings.Contains(line, "//") || !strings.Contains(line, "/*") {
		return nil
	}

	for index, char := range lineToRunes {
		if char == ' '{
			continue
		}
		if 
		
	}

}
