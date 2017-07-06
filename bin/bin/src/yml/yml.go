package yml

import (
    "bufio"
    "log"
    "os"
    "strings"
    "regexp"
    "fmt"
    "os/user"
)

type Yml struct {
  Env map[string]string
  Filepath string
}

func New(f string) (yml * Yml) {
    if f == "" {
        f = GetHome() + "/.env.yml"
    }
    return &Yml{
        Filepath: f,
        Env: ReadFile(f),
    }
}

func (y *Yml) Set(k,v string) {
    y.Env[k] = v
}

func (y Yml) Write(f string) {
  fileHandle, _ := os.Create(f)
  writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
  for k, v := range y.Env {
    fmt.Fprintln(writer, k + ": " + v)
    writer.Flush()
  }
}

func ReadFile(f string)(m map[string]string) {
    file, err := os.Open(f)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    m = make(map[string]string)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        l := scanner.Text()
        reg := regexp.MustCompile(":+?[^http://]")
        //reg2 := regexp.MustCompile("s/^\([\"']\)\(.*\)\1\$/\2/g")
        r := reg.Split(l, 2)
        k, v := strings.TrimSpace(r[0]), strings.TrimSpace(r[1])
        v = strings.Replace(v, "\"", "", -1)
        m[k] = v
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return
}

func GetHome() (h string) {
    usr, err := user.Current()
    if err != nil {
        log.Fatal(err)
    }
    h = usr.HomeDir
    return
}
