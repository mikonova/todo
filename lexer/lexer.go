package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"todo/errdef"
	"todo/lexer/tokens"
)

const StorageFileName = "todos.md"

type TokenStorage struct {
	TokenStream []string
	FileName    string
}

/*
folder is a working directory, ext is the extension for the search
returns list of handles and a formatted error
IMPORTANT - call this with a FULL path from root or from homedir
homedir path starts with '~' (a tilda), root (absolute path) starts with "/" (slash)
In Windows, full path starts with a drive letter
*/
func OpenFiles(folder string, ext string) ([]*os.File, error) {
	if err := os.Chdir(folder); err != nil {
		log.Fatalln(errdef.ErrBase + "Cannot change into user directory")
	}
	if file, err := os.Create(StorageFileName); err != nil {
		log.Fatalln(errdef.ErrBase + "Cannot create the summary file")
	} else {
		file.Close()
	}
	fileList := make([]*os.File, 0)
	folderSep := ([]rune(folder))[0]

	if folderSep == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			err = errors.New(errdef.ErrBase + "cannot find user home dir")
			return fileList, err
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
			fileList = append(fileList, file)
		}
		return nil
	})
	if err != nil {
		err = errors.New(errdef.ErrBase + "error in walking the named dir")
		return fileList, err
	}
	return fileList, nil
}

/*
Tokenize lines of all the files. Takes the result of OpenFiles and a root directory name.
Safe to call in general, but scanner might cause panic in case of an internal error
*/
func GetLines(files []*os.File, folder string) (storage []TokenStorage) {
	var globalDepth = 0
	for _, file := range files {
		tempStorage := TokenStorage{
			FileName: file.Name(),
		}
		scanner := bufio.NewScanner(file)
		var line = 0
		for scanner.Scan() {
			var tokenList []string
			checkScannerError(scanner.Err())
			line++
			tokenList, globalDepth = tokenize(scanner.Text(), globalDepth)
			tempStorage.TokenStream = append(tempStorage.TokenStream, tokenList...)
		}
		storage = append(storage, tempStorage)
		file.Close()
	}
	return storage
}

func checkScannerError(err error) {
	if err != nil {
		log.Panicln(errdef.ErrBase + "scanner error in the tokenizer! Panicking")
	}
}

func tokenize(line string, depth int) ([]string, int) {
	line = strings.TrimSpace(line)
	tokenSlice := make([]string, 0)
	newDepth := depth
	var buf string
	lineToRunes := []rune(line)

	if !strings.Contains(line[:2], "//") || !strings.Contains(line[:2], "/*") {
		return tokenSlice, newDepth
	}

	for index, char := range lineToRunes {
		if lineToRunes[0] == tokens.LineStart && lineToRunes[1] == tokens.Star {
			depth += 1
		}
		if lineToRunes[0] == tokens.Star && lineToRunes[1] == tokens.LineStart {
			depth -= 1
		}
		if !(char == tokens.Space || char == tokens.Comma) {
			buf += string(char)
		}
		if char == tokens.Space {
			tokenSlice = append(tokenSlice, buf)
			buf = ""
			continue
		}
		if lineToRunes[index-1] == tokens.RightBrace {
			buf = line[index:]
			buf += "\n"
		}
		if !strings.Contains(line, string(tokens.LeftBrace)) && !strings.Contains(line, string(tokens.RightBrace)) &&
			lineToRunes[index-1] == tokens.Colon {
			buf = line[index:]
			buf += "\n"
		}
	}
	return tokenSlice, newDepth

}
