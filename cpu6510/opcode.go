package cpu6510

import "fmt"

type OpCodeFunc func(*CPU)

var lookupOpCode = map[byte]OpCodeFunc{
	0x00: BRK,
	0x08: PHP,
	0x18: CLC,
	0x38: SEC,
	0x48: PHA,
	0x58: CLI,
	0x68: PLA,
	0x78: SEI,
	0x88: DEY,
	0x8A: TXA,
	0x98: TYA,
	0xA8: TAY,
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
	"PHP": 0x08,
	"CLC": 0x18,
	"SEC": 0x38,
	"PHA": 0x48,
	"CLI": 0x58,
	"PLA": 0x68,
	"SEI": 0x78,
	"DEY": 0x88,
	"TXA": 0x8A,
	"TYA": 0x98,
	"TAY": 0xA8,
	"TAX": 0xAA,
	"CLV": 0xB8,
	"INY": 0xC8,
	"DEX": 0xCA,
	"CLD": 0xD8,
	"NOP": 0xEA,
	"INX": 0xE8,
	"SED": 0xF8,
}

func raiseStatusRegisterFlags(c *CPU, value byte) {
	if value == 0 {
		c.statusRegister.zeroFlag = true
	}

	if value&0x80 == 0x80 {
		c.statusRegister.negativeFlag = true
	}
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

// PHP - PusH Processor status flags. Pushes the current value of the
// processor status register onto the stack.
func PHP(c *CPU) {
	c.pushOnStack(c.statusRegister.asByte())
	c.programCounter++
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

// PHA - PusH Accumulator. pushes the current value in the accumulator onto
// the stack.
func PHA(c *CPU) {
	c.pushOnStack(c.accumulator)
	c.programCounter++
}

// CLI - CLear Interrupt disable flag
func CLI(c *CPU) {
	c.statusRegister.interruptDisableFlag = false
	c.programCounter++
}

// PLA - PuLl Accumulator. pulls the current value from the stack and places
// it in the accumulator.
func PLA(c *CPU) {
	c.accumulator = c.popFromStack()

	raiseStatusRegisterFlags(c, c.accumulator)

	c.programCounter++
}

// SEI - SEt Interrupt disable flag, preventing the CPU from responding to
// IRQ interrupts.
func SEI(c *CPU) {
	c.statusRegister.interruptDisableFlag = true
	c.programCounter++
}

// DEY - DEcrement Y register. decreases the numerical value held in the Y
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func DEY(c *CPU) {
	y := uint8(c.yRegister)
	y--
	c.yRegister = y

	raiseStatusRegisterFlags(c, c.yRegister)

	c.programCounter++
}

// TXA - Transfer X to A. copies the current value in the X index register to
// the accumulator.
func TXA(c *CPU) {
	c.accumulator = c.xRegister

	raiseStatusRegisterFlags(c, c.accumulator)

	c.programCounter++
}

// TYA - Transfer Y to A. copies the current value in the Y index register to
// the accumulator.
func TYA(c *CPU) {
	c.accumulator = c.yRegister

	raiseStatusRegisterFlags(c, c.accumulator)

	c.programCounter++
}

// TAY - Transfer A to Y. copies the current value in the accumulator to the
// Y index register.
func TAY(c *CPU) {
	c.yRegister = c.accumulator

	raiseStatusRegisterFlags(c, c.yRegister)

	c.programCounter++
}

// TAX - Transfer A to X. copies the current value in the accumulator to the
// X index register.
func TAX(c *CPU) {
	c.xRegister = c.accumulator

	raiseStatusRegisterFlags(c, c.xRegister)

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

	raiseStatusRegisterFlags(c, c.yRegister)

	c.programCounter++
}

// DEX - DEcrement X register. decreases the numerical value held in the X
// index register by one, and "wraps over" when the numerical limits of a
// byte are exceeded.
func DEX(c *CPU) {
	x := uint8(c.xRegister)
	x--
	c.xRegister = x

	raiseStatusRegisterFlags(c, c.xRegister)

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

	raiseStatusRegisterFlags(c, c.xRegister)

	c.programCounter++
}

// SED - SEt Decimal flag
func SED(c *CPU) {
	c.statusRegister.decimalModeFlag = true
	c.programCounter++
}
