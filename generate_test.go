package gobatis_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aryann/difflib"
)

func TestGenerate(t *testing.T) {
	wd := getGoBatis()

	for _, cmd := range []*exec.Cmd{
		exec.Command("go", "install", "github.com/runner-mei/GoBatis/cmd/gobatis"),
		exec.Command("go", "generate", "github.com/runner-mei/GoBatis/gentest"),
	} {
		cmd.Dir = wd
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Logf("%s", out)
			t.Error(err)
			return
		}
	}

	for _, name := range []string{"user", "role", "users"} {
		t.Log("=====================", name)
		actual := readFile(filepath.Join(wd, "gentest/"+name+".gobatis.go"))
		excepted := readFile(filepath.Join(wd, "gentest/"+name+".gobatis.txt"))
		if !reflect.DeepEqual(actual, excepted) {
			results := difflib.Diff(excepted, actual)
			for _, result := range results {
				t.Error(result)
			}
		}
	}
}

func readFile(pa string) []string {
	bs, err := ioutil.ReadFile(pa)
	if err != nil {
		panic(err)
	}

	return splitLines(bs)
}

func splitLines(txt []byte) []string {
	//r := bufio.NewReader(strings.NewReader(s))
	s := bufio.NewScanner(bytes.NewReader(txt))
	var ss []string
	for s.Scan() {
		ss = append(ss, s.Text())
	}
	return ss
}
