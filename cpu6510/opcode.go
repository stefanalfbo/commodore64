package cpu6510

import "fmt"

type OpCodeFunc func(*CPU)

var lookupOpCode = map[byte]OpCodeFunc{
	0x00: BRK,
	0x18: CLC,
	0x38: SEC,
	0x58: CLI,
	0xB8: CLV,
	0xD8: CLD,
	0xF8: SED,
}

func OpCodeAsHex(name string) byte {
	opCode, ok := opCodes[name]
	if !ok {
		panic(fmt.Sprintf("Unknown op code, %s", name))
	}

	return opCode
}

var opCodes = map[string]byte{
	"BRK": 0x00,
	"CLC": 0x18,
	"SEC": 0x38,
	"CLI": 0x58,
	"CLV": 0xB8,
	"CLD": 0xD8,
	"SED": 0xF8,
}

// BRK - BReaKpoint. BRK is intended for use as a debugging tool which
// a programmer may place at specific points in a program, to check the state
// of processor flags at these points in the code.
func BRK(c *CPU) {
	c.statusRegister.breakCommandFlag = true
	c.statusRegister.interruptDisableFlag = true

	// BRK increments the program counter by 2 instead of 1
	c.programCounter += 2
}

// CLC - CLear Carry
func CLC(c *CPU) {
	c.statusRegister.carry = false
	c.programCounter++
}

// SEC - SEt Carry
func SEC(c *CPU) {
	c.statusRegister.carry = true
	c.programCounter++
}

// CLI - CLear Interrupt disable flag
func CLI(c *CPU) {
	c.statusRegister.interruptDisableFlag = false
	c.programCounter++
}

// CLV - CLear oVerflow flag
func CLV(c *CPU) {
	c.statusRegister.overflowFlag = false
	c.programCounter++
}

// CLD - CLear Decimal flag
func CLD(c *CPU) {
	c.statusRegister.decimalModeFlag = false
	c.programCounter++
}

// SED - SEt Decimal flag
func SED(c *CPU) {
	c.statusRegister.decimalModeFlag = true
	c.programCounter++
}
