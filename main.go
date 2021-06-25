package main

import(
    "flag"
    "time"
    "fmt"
    "os"
    "strconv"
    "github.com/TRIKKSS/GoHTTPBASICBrute/utils"
)

var attempt = 0
var start time.Time
var url *string

func worker(id int, username <-chan string, password <-chan string) {
    for j := range username {
        for u := range password {
            attempt++
            if attempt % 1000 == 0 {
                fmt.Println("attempt :", attempt, "\ntime :", time.Since(start))
            }
            if utils.RequestBasicAuth(url, &j, &u) == true {
                KeyFound(j,u)
            }
        }
    }
}

func KeyFound(username string, password string) {
    fmt.Println("key found !")
    fmt.Printf("creds -> %s:%s\n", username, password)
    fmt.Println("program took", time.Since(start))
    os.Exit(0)
}

func main() {
    start = time.Now()
    var usernameList []string
    var passwordList []string

    usernameChan := make(chan string, 100000)
    passwordChan := make(chan string, 100000)

    username := flag.String("U", "", "username")
    password := flag.String("P", "", "password")
    passwordlistName := flag.String("p", "", "password list")
    usernameListName := flag.String("u", "", "username list")
    url = flag.String("url", "", "target url")
    threads := flag.String("t", "10", "threads")
    flag.Parse()

    threadsInt, err := strconv.Atoi(*threads)
    if err != nil {
        fmt.Println("threads must be integer !")
        os.Exit(1)
    }

    if *username == "" && *usernameListName == "" {
        utils.Usage()
    } else if *password == "" && *passwordlistName == "" {
        utils.Usage()
    } else if *url == "" {
        utils.Usage()
    } else {
        if *usernameListName != "" {
            utils.ReadWordlist(*usernameListName, &usernameList)
        } else {
            usernameList = append(usernameList, *username)
        }
        if *passwordlistName != "" {
            utils.ReadWordlist(*passwordlistName, &passwordList)
        } else {
            passwordList = append(passwordList, *password)
        }
    }

    utils.Begin(url, &usernameList, usernameListName, &passwordList, passwordlistName, threads)

    for w := 1; w <= threadsInt; w++ {
        go worker(w, usernameChan, passwordChan)
    }

    for _, users := range usernameList {
        usernameChan <- users
        for _, i := range passwordList {
            passwordChan <- i
        }
    }
}