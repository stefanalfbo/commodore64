package cpu6510

var OP_CODE = map[string]byte{
	"BRK": 0x00,
	"CLC": 0x18,
	"SEC": 0x38,
}

// BRK - BReaKpoint. BRK is intended for use as a debugging tool which
// a programmer may place at specific points in a program, to check the state
// of processor flags at these points in the code.
func (c *CPU) BRK() {
	c.statusRegister.breakCommandFlag = true
	c.statusRegister.interruptDisableFlag = true

	// BRK increments the program counter by 2 instead of 1, therefore we
	// increment it one more time here.
	c.programCounter++

}

// CLC - CLear Carry
func (c *CPU) CLC() {
	c.statusRegister.carry = false
}

// SEC - SEt Carry
func (c *CPU) SEC() {
	c.statusRegister.carry = true
}
