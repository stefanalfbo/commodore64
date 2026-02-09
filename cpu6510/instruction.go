package cpu6510

import "fmt"

type InstructionFunc func(*CPU)

var lookupInstruction = map[byte]InstructionFunc{
	0x00: BRK,
	0x01: ORAIndexedIndirect,
	0x05: ORAZeroPage,
	0x06: ASLZeroPage,
	0x08: PHP,
	0x09: ORAImmediate,
	0x10: BPL,
	0x0A: ASLAccumulator,
	0x0D: ORAAbsolute,
	0x0E: ASLAbsolute,
	0x11: ORAIndirectIndexed,
	0x15: ORAZeroPageX,
	0x16: ASLZeroPageX,
	0x18: CLC,
	0x19: ORAAbsoluteY,
	0x1D: ORAAbsoluteX,
	0x1E: ASLAbsoluteX,
	0x21: ANDIndexedIndirect,
	0x24: BITZeroPage,
	0x25: ANDZeroPage,
	0x26: ROLZeroPage,
	0x28: PLP,
	0x29: ANDImmediate,
	0x2A: ROLAccumulator,
	0x2C: BITAbsolute,
	0x2D: ANDAbsolute,
	0x2E: ROLAbsolute,
	0x30: BMI,
	0x31: ANDIndirectIndexed,
	0x35: ANDZeroPageX,
	0x36: ROLZeroPageX,
	0x38: SEC,
	0x39: ANDAbsoluteY,
	0x3E: ROLAbsoluteX,
	0x3D: ANDAbsoluteX,
	0x41: EORIndexedIndirect,
	0x45: EORZeroPage,
	0x46: LSRZeroPage,
	0x48: PHA,
	0x49: EORImmediate,
	0x4A: LSRAccumulator,
	0x4D: EORAbsolute,
	0x4E: LSRAbsolute,
	0x50: BVC,
	0x51: EORIndirectIndexed,
	0x55: EORZeroPageX,
	0x56: LSRZeroPageX,
	0x58: CLI,
	0x59: EORAbsoluteY,
	0x5D: EORAbsoluteX,
	0x5E: LSRAbsoluteX,
	0x60: RTS,
	0x66: RORZeroPage,
	0x68: PLA,
	0x6A: RORAccumulator,
	0x6E: RORAbsolute,
	0x70: BVS,
	0x76: RORZeroPageX,
	0x78: SEI,
	0x7E: RORAbsoluteX,
	0x88: DEY,
	0x8A: TXA,
	0x90: BCC,
	0x98: TYA,
	0x9A: TXS,
	0xA8: TAY,
	0xAA: TAX,
	0xB0: BCS,
	0xBA: TSX,
	0xB8: CLV,
	0xC0: CPYImmediate,
	0xC1: CMPIndexedIndirect,
	0xC4: CPYZeroPage,
	0xC5: CMPZeroPage,
	0xC8: INY,
	0xC9: CMPImmediate,
	0xCA: DEX,
	0xCC: CPYAbsolute,
	0xCD: CMPAbsolute,
	0xD0: BNE,
	0xD1: CMPIndirectIndexed,
	0xD5: CMPZeroPageX,
	0xD8: CLD,
	0xD9: CMPAbsoluteY,
	0xDD: CMPAbsoluteX,
	0xE0: CPXImmediate,
	0xE4: CPXZeroPage,
	0xEA: NOP,
	0xEC: CPXAbsolute,
	0xE8: INX,
	0xF0: BEQ,
	0xF8: SED,
}

// TODO: Perhaps move this as a helper function
// for the test cases since it is not used anywhere else
func InstructionAsHex(name string) byte {
	instruction, ok := instructions[name]
	if !ok {
		panic(fmt.Sprintf("Unknown instruction, %s", name))
	}

	return instruction
}

