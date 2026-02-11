package cpu6510

// getValueByImmediateAddressingMode - returns the value in memory at the
// current program counter.
func (c *CPU) getValueByImmediateAddressingMode() byte {
	value := c.ram[c.programCounter]
	c.programCounter++
	return value
}

// getValueByAbsoluteAddressingMode - returns the value in memory at the
// address specified by the next two bytes in memory.
func (c *CPU) getValueByAbsoluteAddressingMode() byte {
	address := c.readAddressFromMemory()
	c.programCounter += 2

	return c.readMemory(address)
}

// addressAbsolute - returns the address specified by the next two bytes in
// memory.
func (c *CPU) addressAbsolute() uint16 {
	address := c.readAddressFromMemory()
	c.programCounter += 2

	return address
}

// getValueByAbsoluteXAddressingMode - returns the value in memory at the
// address specified by the next two bytes in memory plus the value of the
// X register.
func (c *CPU) getValueByAbsoluteXAddressingMode() byte {
	address := c.addressAbsoluteX()

	return c.readMemory(address)
}

// addressAbsoluteX - returns the address specified by the next two bytes in
// memory plus the value of the X register.
func (c *CPU) addressAbsoluteX() uint16 {
	address := c.readAddressFromMemory() + uint16(c.xRegister)
	c.programCounter += 2

	return address
}

// getValueByAbsoluteYAddressingMode - returns the value in memory at the
// address specified by the next two bytes in memory plus the value of the
// Y register.
func (c *CPU) getValueByAbsoluteYAddressingMode() byte {
	address := c.readAddressFromMemory() + uint16(c.yRegister)
	c.programCounter += 2

	return c.readMemory(address)
}

// getValueByZeroPageAddressingMode - returns the value in memory at the
// address specified by the next byte in memory.
func (c *CPU) getValueByZeroPageAddressingMode() byte {
	address := c.addressZeroPage()

	return c.readMemory(address)
}

// addressZeroPage - returns the address specified by the next byte in memory.
func (c *CPU) addressZeroPage() uint16 {
	address := uint16(c.ram[c.programCounter])
	c.programCounter++

	return address
}

// getValueByZeroPageXAddressingMode - returns the value in memory at the
// address specified by the next byte in memory plus the value of the X
// register.
func (c *CPU) getValueByZeroPageXAddressingMode() byte {
	address := c.addressZeroPageX()

	return c.readMemory(address)
}

// addressZeroPageX - returns the address specified by the next byte in memory
// plus the value of the X register.
func (c *CPU) addressZeroPageX() uint16 {
	address := uint16(c.ram[c.programCounter] + c.xRegister)
	c.programCounter++

	return address
}

// addressIndexedIndirect - returns the address specified by the zero page
// address plus the X register.
func (c *CPU) addressIndexedIndirect() uint16 {
	zeroPageAddress := uint16(uint8(c.ram[c.programCounter] + c.xRegister))
	c.programCounter++

	lowByte := c.readMemory(zeroPageAddress)
	highByte := c.readMemory(zeroPageAddress + 1)

	return ConvertTwoBytesToAddress(highByte, lowByte)
}

// getValueByIndexedIndirectAddressingMode - returns the value in memory at
// the address specified by the zero page address plus the X register.
// The X register is added to the zero page address to get the low byte of
// the address, and the high byte is the next byte in memory.
func (c *CPU) getValueByIndexedIndirectAddressingMode() byte {
	address := c.addressIndexedIndirect()

	return c.readMemory(address)
}

// getValueByIndirectIndexedAddressingMode - returns the value in memory at
// the address pointed to by the zero page address plus the Y register.
// The Y register is added after the address fetched from the zero page
// address to get the final address.
func (c *CPU) getValueByIndirectIndexedAddressingMode() byte {
	zeroPageAddress := uint16(c.ram[c.programCounter])
	c.programCounter++

	lowByte := c.readMemory(zeroPageAddress)
	highByte := c.readMemory(zeroPageAddress + 1)

	address := ConvertTwoBytesToAddress(highByte, lowByte) + uint16(c.yRegister)

	return c.readMemory(address)
}
