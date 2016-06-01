package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
    "strings"
    "regexp"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func main() {

    // get command line arguments
    // expected: name of file
    var filename string
    var sourceOrDest string

    if(len(os.Args)<3) {
        fmt.Println("Usage: [src|dest|both] [source filename]")
        return
    } else {
        filename = os.Args[2]
        sourceOrDest = os.Args[1]
    }

    //lines, err := readLines("/users/dfranke/Documents/code/g1/ips.txt")
    lines, err := readLines(filename)
    var ipaddr string 
    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }
    // display contents
    var spl string
    spl = "index=firewall ("
    var i = 0
    for _, line := range lines {
    	ipaddr = line
        
        // desanitize
        r := regexp.MustCompile(`\[[\.\,]\]`)
        ipaddr = r.ReplaceAllString(ipaddr,".")
    	
        // remove trailing port info
        r = regexp.MustCompile(`\:\d+`)
        ipaddr = r.ReplaceAllString(ipaddr,"")

        // remove meows
        r = regexp.MustCompile(`h[tx]+p:\/\/`)
        ipaddr = r.ReplaceAllString(ipaddr,"")

        // trim whitespace
        ipaddr = strings.Trim(ipaddr," ")

        // build string based on sourceOrDest entry provided by user
        //   src_ip=x
        //   dest_ip=x
        //   src_ip OR dest_ip=x
        if i>0 {
            spl += " OR "
        }

        if(sourceOrDest=="src") {
            spl += "src_ip=" + ipaddr
        }
        if(sourceOrDest=="dest"||sourceOrDest=="dst") {
            spl += "dest_ip=" + ipaddr
        }
        if(sourceOrDest=="both") {
            spl += "(src_ip=" + ipaddr + " OR dest_ip=" + ipaddr + ")"
        }

    	i++
        fmt.Printf("IP Address: %s\n", ipaddr)
    }
    spl += ")"
    fmt.Println(strings.Repeat("=", 30) + " SNIP " + strings.Repeat("=", 30))
    fmt.Println(spl)
    fmt.Println(strings.Repeat("=", 30) + " /SNIP " + strings.Repeat("=", 30))

}
