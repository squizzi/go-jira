package jira

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

func readJSON(input io.Reader, data interface{}) error {
	content, err := ioutil.ReadAll(input)
	if err != nil {
		return err
	}
	if len(content) == 0 {
		os.Exit(0)
	}
	re := regexp.MustCompile(`issues\/\d*`)
	match := re.FindString(string(content))
	if len(match) == 0 {
		os.Exit(0)
	}
	matches := strings.Split(match, "/")
	fmt.Println(matches[1])
	os.Exit(0)
	return nil
}

func URLJoin(endpoint string, paths ...string) string {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(fmt.Errorf("Unable to parse endpoint: %s", endpoint))
	}
	paths = append([]string{u.Path}, paths...)
	u.Path = path.Join(paths...)
	return u.String()
}
