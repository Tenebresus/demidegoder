package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

var typeMapping = map[int]string {

    0: "BIOS",
    1: "System",
    2: "Base Board",
    3: "Chassis",
    4: "Processor",
    5: "Memory Controller",
    6: "Memory Module",
    7: "Cache",
    8: "Port Connector",
    9: "System Slots",
    10: "On Board Devices",
    11: "OEM Strings",
    12: "System Configuration Options",
    13: "BIOS Language",
    14: "Group Associations",
    15: "System Event Log",
    16: "Physical Memory Array",
    17: "Memory Device",
    18: "32-bit Memory Error",
    19: "Memory Array Mapped Address",
    20: "Memory Device Mapped Address",
    21: "Built-in Pointing Device",
    22: "Portable Battery",
    23: "System Reset",
    24: "Hardware Security",
    25: "System Power Controls",
    26: "Voltage Probe",
    27: "Cooling Device",
    28: "Temperature Probe",
    29: "Electrical Current Probe",
    30: "Out-of-band Remote Access",
    31: "Boot Integrity Services",
    32: "System Boot",
    33: "64-bit Memory Error",
    34: "Management Device",
    35: "Management Device Component",
    36: "Management Device Threshold Data",
    37: "Memory Channel",
    38: "IPMI Device",
    39: "Power Supply",
    40: "Additional Information",
    41: "Onboard Device",

}

type Parser struct {

    Dmidecode string
    dmiTypes []DMIType

}


func Parse(dmidecode string) []byte {

    p := Parser {

        Dmidecode: dmidecode,

    }

    matches := p.getMatches()

    for _, match := range matches {

        dmiType := typeMapping[getDMIType(match)]

        if dmiType != "" {

            properties := getProperties(match) 

            dmitype := DMIType {

                Name: dmiType,
                Properties: properties,

            }

            p.dmiTypes = append(p.dmiTypes, dmitype)

        }

    }

    ret, _ := json.Marshal(p.dmiTypes)
    return ret

}

func getProperties(match string) map[string]string {

    regex := regexp.MustCompile(`(?m)^\t([A-z].+): (.+)$`)
    matches := regex.FindAllStringSubmatch(match, -1)

    properties := make(map[string]string)

    for _, property := range matches {

        properties[property[1]] = property[2]

    }

    return properties

}

func (p *Parser) getMatches() []string {

    regex, err := regexp.Compile(`(?m)Handle[\s\S]+?\n\n`)

    if err != nil {

        fmt.Println(err)
        return []string{}

    }

    return regex.FindAllString(p.Dmidecode, -1)

}

func getDMIType(match string) int {

    regex := regexp.MustCompile(`DMI type ([0-9]+)`)
    matches := regex.FindStringSubmatch(match)

    ret, _ := strconv.Atoi(matches[1])

    return ret

}
