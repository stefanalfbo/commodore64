package cpu6510

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

// ASLAccumulator - Arithmetic Shift Left. ASL shifts all bits in the
// accumulator.
func ASLAccumulator(c *CPU) {
	c.programCounter++

	c.statusRegister.carryFlag = setCarryFlag(c.accumulator)

	c.accumulator <<= 1

	raiseStatusRegisterFlags(c, c.accumulator)

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

// LSR - Logical Shift Right. LSR shifts all bits in the memory location
// specified by the address.
func lsr(c *CPU, getAddress func() uint16) {
	c.programCounter++

	address := getAddress()

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value >>= 1

	raiseStatusRegisterFlags(c, value)
	// The negative status flag is always unconditionally cleared.
	c.statusRegister.negativeFlag = false

	c.writeMemory(address, value)
}

// LSRAccumulator - Logical Shift Right. LSR shifts all bits in the accumulator
// register.
func LSRAccumulator(c *CPU) {
	c.programCounter++

	c.statusRegister.carryFlag = setCarryFlag(c.accumulator)

	c.accumulator >>= 1

	raiseStatusRegisterFlags(c, c.accumulator)
	// The negative status flag is always unconditionally cleared.
	c.statusRegister.negativeFlag = false
}

// LSRZeroPage - Logical Shift Right. LSR shifts all bits in the memory
// location specified by the single byte address.
func LSRZeroPage(c *CPU) {
	lsr(c, c.addressZeroPage)
}

// LSRZeroPageX - Logical Shift Right. LSR shifts all bits in the memory
// location specified by the single byte address plus the X index register.
func LSRZeroPageX(c *CPU) {
	lsr(c, c.addressZeroPageX)
}

// LSRAbsolute - Logical Shift Right. LSR shifts all bits in the memory
// location specified by the two byte address.
func LSRAbsolute(c *CPU) {
	lsr(c, c.addressAbsolute)
}

// LSRAbsoluteX - Logical Shift Right. LSR shifts all bits in the memory
// location specified by the two byte address plus the X index register.
func LSRAbsoluteX(c *CPU) {
	lsr(c, c.addressAbsoluteX)
}

func rol(c *CPU, getAddress func() uint16) {
	c.programCounter++

	address := getAddress()

	value := c.readMemory(address)

	carry := c.statusRegister.carryFlag

	c.statusRegister.carryFlag = value&0x80 == 0x80

	value <<= 1

	if carry {
		value |= 0x01
	}

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)
}

// ROLAccumulator - Rotate Left. ROL shifts all bits in the accumulator register.
func ROLAccumulator(c *CPU) {
	c.programCounter++

	carry := c.statusRegister.carryFlag

	c.statusRegister.carryFlag = c.accumulator&0x80 == 0x80

	c.accumulator <<= 1

	if carry {
		c.accumulator |= 0x01
	}

	raiseStatusRegisterFlags(c, c.accumulator)
}

// ROLZeroPage - Rotate Left. ROL shifts all bits in the memory location
// specified by the single byte address.
func ROLZeroPage(c *CPU) {
	rol(c, c.addressZeroPage)
}

// ROLZeroPageX - Rotate Left. ROL shifts all bits in the memory location
// specified by the single byte address plus the X index register.
func ROLZeroPageX(c *CPU) {
	rol(c, c.addressZeroPageX)
}

// ROLAbsolute - Rotate Left. ROL shifts all bits in the memory location
// specified by the two byte address.
func ROLAbsolute(c *CPU) {
	rol(c, c.addressAbsolute)
}

// ROLAbsoluteX - Rotate Left. ROL shifts all bits in the memory location
// specified by the two byte address plus the X index register.
func ROLAbsoluteX(c *CPU) {
	rol(c, c.addressAbsoluteX)
}

func ror(c *CPU, getAddress func() uint16) {
	c.programCounter++

	address := getAddress()

	value := c.readMemory(address)

	carry := c.statusRegister.carryFlag

	c.statusRegister.carryFlag = value&0x01 == 0x01

	value >>= 1

	if carry {
		value |= 0x80
	}

	raiseStatusRegisterFlags(c, value)

	c.writeMemory(address, value)
}

// ROR - Rotate Right. ROR shifts all bits in the accumulator register.
func RORAccumulator(c *CPU) {
	c.programCounter++

	carry := c.statusRegister.carryFlag

	c.statusRegister.carryFlag = c.accumulator&0x01 == 0x01

	c.accumulator >>= 1

	if carry {
		c.accumulator |= 0x80
	}

	raiseStatusRegisterFlags(c, c.accumulator)
}

// RORZeroPage - Rotate Right. ROR shifts all bits in the memory location
// specified by the single byte address.
func RORZeroPage(c *CPU) {
	ror(c, c.addressZeroPage)
}

// RORZeroPageX - Rotate Right. ROR shifts all bits in the memory location
// specified by the single byte address plus the X index register.
func RORZeroPageX(c *CPU) {
	ror(c, c.addressZeroPageX)
}

// RORAbsolute - Rotate Right. ROR shifts all bits in the memory location
// specified by the two byte address.
func RORAbsolute(c *CPU) {
	ror(c, c.addressAbsolute)
}

// RORAbsoluteX - Rotate Right. ROR shifts all bits in the memory location
// specified by the two byte address plus the X index register.
func RORAbsoluteX(c *CPU) {
	ror(c, c.addressAbsoluteX)
}
