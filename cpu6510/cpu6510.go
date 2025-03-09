package cpu6510

const memorySize = 65536

type StatusRegister struct {
	// Indicates when a bit of the result is to be carried to or borrowed
	// from another byte. Also used for rotate and shift operations.
	carry bool
	// If set IRQ will be prevented (masked), except non-maskable
	// interrupts (NMI).
	interruptDisableFlag bool
	// If set arithmetic operations are calculated in decimal mode
	// (otherwise usually in binary mode).
	decimalModeFlag bool
	// Indicates that interrupt request has been triggered by an BRK
	// opcode (not an IRQ).
	breakCommandFlag bool
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
			carry:                false,
			interruptDisableFlag: false,
			decimalModeFlag:      false,
			breakCommandFlag:     false,
		},
	}
}

// Next fetches the next instruction from memory.
func (c *CPU) next() byte {
	instruction := c.ram[c.programCounter]

	return instruction
}

// Execute executes the instruction.
func (c *CPU) execute(instruction byte) {
	switch instruction {
	case 0x00:
		c.BRK()
	case 0x18:
		c.CLC()
	case 0x38:
		c.SEC()
	case 0xD8:
		c.CLD()
	case 0xF8:
		c.SED()
	default:
		panic("Unknown instruction")
	}

	// Should the program counter be incremented here? See the BRK opcode.
	c.programCounter++
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
