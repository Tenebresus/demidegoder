# dmideGOder

A (very) simple dmidecode output decoder. 

## basic examples

dmidecode output (parser/test.txt):

```
# dmidecode 3.3
Getting SMBIOS data from sysfs.
SMBIOS 2.7 present.
127 structures occupying 5159 bytes.
Table at 0x000EC700.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
	Vendor: American Megatrends Inc.
	Version: 3.2a
	Release Date: 10/28/2015
	Address: 0xF0000
	Runtime Size: 64 kB
	ROM Size: 4 MB
	Characteristics:
		PCI is supported
		BIOS is upgradeable
		BIOS shadowing is allowed
		Boot from CD is supported
		Selectable boot is supported
		BIOS ROM is socketed
		EDD is supported
		5.25"/1.2 MB floppy services are supported (int 13h)
		3.5"/720 kB floppy services are supported (int 13h)
		3.5"/2.88 MB floppy services are supported (int 13h)
		Print screen service is supported (int 5h)
		8042 keyboard services are supported (int 9h)
		Serial services are supported (int 14h)
		Printer services are supported (int 17h)
		ACPI is supported
		USB legacy is supported
		BIOS boot specification is supported
		Function key-initiated network boot is supported
		Targeted content distribution is supported
		UEFI is supported
	BIOS Revision: 4.6
...

```

**Parsing the output**

`parser.Parse()` returns a []byte encoded json value 

```golang


file, err := os.ReadFile("parser/test.txt")
test := parser.Parse(string(file))

```

**Encoding JSON to dmidecode struct**

You can use `[]parser.DMIType` in combination with `json.Unmarshall` to convert your json encoded dmidecode output to a struct

```golang

var dmiTypes []parser.DMIType
json.Unmarshal(test, &dmiTypes)

```
