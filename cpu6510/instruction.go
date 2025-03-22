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
	0x28: PLP,
	0x38: SEC,
	0x48: PHA,
	0x58: CLI,
	0x60: RTS,
	0x68: PLA,
	0x78: SEI,
	0x88: DEY,
	0x8A: TXA,
	0x98: TYA,
	0x9A: TXS,
	0xA8: TAY,
	0xAA: TAX,
	0xBA: TSX,
	0xB8: CLV,
	0xC1: CMPIndexedIndirect,
	0xC5: CMPZeroPage,
	0xC8: INY,
	0xC9: CMPImmediate,
	0xCA: DEX,
	0xCD: CMPAbsolute,
	0xD1: CMPIndirectIndexed,
	0xD5: CMPZeroPageX,
	0xD8: CLD,
	0xD9: CMPAbsoluteY,
	0xDD: CMPAbsoluteX,
	0xEA: NOP,
	0xE8: INX,
	0xF8: SED,
}

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
	"PLP":                0x28,
	"SEC":                0x38,
	"PHA":                0x48,
	"CLI":                0x58,
	"RTS":                0x60,
	"PLA":                0x68,
	"SEI":                0x78,
	"DEY":                0x88,
	"TXA":                0x8A,
	"TYA":                0x98,
	"TXS":                0x9A,
	"TAY":                0xA8,
	"TAX":                0xAA,
	"TSX":                0xBA,
	"CLV":                0xB8,
	"CMPIndexedIndirect": 0xC1,
	"CMPZeroPage":        0xC5,
	"INY":                0xC8,
	"CMPImmediate":       0xC9,
	"DEX":                0xCA,
	"CMPAbsolute":        0xCD,
	"CMPIndirectIndexed": 0xD1,
	"CMPZeroPageX":       0xD5,
	"CLD":                0xD8,
	"CMPAbsoluteY":       0xD9,
	"CMPAbsoluteX":       0xDD,
	"NOP":                0xEA,
	"INX":                0xE8,
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

// ORA - OR with Accumulator. ORA performs a logical OR between the value in
// the accumulator and the given value, and stores the result in the
// accumulator.
func ora(c *CPU, value byte) {
	c.accumulator |= value

	raiseStatusRegisterFlags(c, c.accumulator)
}

// ORAImmediate - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAImmediate(c *CPU) {
	c.programCounter++

	value := c.getValueByImmediateAddressingMode()

	ora(c, value)
}

// ORAAbsolute - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsolute(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteAddressingMode()

	ora(c, value)
}

// ORAAbsoluteX - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsoluteX(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteXAddressingMode()

	ora(c, value)
}

// ORAAbsoluteY - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsoluteY(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteYAddressingMode()

	ora(c, value)
}

// ORAZeroPage - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAZeroPage(c *CPU) {
	c.programCounter++

	value := c.getValueByZeroPageAddressingMode()

	ora(c, value)
}

// ORAZeroPageX - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAZeroPageX(c *CPU) {
	c.programCounter++

	value := c.getValueByZeroPageXAddressingMode()

	ora(c, value)
}

// ORAIndexedIndirect - OR with Accumulator. ORA performs a logical OR
// between the value in the accumulator and the value in memory, and stores
// the result in the accumulator.
func ORAIndexedIndirect(c *CPU) {
	c.programCounter++

	value := c.getValueByIndexedIndirectAddressingMode()

	ora(c, value)
}

// ORAIndirectIndexed - OR with Accumulator. ORA performs a logical OR
// between the value in the accumulator and the value in memory, and stores
// the result in the accumulator.
func ORAIndirectIndexed(c *CPU) {
	c.programCounter++

	value := c.getValueByIndirectIndexedAddressingMode()

	ora(c, value)
}

// ASLZeroPage - Arithmetic Shift Left. ASL shifts all bits in the memory
// location specified by the single byte address.
func ASLZeroPage(c *CPU) {
	c.programCounter++

	address := uint16(c.ram[c.programCounter])

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value <<= 1

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)

	c.programCounter++
}

// ASLAccumulator - Arithmetic Shift Left. ASL shifts all bits in the
// accumulator.
func ASLAccumulator(c *CPU) {
	c.statusRegister.carryFlag = c.accumulator&0x80 == 0x80

	c.accumulator <<= 1

	raiseStatusRegisterFlags(c, c.accumulator)

	c.programCounter++
}

// ASLAbsolute - Arithmetic Shift Left. ASL shifts all bits in the memory
// location specified by the two byte address.
func ASLAbsolute(c *CPU) {
	c.programCounter++

	address := c.readAddressFromMemory()

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value <<= 1

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)

	c.programCounter += 2
}

// PHP - PusH Processor status flags. Pushes the current value of the
// processor status register onto the stack.
func PHP(c *CPU) {
	c.pushOnStack(c.statusRegister.asByte())
	c.programCounter++
}

// ASLZeroPageX - Arithmetic Shift Left. ASL shifts all bits in the memory
// location specified by the single byte address plus the X index register.
func ASLZeroPageX(c *CPU) {
	c.programCounter++

	address := uint16(c.ram[c.programCounter] + c.xRegister)

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value <<= 1

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)

	c.programCounter++
}

// CLC - CLear Carry
func CLC(c *CPU) {
	c.statusRegister.carryFlag = false
	c.programCounter++
}

// ASLAbsoluteX - Arithmetic Shift Left. ASL shifts all bits in the memory
// location specified by the two byte address plus the X index register.
func ASLAbsoluteX(c *CPU) {
	c.programCounter++

	// Indexed absolute addressing is an addressing mode in which the
	// contents of the X register is added to a given base address, to
	// obtain the "target" address.
	address := c.readAddressFromMemory() + uint16(c.xRegister)

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value <<= 1

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)

	c.programCounter += 2
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

func cmp(c *CPU, value byte) {
	tmp := c.accumulator - value

	raiseStatusRegisterFlags(c, tmp)

	c.statusRegister.carryFlag = c.accumulator >= value
}

// CMPImmediate - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPImmediate(c *CPU) {
	c.programCounter++

	value := c.getValueByImmediateAddressingMode()

	cmp(c, value)
}

// CMPAbsolute - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsolute(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteAddressingMode()

	cmp(c, value)
}

// CMPZeroPageX - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPZeroPageX(c *CPU) {
	c.programCounter++

	value := c.getValueByZeroPageXAddressingMode()

	cmp(c, value)
}

// CMPAbsoluteY - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsoluteY(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteYAddressingMode()

	cmp(c, value)
}

// CMPAbsoluteX - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsoluteX(c *CPU) {
	c.programCounter++

	value := c.getValueByAbsoluteXAddressingMode()

	cmp(c, value)
}

// CMPIndexedIndirect - CoMPare. CMP compares the value in the accumulator with
// the value in memory, and sets the zero and negative flags in the status
// register based on the result.
func CMPIndexedIndirect(c *CPU) {
	c.programCounter++

	value := c.getValueByIndexedIndirectAddressingMode()

	cmp(c, value)
}

// CMPIndirectIndexed - CoMPare. CMP compares the value in the accumulator with
// the value in memory, and sets the zero and negative flags in the status
// register based on the result.
func CMPIndirectIndexed(c *CPU) {
	c.programCounter++

	value := c.getValueByIndirectIndexedAddressingMode()

	cmp(c, value)

	c.programCounter++
}

// CMPZeroPage - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPZeroPage(c *CPU) {
	c.programCounter++

	value := c.getValueByZeroPageAddressingMode()

	cmp(c, value)
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
