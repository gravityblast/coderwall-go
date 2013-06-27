# Coderwall

Coderwall API client for [go](http://golang.org/ "Go lang").

## Installation

    go get github.com/pilu/coderwall-go

## Usage

    package main

    import (
      "github.com/pilu/coderwall-go"
      "fmt"
      "flag"
      "os"
    )

    func usage() {
      fmt.Fprintf(os.Stderr, "usage: coderwall USERNAME\n")
      flag.PrintDefaults()
      os.Exit(1)
    }

    func main() {
      flag.Usage = usage
      flag.Parse()

      args := flag.Args()
      if len(args) < 1 {
        usage()
        os.Exit(1);
      }
      client := coderwall.NewClient()
      profile, err := client.GetProfile(args[0])
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      fmt.Printf("%s, %s (%s) - %d endorsement\n\n", profile.Username, profile.Name, profile.Location, profile.Endorsements)
      fmt.Printf("Badges (%d):\n\n", len(profile.Badges))
      for _, badge := range(profile.Badges) {
        fmt.Printf("\t%s\n\t\t%s\n\n", badge.Name, badge.Description)
      }
    }
