package cpu6510

const memorySize = 65536

type StatusRegister struct {
	// Indicates when a bit of the result is to be carried to or borrowed
	// from another byte. Also used for rotate and shift operations.
	carry bool
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
}

// NewCPU creates a new CPU6510 processor.
func NewCPU() *CPU {
	return &CPU{
		programCounter: 0,
		statusRegister: StatusRegister{
			carry: false,
		},
	}
}

// Next fetches the next instruction from memory.
func (c *CPU) Next() byte {
	instruction := c.ram[c.programCounter]

	c.programCounter++

	return instruction
}