var instructions = map[string]byte{
	"BRK":                0x00,
	"ORAIndexedIndirect": 0x01,
	"ORAZeroPage":        0x05,
	"ASLZeroPage":        0x06,
	"PHP":                0x08,
	"ORAImmediate":       0x09,
	"BPL":                0x10,
	"ASLAccumulator":     0x0A,
	"ORAAbsolute":        0x0D,
	"ASLAbsolute":        0x0E,
	"ORAIndirectIndexed": 0x11,
	"ORAZeroPageX":       0x15,
	"ASLZeroPageX":       0x16,
	"CLC":                0x18,
	"ORAAbsoluteY":       0x19,
	"ORAAbsoluteX":       0x1D,
	"ASLAbsoluteX":       0x1E,
	"ANDIndexedIndirect": 0x21,
	"BITZeroPage":        0x24,
	"ANDZeroPage":        0x25,
	"ROLZeroPage":        0x26,
	"PLP":                0x28,
	"ANDImmediate":       0x29,
	"ROLAccumulator":     0x2A,
	"BITAbsolute":        0x2C,
	"ROLAbsolute":        0x2E,
	"ANDAbsolute":        0x2D,
	"BMI":                0x30,
	"ANDIndirectIndexed": 0x31,
	"ANDZeroPageX":       0x35,
	"ROLZeroPageX":       0x36,
	"SEC":                0x38,
	"ANDAbsoluteY":       0x39,
	"ANDAbsoluteX":       0x3D,
	"ROLAbsoluteX":       0x3E,
	"EORIndexedIndirect": 0x41,
	"EORZeroPage":        0x45,
	"LSRZeroPage":        0x46,
	"PHA":                0x48,
	"EORImmediate":       0x49,
	"LSRAccumulator":     0x4A,
	"EORAbsolute":        0x4D,
	"LSRAbsolute":        0x4E,
	"BVC":                0x50,
	"EORIndirectIndexed": 0x51,
	"EORZeroPageX":       0x55,
	"LSRZeroPageX":       0x56,
	"CLI":                0x58,
	"EORAbsoluteY":       0x59,
	"EORAbsoluteX":       0x5D,
	"LSRAbsoluteX":       0x5E,
	"RTS":                0x60,
	"RORZeroPage":        0x66,
	"PLA":                0x68,
	"RORAccumulator":     0x6A,
	"RORAbsolute":        0x6E,
	"BVS":                0x70,
	"RORZeroPageX":       0x76,
	"SEI":                0x78,
	"RORAbsoluteX":       0x7E,
	"DEY":                0x88,
	"TXA":                0x8A,
	"BCC":                0x90,
	"TYA":                0x98,
	"TXS":                0x9A,
	"TAY":                0xA8,
	"TAX":                0xAA,
	"BCS":                0xB0,
	"TSX":                0xBA,
	"CLV":                0xB8,
	"CPYImmediate":       0xC0,
	"CMPIndexedIndirect": 0xC1,
	"CPYZeroPage":        0xC4,
	"CMPZeroPage":        0xC5,
	"INY":                0xC8,
	"CMPImmediate":       0xC9,
	"DEX":                0xCA,
	"CPYAbsolute":        0xCC,
	"CMPAbsolute":        0xCD,
	"BNE":                0xD0,
	"CMPIndirectIndexed": 0xD1,
	"CMPZeroPageX":       0xD5,
	"CLD":                0xD8,
	"CMPAbsoluteY":       0xD9,
	"CMPAbsoluteX":       0xDD,
	"CPXImmediate":       0xE0,
	"CPXZeroPage":        0xE4,
	"NOP":                0xEA,
	"CPXAbsolute":        0xEC,
	"INX":                0xE8,
	"BEQ":                0xF0,
	"SED":                0xF8,
}

// ConvertTwoBytesToAddress - converts two bytes into a single address.
func ConvertTwoBytesToAddress(highByte, lowByte byte) uint16 {
	return (uint16(highByte) << 8) | uint16(lowByte)
}

// setCarryFlag - returns true if the originally value in bit 7
// was set, otherwise false.
func setCarryFlag(value byte) bool {
	return value&0x80 == 0x80
}

// raiseStatusRegisterFlags - sets the zero and negative flags in the status
// register based on the value passed in.
func raiseStatusRegisterFlags(c *CPU, value byte) {
	c.statusRegister.zeroFlag = value == 0

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

// PLP - PuLl Processor status register flags. Pulls the current value from
// the stack and places it in the processor status register.
func PLP(c *CPU) {
	value := c.popFromStack()

	c.statusRegister = newStatusRegister(value)

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

// RTS - ReTurn from Subroutine. pulls the program counter from the stack and
// places it in the program counter.
func RTS(c *CPU) {
	lowByte := c.popFromStack()
	highByte := c.popFromStack()

	programCounterAddress := ConvertTwoBytesToAddress(highByte, lowByte)

	c.programCounter = programCounterAddress
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

// TXS - Transfer X to Stack pointer. copies the current value in the X index
// register to the stack pointer.
func TXS(c *CPU) {
	c.stackPointer = c.xRegister
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

// TSX - Transfer Stack pointer to X. copies the current value in the stack
// pointer to the X index register.
func TSX(c *CPU) {
	c.xRegister = c.stackPointer

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
