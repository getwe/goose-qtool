package main

import (
    "fmt"
    "os"
    "net"
    "bytes"
    "time"
    "io/ioutil"
    flags "github.com/jessevdk/go-flags"
)

func main() {
    defer func() {
        if r := recover();r != nil {
            fmt.Println(r)
            os.Exit(1)
        }
    }()

    var opts struct {

        Ip string `short:"i" long:"ip" description:"server ip"`

        Port string `short:"p" long:"port" description:"server port"`

        Cmd string `short:"c" long:"cmd" description:"query command"`

    }
    parser := flags.NewParser(&opts,flags.HelpFlag)
    _,err := parser.ParseArgs(os.Args)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if 0 == len(opts.Ip) || 0 == len(opts.Port) || 0 == len(opts.Cmd) {
        parser.WriteHelp(os.Stderr)
        os.Exit(1)
    }
    work(opts.Ip,opts.Port,opts.Cmd)
}

func work(ip,port,cmd string) {

    conn,err := net.DialTimeout("tcp",fmt.Sprintf("%s:%s",ip,port),time.Second)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    reqbuf := bytes.NewBufferString(cmd)
    writeLen,err := conn.Write(reqbuf.Bytes())
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("write data len[%d]\n",writeLen)

    resbytes,err := ioutil.ReadAll(conn)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    resbuf := bytes.NewBuffer(resbytes)
    fmt.Println(resbuf.String())
}
