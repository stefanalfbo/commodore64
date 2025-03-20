package cpu6510

import "fmt"

// The memory size of the CPU6510 is 64KB (65536 Bytes).
const memorySize = 65536

// The stack is located in the memory range $0100-$01FF.
const stackBase uint16 = 0x0100

type StatusRegister struct {
	// Indicates when a bit of the result is to be carried to or borrowed
	// from another byte. Also used for rotate and shift operations.
	carryFlag bool // C
	// True indicates that the result of an operation is equal to zero.
	zeroFlag bool // Z
	// If set IRQ will be prevented (masked), except non-maskable
	// interrupts (NMI).
	interruptDisableFlag bool // I
	// If set arithmetic operations are calculated in decimal mode
	// (otherwise usually in binary mode).
	decimalModeFlag bool // D
	// Indicates that interrupt request has been triggered by an BRK
	// instruction (not an IRQ).
	breakCommandFlag bool // B
	// Cannot be changed, usually set to 1.
	unusedFlag bool
	// Indicates that a result of an signed arithmetic operation
	// exceeds the signed value range (-128 to 127).
	overflowFlag bool // V
	// A value of true indicates that the result is negative (bit 7 is
	// set, for a two's complement representation).
	negativeFlag bool // N
}

// newStatusRegister creates a new status register with the given value.
func newStatusRegister(value byte) StatusRegister {
	return StatusRegister{
		carryFlag:            value&0x01 == 0x01,
		zeroFlag:             value&0x02 == 0x02,
		interruptDisableFlag: value&0x04 == 0x04,
		decimalModeFlag:      value&0x08 == 0x08,
		breakCommandFlag:     value&0x10 == 0x10,
		unusedFlag:           value&0x20 == 0x20,
		overflowFlag:         value&0x40 == 0x40,
		negativeFlag:         value&0x80 == 0x80,
	}
}

// asByte returns the status register as a byte.
func (sr *StatusRegister) asByte() byte {
	var value byte

	if sr.carryFlag {
		value |= 1 << 0
	}
	if sr.zeroFlag {
		value |= 1 << 1
	}
	if sr.interruptDisableFlag {
		value |= 1 << 2
	}
	if sr.decimalModeFlag {
		value |= 1 << 3
	}
	if sr.breakCommandFlag {
		value |= 1 << 4
	}
	if sr.unusedFlag {
		value |= 1 << 5
	}
	if sr.overflowFlag {
		value |= 1 << 6
	}
	if sr.negativeFlag {
		value |= 1 << 7
	}

	return value
}

// The CPU struct represents the CPU6510 processor.
type CPU struct {
	// The Program Counter, also known as Instruction Pointer, IP, or PC is a
	// 16 bit register in the processor that contains the address in RAM of the
	// currently executing instruction on the CPU6510.
	programCounter uint16
	// The Status Register is a hardware register which records the condition
	// of the CPU as a result of arithmetic, logical or command operations.
	// The purpose of the Processor Status Register is to hold information
	// about the most recently performed ALU operation, control the enabling
	// and disabling of interrupts and set the CPU operating mode.
	statusRegister StatusRegister
	// The Random Access Memory (RAM) is a 64KB (65536 Bytes) memory that
	// stores the program and data that the CPU6510 will execute.
	ram [memorySize]byte
	// The X index register is an 8-Bit data register. The register is used
	// in the Indexed Indirect, and Absolute indexed by X addressing modes.
	xRegister byte
	// The Y index register is an 8-Bit data register. The register is used
	// in the Indirect Indexed, and Absolute indexed by Y addressing modes.
	// Also used in conjuction with the Accumulator (A) to form either memory
	// address locations (A/Y) or 16-Bit signed values (A/Y) for Floating
	// point arithmetic.
	yRegister byte
	// The accumulator (abbreviated as A) is a 8-bit data register.The
	// accumulator primarily serves as register for arithmetic and logical
	// operations.
	accumulator byte
	// The Stack Pointer is a 8-bit register. The stack pointer always points
	// to the next available location on the stack.
	//
	// The first entry in the stack is $01FF and the last entry is $0100. When
	// the stack is empty SP = $FF.
	stackPointer byte
}

// NewCPU creates a new CPU6510 processor.
func NewCPU() *CPU {
	return &CPU{
		programCounter: 0,
		statusRegister: StatusRegister{
			carryFlag:            false,
			zeroFlag:             false,
			interruptDisableFlag: false,
			decimalModeFlag:      false,
			breakCommandFlag:     false,
			unusedFlag:           true,
			overflowFlag:         false,
			negativeFlag:         false,
		},
		xRegister:    0,
		yRegister:    0,
		accumulator:  0,
		stackPointer: 0xFF,
	}
}

func (c *CPU) pushOnStack(value byte) {
	c.ram[stackBase+uint16(c.stackPointer)] = value
	c.stackPointer--
}

func (c *CPU) popFromStack() byte {
	c.stackPointer++
	return c.ram[stackBase+uint16(c.stackPointer)]
}

// Next fetches the next instruction from memory.
func (c *CPU) next() byte {
	instruction := c.ram[c.programCounter]

	return instruction
}

// readAddressFromMemory reads the address from the next two bytes in memory.
func (c *CPU) readAddressFromMemory() uint16 {
	var lowByte byte = c.ram[c.programCounter]
	var highByte byte = c.ram[c.programCounter+1]

	return ConvertTwoBytesToAddress(highByte, lowByte)
}

// readMemory reads the byte at the given address in memory.
func (c *CPU) readMemory(address uint16) byte {
	return c.ram[address]
}

// writeMemory writes the byte at the given address in memory.
func (c *CPU) writeMemory(address uint16, value byte) {
	c.ram[address] = value
}

// Execute executes the instruction.
func (c *CPU) execute(instruction byte) {
	if runInstruction, ok := lookupInstruction[instruction]; ok {
		runInstruction(c)
	} else {
		panic(fmt.Sprintf("Unknown instruction, %x", instruction))
	}
}

// Run the CPU.
func (c *CPU) Run() {
	for {
		instruction := c.next()
		c.execute(instruction)

		// Exit the loop when the instruction is BRK, 0x00.
		if instruction == 0x00 {
			break
		}
	}
}
