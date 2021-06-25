package utils

import(
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
    "net/http"
)

func Begin(url *string, username *[]string, usersWordlist *string, passwd *[]string, passWordlist *string, threads *string) {
    fmt.Println("\n====================================================")
    fmt.Printf("target : %s\n", *url)
    fmt.Println("====================================================")
    if len(*username) == 1 {
        fmt.Printf("username : %s\n", (*username)[0])
    } else {
        fmt.Printf("username list : %s\n", *usersWordlist)
    }
    fmt.Println("====================================================")
    if len(*passwd) == 1 {
        fmt.Printf("password : %s\n", (*passwd)[0])
    } else {
        fmt.Printf("password list : %s\n", *passWordlist)
    }
    fmt.Println("====================================================")
    fmt.Printf("threads : %s\n", *threads)
    fmt.Println("====================================================\n")
}


func ReadWordlist(wordlistFile string, wordlist *[]string) {
    file, err := os.Open(wordlistFile)

    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        *wordlist = append(*wordlist, scanner.Text())
    }
    file.Close()
}


func CheckUrl(url string) {
    if strings.Split(url, ":")[0] != "http" {
        fmt.Println("invalid protocol : must be http")
        os.Exit(1)
    }
}


func RequestBasicAuth(url *string, username *string, passwd *string) bool {
    var resp *http.Response
    for true {
        var err error
        client := &http.Client{}
        req, _ := http.NewRequest("GET", *url, nil)
        req.SetBasicAuth(*username, *passwd)
        resp, err = client.Do(req)
        if err == nil{
            break
        }
        fmt.Println(err)
    }
    if resp.StatusCode == 401 {
        return false
    }
    return true
}


func Usage() {
    fmt.Printf("usage : %s [option]\n\t-u username list\n\t-p password list\n\t-U username\n\t-P password\n\t-url target url\n\t-t threads (default 10)\n", os.Args[0])
    os.Exit(1)
}