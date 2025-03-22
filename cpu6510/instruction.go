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
	0x21: ANDIndexedIndirect,
	0x25: ANDZeroPage,
	0x28: PLP,
	0x29: ANDImmediate,
	0x2D: ANDAbsolute,
	0x31: ANDIndirectIndexed,
	0x35: ANDZeroPageX,
	0x38: SEC,
	0x39: ANDAbsoluteY,
	0x3D: ANDAbsoluteX,
	0x41: EORIndexedIndirect,
	0x45: EORZeroPage,
	0x48: PHA,
	0x49: EORImmediate,
	0x4D: EORAbsolute,
	0x51: EORIndirectIndexed,
	0x55: EORZeroPageX,
	0x58: CLI,
	0x59: EORAbsoluteY,
	0x5D: EORAbsoluteX,
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
	"ANDIndexedIndirect": 0x21,
	"ANDZeroPage":        0x25,
	"PLP":                0x28,
	"ANDImmediate":       0x29,
	"ANDAbsolute":        0x2D,
	"ANDIndirectIndexed": 0x31,
	"ANDZeroPageX":       0x35,
	"SEC":                0x38,
	"ANDAbsoluteY":       0x39,
	"ANDAbsoluteX":       0x3D,
	"EORIndexedIndirect": 0x41,
	"EORZeroPage":        0x45,
	"PHA":                0x48,
	"EORImmediate":       0x49,
	"EORAbsolute":        0x4D,
	"EORIndirectIndexed": 0x51,
	"EORZeroPageX":       0x55,
	"CLI":                0x58,
	"EORAbsoluteY":       0x59,
	"EORAbsoluteX":       0x5D,
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
func ora(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	c.accumulator |= value

	raiseStatusRegisterFlags(c, c.accumulator)
}

// ORAImmediate - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAImmediate(c *CPU) {
	ora(c, c.getValueByImmediateAddressingMode)
}

// ORAAbsolute - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsolute(c *CPU) {
	ora(c, c.getValueByAbsoluteAddressingMode)
}

// ORAAbsoluteX - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsoluteX(c *CPU) {
	ora(c, c.getValueByAbsoluteXAddressingMode)
}

// ORAAbsoluteY - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAAbsoluteY(c *CPU) {
	ora(c, c.getValueByAbsoluteYAddressingMode)
}

// ORAZeroPage - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAZeroPage(c *CPU) {
	ora(c, c.getValueByZeroPageAddressingMode)
}

// ORAZeroPageX - OR with Accumulator. ORA performs a logical OR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ORAZeroPageX(c *CPU) {
	ora(c, c.getValueByZeroPageXAddressingMode)
}

// ORAIndexedIndirect - OR with Accumulator. ORA performs a logical OR
// between the value in the accumulator and the value in memory, and stores
// the result in the accumulator.
func ORAIndexedIndirect(c *CPU) {
	ora(c, c.getValueByIndexedIndirectAddressingMode)
}

// ORAIndirectIndexed - OR with Accumulator. ORA performs a logical OR
// between the value in the accumulator and the value in memory, and stores
// the result in the accumulator.
func ORAIndirectIndexed(c *CPU) {
	ora(c, c.getValueByIndirectIndexedAddressingMode)
}

// AND - Logical AND. AND performs a logical AND between the value in the
// accumulator and the given value, and stores the result in the accumulator.
func and(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	c.accumulator &= value

	raiseStatusRegisterFlags(c, c.accumulator)
}

// ANDImmediate - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDImmediate(c *CPU) {
	and(c, c.getValueByImmediateAddressingMode)
}

// ANDAbsolute - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDAbsolute(c *CPU) {
	and(c, c.getValueByAbsoluteAddressingMode)
}

// ANDAbsoluteX - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDAbsoluteX(c *CPU) {
	and(c, c.getValueByAbsoluteXAddressingMode)
}

// ANDAbsoluteY - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDAbsoluteY(c *CPU) {
	and(c, c.getValueByAbsoluteYAddressingMode)
}

