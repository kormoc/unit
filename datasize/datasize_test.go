package datasize

import "testing"

func Test(t *testing.T) {

    datasize := Datasize(500) * Megabyte

    if DatasizeSIBit(datasize).String() != "4Gbit" {
        t.Fatalf("DatasizeSIBit has failed\n\tExpected: %v\n\tFound: %v\n", "4Gbit", DatasizeSIBit(datasize).String())
    }

    if DatasizeSIByte(datasize).String() != "500M" {
        t.Fatalf("DatasizeSIByte has failed\n\tExpected: %v\n\tFound: %v\n", "500M", DatasizeSIByte(datasize).String())
    }

    if DatasizeIECBit(datasize).String() != "3.725Gibit" {
        t.Fatalf("DatasizeIECBit has failed\n\tExpected: %v\n\tFound: %v\n", "3.725Gibit", DatasizeIECBit(datasize).String())
    }

    if DatasizeIECByte(datasize).String() != "476.837Mi" {
        t.Fatalf("DatasizeIECByte has failed\n\tExpected: %v\n\tFound: %v\n", "476.837Mi", DatasizeIECByte(datasize).String())
    }
}

func TestOutputStringMaxPercision(t *testing.T) {
    datasize := Datasize(500) * Megabyte
    SetOutputStringMaxPercision(0)

    if DatasizeIECByte(datasize).String() != "476Mi" {
        t.Fatalf("DatasizeIECByte has failed\n\tExpected: %v\n\tFound: %v\n", "476Mi", DatasizeIECByte(datasize).String())
    }
}

func TestOutputStringMaxLevels(t *testing.T) {
    datasize := Datasize(500) * Megabyte
    SetOutputStringMaxLevels(2)

    if DatasizeIECByte(datasize).String() != "476Mi857Ki" {
        t.Fatalf("DatasizeIECByte has failed\n\tExpected: %v\n\tFound: %v\n", "476Mi857Ki", DatasizeIECByte(datasize).String())
    }
}
