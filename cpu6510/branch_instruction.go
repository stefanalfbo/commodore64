package cpu6510

// BPL - Branch if PLus. BPL branches to the given address if the negative
// flag in the status register is not set.
func BPL(c *CPU) {
	c.programCounter++

	if !c.statusRegister.negativeFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BMI - Branch if Minus. BMI branches to the given address if the negative
// flag in the status register is set.
func BMI(c *CPU) {
	c.programCounter++

	if c.statusRegister.negativeFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BVC - Branch if oVerflow Clear. BVC branches to the given address if the
// overflow flag in the status register is not set.
func BVC(c *CPU) {
	c.programCounter++

	if !c.statusRegister.overflowFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BVS - Branch if oVerflow Set. BVS branches to the given address if the
// overflow flag in the status register is set.
func BVS(c *CPU) {
	c.programCounter++

	if c.statusRegister.overflowFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BCC - Branch if Carry Clear. BCC branches to the given address if the
// carry flag in the status register is not set.
func BCC(c *CPU) {
	c.programCounter++

	if !c.statusRegister.carryFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BCS - Branch if Carry Set. BCS branches to the given address if the
// carry flag in the status register is set.
func BCS(c *CPU) {
	c.programCounter++

	if c.statusRegister.carryFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BNE - Branch if Not Equal. BNE branches to the given address if the
// zero flag in the status register is not set.
func BNE(c *CPU) {
	c.programCounter++

	if !c.statusRegister.zeroFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}

// BEQ - Branch if Equal. BEQ branches to the given address if the zero
// flag in the status register is set.
func BEQ(c *CPU) {
	c.programCounter++

	if c.statusRegister.zeroFlag {
		c.programCounter = uint16(int16(c.programCounter) + int16(c.ram[c.programCounter]))
	}
}
