package cpu6510

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

func cpx(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	tmp := c.xRegister - value

	raiseStatusRegisterFlags(c, tmp)

	c.statusRegister.carryFlag = c.xRegister >= value
}

// CPXImmediate - CoMPare X register. CPX compares the value in the X index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPXImmediate(c *CPU) {
	cpx(c, c.getValueByImmediateAddressingMode)
}

// CPXAbsolute - CoMPare X register. CPX compares the value in the X index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPXAbsolute(c *CPU) {
	cpx(c, c.getValueByAbsoluteAddressingMode)
}

// CPXZeroPage - CoMPare X register. CPX compares the value in the X index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPXZeroPage(c *CPU) {
	cpx(c, c.getValueByZeroPageAddressingMode)
}

func cpy(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	tmp := c.yRegister - value

	raiseStatusRegisterFlags(c, tmp)

	c.statusRegister.carryFlag = c.yRegister >= value
}

// CPYImmediate - CoMPare Y register. CPY compares the value in the Y index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPYImmediate(c *CPU) {
	cpy(c, c.getValueByImmediateAddressingMode)
}

// CPYAbsolute - CoMPare Y register. CPY compares the value in the Y index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPYAbsolute(c *CPU) {
	cpy(c, c.getValueByAbsoluteAddressingMode)
}

// CPYZeroPage - CoMPare Y register. CPY compares the value in the Y index
// register with the value in memory, and sets the zero and negative flags in
// the status register based on the result.
func CPYZeroPage(c *CPU) {
	cpy(c, c.getValueByZeroPageAddressingMode)
}

// BITZeroPage - BIT sets the zero flag as though the value in the address tested
// were ANDed with the accumulator. The negative flag and overflow flag are set to
// bits 7 and 6 respectively of the value in the address tested.
func BITZeroPage(c *CPU) {
	bit(c, c.getValueByZeroPageAddressingMode)
}

// BITZeroPage - BIT sets the zero flag as though the value in the address tested
// were ANDed with the accumulator. The negative flag and overflow flag are set to
// bits 7 and 6 respectively of the value in the address tested.
func BITAbsolute(c *CPU) {
	bit(c, c.getValueByAbsoluteAddressingMode)
}

// BIT - BIT sets the zero flag as though the value in the address tested
// were ANDed with the accumulator. The negative flag and overflow flag are set to
// bits 7 and 6 respectively of the value in the address tested.
func bit(c *CPU, getValue func() byte) {
	c.programCounter++

	value := getValue()

	c.statusRegister.overflowFlag = value&0b01000000 != 0
	c.statusRegister.negativeFlag = value&0b10000000 != 0

	c.statusRegister.zeroFlag = value&c.accumulator == 0
}
