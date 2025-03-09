package cpu6510

// BRK - BReaKpoint, 0x00. BRK is intended for use as a debugging tool which
// a programmer may place at specific points in a program, to check the state
// of processor flags at these points in the code.
func (c *CPU) BRK() {
	c.statusRegister.breakCommandFlag = true
	c.statusRegister.interruptDisableFlag = true

	// BRK increments the program counter by 2 instead of 1, therefore we
	// increment it one more time here.
	c.programCounter++

}

// CLC - CLear Carry, 0x18
func (c *CPU) CLC() {
	c.statusRegister.carry = false
}
