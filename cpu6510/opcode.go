package cpu6510

import "fmt"

type OpCodeFunc func(*CPU)

var lookupOpCode = map[byte]OpCodeFunc{
	0x00: BRK,
	0x18: CLC,
	0x38: SEC,
	0x58: CLI,
	0x88: DEY,
	0xAA: TAX,
	0xB8: CLV,
	0xC8: INY,
	0xCA: DEX,
	0xD8: CLD,
	0xEA: NOP,
	0xE8: INX,
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
	"DEY": 0x88,
	"TAX": 0xAA,
	"CLV": 0xB8,
	"INY": 0xC8,
	"DEX": 0xCA,
	"CLD": 0xD8,
	"NOP": 0xEA,
	"INX": 0xE8,
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
	c.statusRegister.carryFlag = false
	c.programCounter++
}

// SEC - SEt Carry
func SEC(c *CPU) {
	c.statusRegister.carryFlag = true
	c.programCounter++
}

// CLI - CLear Interrupt disable flag
func CLI(c *CPU) {
	c.statusRegister.interruptDisableFlag = false
	c.programCounter++
}

// DEY - DEcrement Y register. decreases the numerical value held in the Y
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func DEY(c *CPU) {
	y := uint8(c.yRegister)
	y--
	c.yRegister = y

	if c.yRegister == 0 {
		c.statusRegister.zeroFlag = true
	}

	if c.yRegister&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}

	c.programCounter++
}

// TAX - Transfer A to X. copies the current value in the accumulator to the
// X index register.
func TAX(c *CPU) {
	c.xRegister = c.accumulator

	if c.xRegister == 0 {
		c.statusRegister.zeroFlag = true
	}

	if c.xRegister&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}

	c.programCounter++
}

// CLV - CLear oVerflow flag
func CLV(c *CPU) {
	c.statusRegister.overflowFlag = false
	c.programCounter++
}

// INY - INcrement Y register. increases the numerical value held in the Y
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func INY(c *CPU) {
	y := uint8(c.yRegister)
	y++
	c.yRegister = y

	if c.yRegister == 0 {
		c.statusRegister.zeroFlag = true
	}

	if c.yRegister&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}

	c.programCounter++
}

// DEX - DEcrement X register. decreases the numerical value held in the X
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func DEX(c *CPU) {
	x := uint8(c.xRegister)
	x--
	c.xRegister = x

	if c.xRegister == 0 {
		c.statusRegister.zeroFlag = true
	}

	if c.xRegister&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}

	c.programCounter++
}

// CLD - CLear Decimal flag
func CLD(c *CPU) {
	c.statusRegister.decimalModeFlag = false
	c.programCounter++
}

// NOP - No OPeration.
func NOP(c *CPU) {
	c.programCounter++
}

// INX - INcrement X register. increases the numerical value held in the X
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func INX(c *CPU) {
	x := uint8(c.xRegister)
	x++
	c.xRegister = x

	if c.xRegister == 0 {
		c.statusRegister.zeroFlag = true
	}

	if c.xRegister&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}

	c.programCounter++
}

// SED - SEt Decimal flag
func SED(c *CPU) {
	c.statusRegister.decimalModeFlag = true
	c.programCounter++
}
