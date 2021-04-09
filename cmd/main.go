package main

import (
    "fmt"
    "github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"

    "hsdevice"
    "hsdevice/cmd/driver"
)

func main() {
    sd := driver.HansongDevice{}
    fmt.Printf("the type of driver is %T\n", sd)

    fmt.Printf("Hello world, with version: %s\n", hsdevice.Version)
    startup.Bootstrap("hansong-device",
        "1.0",
        &sd)
}
