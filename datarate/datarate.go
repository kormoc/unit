package datarate

import "fmt"
import "github.com/kormoc/unit"
import "github.com/kormoc/unit/datasize"
import "strings"
import "time"

type Datarate float64
type DatarateSIBit   Datarate
type DatarateSIByte  Datarate
type DatarateIECBit  Datarate
type DatarateIECByte Datarate

var outputStringMaxPercision int = 3

var SIBits = [...]unit.Unit {
    {Name: "yottabit per second", Size: float64(datasize.Yottabit), Suffix: "Ybit/s"},
    {Name: "zettabit per second", Size: float64(datasize.Zettabit), Suffix: "Zbit/s"},
    {Name:   "exabit per second", Size: float64(datasize.Exabit  ), Suffix: "Ebit/s"},
    {Name:  "petabit per second", Size: float64(datasize.Petabit ), Suffix: "Pbit/s"},
    {Name:  "terabit per second", Size: float64(datasize.Terabit ), Suffix: "Tbit/s"},
    {Name:  "gigabit per second", Size: float64(datasize.Gigabit ), Suffix: "Gbit/s"},
    {Name:  "megabit per second", Size: float64(datasize.Megabit ), Suffix: "Mbit/s"},
    {Name:  "kilobit per second", Size: float64(datasize.Kilobit ), Suffix: "kbit/s"},
    {Name:      "bit per second", Size: float64(datasize.Bit     ), Suffix: "bit/s"},
}

var IECBits = [...]unit.Unit {
    {Name: "yobibit per second", Size: float64(datasize.Yobibit), Suffix: "Yibit/s"},
    {Name: "zebibit per second", Size: float64(datasize.Zebibit), Suffix: "Zibit/s"},
    {Name: "exbibit per second", Size: float64(datasize.Exbibit), Suffix: "Eibit/s"},
    {Name: "pebibit per second", Size: float64(datasize.Pebibit), Suffix: "Pibit/s"},
    {Name: "tebibit per second", Size: float64(datasize.Tebibit), Suffix: "Tibit/s"},
    {Name: "gibibit per second", Size: float64(datasize.Gibibit), Suffix: "Gibit/s"},
    {Name: "mebibit per second", Size: float64(datasize.Mebibit), Suffix: "Mibit/s"},
    {Name: "kibibit per second", Size: float64(datasize.Kibibit), Suffix: "Kibit/s"},
    {Name:     "bit per second", Size: float64(datasize.Bit    ), Suffix: "bit/s"},
}

var SIBytes = [...]unit.Unit {
    {Name: "yottabyte per second", Size: float64(datasize.Yottabyte), Suffix: "YB/s"},
    {Name: "zettabyte per second", Size: float64(datasize.Zettabyte), Suffix: "ZB/s"},
    {Name:   "exabyte per second", Size: float64(datasize.Exabyte  ), Suffix: "EB/s"},
    {Name:  "petabyte per second", Size: float64(datasize.Petabyte ), Suffix: "PB/s"},
    {Name:  "terabyte per second", Size: float64(datasize.Terabyte ), Suffix: "TB/s"},
    {Name:  "gigabyte per second", Size: float64(datasize.Gigabyte ), Suffix: "GB/s"},
    {Name:  "megabyte per second", Size: float64(datasize.Megabyte ), Suffix: "MB/s"},
    {Name:  "kilobyte per second", Size: float64(datasize.Kilobyte ), Suffix: "kB/s"},
    {Name:      "byte per second", Size: float64(datasize.Byte     ), Suffix: "B/s"},
}

var IECBytes = [...]unit.Unit {
    {Name: "yobibyte per second", Size: float64(datasize.Yobibyte), Suffix: "YiB/s"},
    {Name: "zebibyte per second", Size: float64(datasize.Zebibyte), Suffix: "ZiB/s"},
    {Name: "exbibyte per second", Size: float64(datasize.Exbibyte), Suffix: "EiB/s"},
    {Name: "pebibyte per second", Size: float64(datasize.Pebibyte), Suffix: "PiB/s"},
    {Name: "tebibyte per second", Size: float64(datasize.Tebibyte), Suffix: "TiB/s"},
    {Name: "gibibyte per second", Size: float64(datasize.Gibibyte), Suffix: "GiB/s"},
    {Name: "mebibyte per second", Size: float64(datasize.Mebibyte), Suffix: "MiB/s"},
    {Name: "kibibyte per second", Size: float64(datasize.Kibibyte), Suffix: "KiB/s"},
    {Name:     "byte per second", Size: float64(datasize.Byte    ), Suffix: "B/s"},
}

func NewDatarate(datasize datasize.Datasize, duration time.Duration) Datarate {
    return Datarate(float64(datasize) / duration.Seconds())
}

func SetOutputStringMaxPercision(maxPercision int) {
    outputStringMaxPercision = maxPercision

}

func buildHumanString(value float64, mapping [9]unit.Unit) (dataRate string) {
    for _, v := range mapping {
        if value >= v.Size {
            // Last level to process, display a float
            vals := strings.SplitN(fmt.Sprintf("%f", value / v.Size), ".", 2)
            dataRate += vals[0]
            dataRate += strings.TrimRight("."+vals[1][:outputStringMaxPercision], ".0")
            dataRate += v.Suffix
            break
        }
    }
    return dataRate
}

func (d DatarateSIBit) String() string {
    return buildHumanString(float64(d), SIBits)
}

func (d DatarateSIByte) String() string {
    return buildHumanString(float64(d), SIBytes)
}

func (d DatarateIECBit) String() string {
    return buildHumanString(float64(d), IECBits)
}

func (d DatarateIECByte) String() string {
    return buildHumanString(float64(d), IECBytes)
}

// Helper consts

const (
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
    Hour        = time.Hour
)

const (
    Bit      = datasize.Bit
    Kilobit  = datasize.Kilobit
    Megabit  = datasize.Megabit
    Gigabit  = datasize.Gigabit
    Terabit  = datasize.Terabit
    Petabit  = datasize.Petabit
    Exabit   = datasize.Exabit
    Zettabit = datasize.Zettabit
    Yottabit = datasize.Yottabit
)

const (
    Kibibit = datasize.Kibibit
    Mebibit = datasize.Mebibit
    Gibibit = datasize.Gibibit
    Tebibit = datasize.Tebibit
    Pebibit = datasize.Pebibit
    Exbibit = datasize.Exbibit
    Zebibit = datasize.Zebibit
    Yobibit = datasize.Yobibit
)

const (
    Byte       = datasize.Byte
    Kilobyte   = datasize.Kilobyte
    Megabyte   = datasize.Megabyte
    Gigabyte   = datasize.Gigabyte
    Terabyte   = datasize.Terabyte
    Petabyte   = datasize.Petabyte
    Exabyte    = datasize.Exabyte
    Zettabyte  = datasize.Zettabyte
    Yottabyte  = datasize.Yottabyte
)

const (
    Kibibyte = datasize.Kibibyte
    Mebibyte = datasize.Mebibyte
    Gibibyte = datasize.Gibibyte
    Tebibyte = datasize.Tebibyte
    Pebibyte = datasize.Pebibyte
    Exbibyte = datasize.Exbibyte
    Zebibyte = datasize.Zebibyte
    Yobibyte = datasize.Yobibyte
)
