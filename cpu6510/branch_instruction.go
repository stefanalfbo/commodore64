package cpu6510

func branchOnFlag(c *CPU, flag bool) {
	c.IncrementProgramCounter()

	if flag {
		c.SetProgramCounter(uint16(int16(c.ProgramCounter()) + int16(c.ReadRam(c.ProgramCounter()))))
	}
}

// BPL - Branch if PLus. BPL branches to the given address if the negative
// flag in the status register is not set.
func BPL(c *CPU) {
	branchOnFlag(c, !c.StatusRegister().negativeFlag)
}

// BMI - Branch if Minus. BMI branches to the given address if the negative
// flag in the status register is set.
func BMI(c *CPU) {
	branchOnFlag(c, c.StatusRegister().negativeFlag)
}

// BVC - Branch if oVerflow Clear. BVC branches to the given address if the
// overflow flag in the status register is not set.
func BVC(c *CPU) {
	branchOnFlag(c, !c.StatusRegister().overflowFlag)
}

// BVS - Branch if oVerflow Set. BVS branches to the given address if the
// overflow flag in the status register is set.
func BVS(c *CPU) {
	branchOnFlag(c, c.StatusRegister().overflowFlag)
}

// BCC - Branch if Carry Clear. BCC branches to the given address if the
// carry flag in the status register is not set.
func BCC(c *CPU) {
	branchOnFlag(c, !c.StatusRegister().carryFlag)
}

// BCS - Branch if Carry Set. BCS branches to the given address if the
// carry flag in the status register is set.
func BCS(c *CPU) {
	branchOnFlag(c, c.StatusRegister().carryFlag)
}

// BNE - Branch if Not Equal. BNE branches to the given address if the
// zero flag in the status register is not set.
func BNE(c *CPU) {
	branchOnFlag(c, !c.StatusRegister().zeroFlag)
}

// BEQ - Branch if Equal. BEQ branches to the given address if the zero
// flag in the status register is set.
func BEQ(c *CPU) {
	branchOnFlag(c, c.StatusRegister().zeroFlag)
}
