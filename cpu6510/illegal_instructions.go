package cpu6510

// JAM - JaM/KIL. Illegal opcode that halts the CPU.
func JAM(c *CPU) {
	c.isJammed = true
}

// SLOIndexedIndirect - Shift Left then ORA with accumulator.
func SLOIndexedIndirect(c *CPU) {
	c.programCounter++

	address := c.addressIndexedIndirect()

	value := c.readMemory(address)

	c.statusRegister.carryFlag = setCarryFlag(value)

	value <<= 1

	c.writeMemory(address, value)

	c.accumulator |= value

	raiseStatusRegisterFlags(c, c.accumulator)
}
