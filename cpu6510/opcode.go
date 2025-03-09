package cpu6510

// CLC - CLear Carry
func (c *CPU) CLC() {
	c.statusRegister.carry = false
}
