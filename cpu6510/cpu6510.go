package cpu6510

import "fmt"

const memorySize = 65536

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
	// opcode (not an IRQ).
	breakCommandFlag bool // B
	// Indicates that a result of an signed arithmetic operation
	// exceeds the signed value range (-128 to 127).
	overflowFlag bool // V
	// A value of true indicates that the result is negative (bit 7 is
	// set, for a two's complement representation).
	negativeFlag bool // N
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
			overflowFlag:         false,
			negativeFlag:         false,
		},
		xRegister:   0,
		yRegister:   0,
		accumulator: 0,
	}
}

// Next fetches the next instruction from memory.
func (c *CPU) next() byte {
	instruction := c.ram[c.programCounter]

	return instruction
}

// Execute executes the instruction.
func (c *CPU) execute(instruction byte) {
	if runOpCode, ok := lookupOpCode[instruction]; ok {
		runOpCode(c)
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