// ANDZeroPage - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDZeroPage(c *CPU) {
	and(c, c.getValueByZeroPageAddressingMode)
}

// ANDZeroPageX - Logical AND. AND performs a logical AND between the value in
// the accumulator and the value in memory, and stores the result in the
// accumulator.
func ANDZeroPageX(c *CPU) {
	and(c, c.getValueByZeroPageXAddressingMode)
}

// ANDIndexedIndirect - Logical AND. AND performs a logical AND between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ANDIndexedIndirect(c *CPU) {
	and(c, c.getValueByIndexedIndirectAddressingMode)
}

// ANDIndirectIndexed - Logical AND. AND performs a logical AND between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func ANDIndirectIndexed(c *CPU) {
	and(c, c.getValueByIndirectIndexedAddressingMode)
}

func eor(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	c.accumulator ^= value

	raiseStatusRegisterFlags(c, c.accumulator)
}

// EORImmediate - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORImmediate(c *CPU) {
	eor(c, c.getValueByImmediateAddressingMode)
}

// EORAbsolute - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORAbsolute(c *CPU) {
	eor(c, c.getValueByAbsoluteAddressingMode)
}

// EORAbsoluteX - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORAbsoluteX(c *CPU) {
	eor(c, c.getValueByAbsoluteXAddressingMode)
}

// EORAbsoluteY - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORAbsoluteY(c *CPU) {
	eor(c, c.getValueByAbsoluteYAddressingMode)
}

// EORZeroPage - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORZeroPage(c *CPU) {
	eor(c, c.getValueByZeroPageAddressingMode)
}

// EORZeroPageX - Exclusive OR. EOR performs a logical XOR between the value
// in the accumulator and the value in memory, and stores the result in the
// accumulator.
func EORZeroPageX(c *CPU) {
	eor(c, c.getValueByZeroPageXAddressingMode)
}

// EORIndexedIndirect - Exclusive OR. EOR performs a logical XOR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func EORIndexedIndirect(c *CPU) {
	eor(c, c.getValueByIndexedIndirectAddressingMode)
}

// EORIndirectIndexed - Exclusive OR. EOR performs a logical XOR between the
// value in the accumulator and the value in memory, and stores the result in
// the accumulator.
func EORIndirectIndexed(c *CPU) {
	eor(c, c.getValueByIndirectIndexedAddressingMode)
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

func cmp(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	tmp := c.accumulator - value

	raiseStatusRegisterFlags(c, tmp)

	c.statusRegister.carryFlag = c.accumulator >= value
}

// CMPImmediate - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPImmediate(c *CPU) {
	cmp(c, c.getValueByImmediateAddressingMode)
}

// CMPAbsolute - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsolute(c *CPU) {
	cmp(c, c.getValueByAbsoluteAddressingMode)
}

// CMPZeroPageX - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPZeroPageX(c *CPU) {
	cmp(c, c.getValueByZeroPageXAddressingMode)
}

// CMPAbsoluteY - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsoluteY(c *CPU) {
	cmp(c, c.getValueByAbsoluteYAddressingMode)
}

// CMPAbsoluteX - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPAbsoluteX(c *CPU) {
	cmp(c, c.getValueByAbsoluteXAddressingMode)
}

// CMPIndexedIndirect - CoMPare. CMP compares the value in the accumulator with
// the value in memory, and sets the zero and negative flags in the status
// register based on the result.
func CMPIndexedIndirect(c *CPU) {
	cmp(c, c.getValueByIndexedIndirectAddressingMode)
}

// CMPIndirectIndexed - CoMPare. CMP compares the value in the accumulator with
// the value in memory, and sets the zero and negative flags in the status
// register based on the result.
func CMPIndirectIndexed(c *CPU) {
	cmp(c, c.getValueByIndirectIndexedAddressingMode)

	c.programCounter++
}

// CMPZeroPage - CoMPare. CMP compares the value in the accumulator with the
// value in memory, and sets the zero and negative flags in the status register
// based on the result.
func CMPZeroPage(c *CPU) {
	cmp(c, c.getValueByZeroPageAddressingMode)
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
