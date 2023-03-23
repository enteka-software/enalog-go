# EnaLog Go

### Installation

```bash
go get github.com/enteka-software/enalog-go
```

### Usage

```go
package main

import (
    "os"
    "github.com/enteka-software/enalog-go"   
)

func main() {
    client, _ := enalog.New(os.Getenv("ENALOG_API_KEY"))
    
    event := enalog.Event{
        Project: "landing-page",
        Name: "user-joined-waitlist",
        Push: false,
        Description: "User joined waitlist on the landing page",
        Icon: "🚀",
        Tags: []string{},
        Meta: map[string]string{}
    }
    
    _, _ := client.PushEvent(event)
}
```