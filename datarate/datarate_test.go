package datarate

import "testing"
import "time"
import "github.com/kormoc/unit/datasize"

func Test(t *testing.T) {
    datarate := NewDatarate(500 * datasize.Megabyte, 17*time.Second)

    if DatarateSIBit(datarate).String() != "235.294Mbit/s" {
        t.Fatalf("DatarateSIBit has failed\n\tExpected: %v\n\tFound: %v\n", "235.294Mbit/s", DatarateSIBit(datarate).String())
    }

    if DatarateSIByte(datarate).String() != "29.411MB/s" {
        t.Fatalf("DatarateSIByte has failed\n\tExpected: %v\n\tFound: %v\n", "29.411MB/s", DatarateSIByte(datarate).String())
    }

    if DatarateIECBit(datarate).String() != "224.393Mibit/s" {
        t.Fatalf("DatarateIECBit has failed\n\tExpected: %v\n\tFound: %v\n", "224.393Mibit/s", DatarateIECBit(datarate).String())
    }

    if DatarateIECByte(datarate).String() != "28.049MiB/s" {
        t.Fatalf("DatarateIECByte has failed\n\tExpected: %v\n\tFound: %v\n", "28.049MiB/s", DatarateIECByte(datarate).String())
    }
}

func TestOutputStringMaxPercision(t *testing.T) {
    datarate := NewDatarate(500 * datasize.Megabyte, 17*time.Second)
    SetOutputStringMaxPercision(0)

    if DatarateIECByte(datarate).String() != "28MiB/s" {
        t.Fatalf("DatarateIECByte has failed\n\tExpected: %v\n\tFound: %v\n", "28MiB/s", DatarateIECByte(datarate).String())
    }
}