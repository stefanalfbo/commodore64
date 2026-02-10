package cpu6510

// JAM - JaM/KIL. Illegal opcode that halts the CPU.
func JAM(c *CPU) {
	c.isJammed = true
}
