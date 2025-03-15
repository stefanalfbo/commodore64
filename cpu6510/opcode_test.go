package cpu6510

import (
	"testing"
)

func TestConvertTwoBytesToAddress(t *testing.T) {
	tests := []struct {
		high     byte
		low      byte
		expected uint16
	}{
		{0x03, 0x00, 0x0300},
		{0x00, 0xFF, 0x00FF},
		{0xFF, 0xFF, 0xFFFF},
		{0x12, 0x34, 0x1234},
		{0xAB, 0xCD, 0xABCD},
		{0x00, 0x00, 0x0000},
		{0x02, 0x00, 0x0200},
	}

	for _, test := range tests {
		address := convertTwoBytesToAddress(test.high, test.low)
		if address != test.expected {
			t.Errorf("Address should be %04x, got %04x for high:%02x low:%02x", test.expected, address, test.high, test.low)
		}
	}
}

func TestBRK(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 2

	cpu.execute(OpCodeAsHex("BRK"))

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if !cpu.statusRegister.breakCommandFlag {
		t.Errorf("Break status flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented by 2")
	}
}

func TestASLZeroPage(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the single byte address", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x03
		cpu.ram[0x03] = 0x03

		cpu.execute(OpCodeAsHex("ASLZeroPage"))

		if cpu.ram[0x03] != 0x06 {
			t.Errorf("Memory location should be shifted left")
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the single byte address and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x03
		cpu.ram[0x03] = 0x80

		cpu.execute(OpCodeAsHex("ASLZeroPage"))

		if cpu.ram[0x03] != 0x00 {
			t.Errorf("Memory location should be shifted left")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the single byte address and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[cpu.programCounter+1] = 0x03
		cpu.ram[0x03] = 0x40

		cpu.execute(OpCodeAsHex("ASLZeroPage"))

		if cpu.ram[0x03] != 0x80 {
			t.Errorf("Memory location should be shifted left")
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestASLAccumulator(t *testing.T) {
	t.Run("Shift all bits in the accumulator", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0x03

		cpu.execute(OpCodeAsHex("ASLAccumulator"))

		if cpu.accumulator != 0x06 {
			t.Errorf("Accumulator should be shifted left")
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the accumulator and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0x80

		cpu.execute(OpCodeAsHex("ASLAccumulator"))

		if cpu.accumulator != 0x00 {
			t.Errorf("Accumulator should be shifted left")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the accumulator and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0x40

		cpu.execute(OpCodeAsHex("ASLAccumulator"))

		if cpu.accumulator != 0x80 {
			t.Errorf("Accumulator should be shifted left")
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestASLAbsolute(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the two byte address", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.ram[0x0201] = 0x03

		cpu.execute(OpCodeAsHex("ASLAbsolute"))
		if cpu.ram[0x0201] != 0x06 {
			t.Errorf("Memory location should be shifted left, expected 0x06, got %02x", cpu.ram[0x0201])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the two byte address and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.ram[0x0201] = 0x80

		cpu.execute(OpCodeAsHex("ASLAbsolute"))
		if cpu.ram[0x0201] != 0x00 {
			t.Errorf("Memory location should be shifted left, expected 0x00, got %02x", cpu.ram[0x0200])
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the two byte address and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.ram[0x0201] = 0x40

		cpu.execute(OpCodeAsHex("ASLAbsolute"))
		if cpu.ram[0x0201] != 0x80 {
			t.Errorf("Memory location should be shifted left, expected 0x80, got %02x", cpu.ram[0x0200])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestPHP(t *testing.T) {
	t.Run("Push processor status register flags", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.statusRegister.carryFlag = true
		cpu.statusRegister.zeroFlag = true
		cpu.statusRegister.interruptDisableFlag = true
		cpu.statusRegister.decimalModeFlag = true
		cpu.statusRegister.breakCommandFlag = true
		cpu.statusRegister.overflowFlag = true
		cpu.statusRegister.negativeFlag = true

		cpu.execute(OpCodeAsHex("PHP"))

		if cpu.ram[0x01FF] != 0xFF {
			t.Errorf("Status register should be pushed onto the stack")
		}

		if cpu.stackPointer != 0xFE {
			t.Errorf("Stack pointer should be decremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Push processor status register flags when all is cleared", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1

		cpu.execute(OpCodeAsHex("PHP"))

		if cpu.ram[0x01FF] != 0x20 {
			t.Errorf("Status register should be pushed onto the stack")
		}

		if cpu.stackPointer != 0xFE {
			t.Errorf("Stack pointer should be decremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestASLZeroPageX(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the single byte address and the X register", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.xRegister = 0x01
		cpu.ram[0x02] = 0x03

		cpu.execute(OpCodeAsHex("ASLZeroPageX"))

		if cpu.ram[0x02] != 0x06 {
			t.Errorf("Memory location should be shifted left, expected 0x06, got %02x", cpu.ram[0x02])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the single byte address and the X register and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.xRegister = 0x01
		cpu.ram[0x02] = 0x80

		cpu.execute(OpCodeAsHex("ASLZeroPageX"))

		if cpu.ram[0x02] != 0x00 {
			t.Errorf("Memory location should be shifted left, expected 0x00, got %02x", cpu.ram[0x02])
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the single byte address and the X register and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.xRegister = 0x01
		cpu.ram[0x02] = 0x40

		cpu.execute(OpCodeAsHex("ASLZeroPageX"))

		if cpu.ram[0x02] != 0x80 {
			t.Errorf("Memory location should be shifted left, expected 0x80, got %02x", cpu.ram[0x02])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestCLC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.carryFlag = true

	cpu.execute(OpCodeAsHex("CLC"))

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestASLAbsoluteX(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the two byte address and the X register", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.xRegister = 0x01

		cpu.ram[0x0202] = 0x03

		cpu.execute(OpCodeAsHex("ASLAbsoluteX"))
		if cpu.ram[0x0202] != 0x06 {
			t.Errorf("Memory location should be shifted left, expected 0x06, got %02x", cpu.ram[0x0202])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the two byte address and the X register and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.xRegister = 0x01

		cpu.ram[0x0202] = 0x80

		cpu.execute(OpCodeAsHex("ASLAbsoluteX"))
		if cpu.ram[0x0202] != 0x00 {
			t.Errorf("Memory location should be shifted left, expected 0x00, got %02x", cpu.ram[0x0202])
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits in the memory location specified by the two byte address and the X register and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.ram[cpu.programCounter+2] = 0x02
		cpu.xRegister = 0x01

		cpu.ram[0x0202] = 0x40

		cpu.execute(OpCodeAsHex("ASLAbsoluteX"))

		if cpu.ram[0x0202] != 0x80 {
			t.Errorf("Memory location should be shifted left, expected 0x80, got %02x", cpu.ram[0x0202])
		}

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestPLP(t *testing.T) {
	t.Run("Pull processor status register flags when all is cleared", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x20
		cpu.stackPointer = 0xFE

		cpu.execute(OpCodeAsHex("PLP"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.interruptDisableFlag {
			t.Errorf("Interrupt disable flag should be cleared")
		}

		if cpu.statusRegister.decimalModeFlag {
			t.Errorf("Decimal mode flag should be cleared")
		}

		if cpu.statusRegister.breakCommandFlag {
			t.Errorf("Break command flag should be cleared")
		}

		if cpu.statusRegister.overflowFlag {
			t.Errorf("Overflow flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull processor status register flags when all is set", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0xFF
		cpu.stackPointer = 0xFE

		cpu.execute(OpCodeAsHex("PLP"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if !cpu.statusRegister.interruptDisableFlag {
			t.Errorf("Interrupt disable flag should be set")
		}

		if !cpu.statusRegister.decimalModeFlag {
			t.Errorf("Decimal mode flag should be set")
		}

		if !cpu.statusRegister.breakCommandFlag {
			t.Errorf("Break command flag should be set")
		}

		if !cpu.statusRegister.overflowFlag {
			t.Errorf("Overflow flag should be set")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestSEC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SEC"))

	if !cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestPHA(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.accumulator = 0x03

	cpu.execute(OpCodeAsHex("PHA"))

	if cpu.ram[0x01FF] != cpu.accumulator {
		t.Errorf("Accumulator should be pushed onto the stack")
	}

	if cpu.stackPointer != 0xFE {
		t.Errorf("Stack pointer should be decremented")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLI(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.interruptDisableFlag = true

	cpu.execute(OpCodeAsHex("CLI"))

	if cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestRTS(t *testing.T) {
	t.Run("Return from subroutine", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[0x01FE] = 0x03
		cpu.ram[0x01FF] = 0x00
		cpu.stackPointer = 0xFD

		cpu.execute(OpCodeAsHex("RTS"))

		if cpu.programCounter != 0x0004 {
			t.Errorf("Program counter should be set to the address on the stack")
		}
	})
}

func TestPLA(t *testing.T) {
	t.Run("Pull accumulator", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x03
		cpu.stackPointer = 0xFE

		cpu.execute(OpCodeAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull accumulator and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x00
		cpu.stackPointer = 0xFE

		cpu.execute(OpCodeAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull accumulator and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x80
		cpu.stackPointer = 0xFE

		cpu.execute(OpCodeAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestSEI(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SEI"))

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLV(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.overflowFlag = true

	cpu.execute(OpCodeAsHex("CLV"))

	if cpu.statusRegister.overflowFlag {
		t.Errorf("Overflow flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestNOP(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("NOP"))

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestTXS(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.xRegister = 0x03

	cpu.execute(OpCodeAsHex("TXS"))

	if cpu.stackPointer != cpu.xRegister {
		t.Errorf("Stack pointer should be set to the value of the X register")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLD(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.decimalModeFlag = true

	cpu.execute(OpCodeAsHex("CLD"))

	if cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSED(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SED"))

	if !cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestTSX(t *testing.T) {

	t.Run("Transfer stack pointer to X", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.stackPointer = 0x03

		cpu.execute(OpCodeAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Transfer stack pointer to X and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.stackPointer = 0x00

		cpu.execute(OpCodeAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Transfer stack pointer to X and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.stackPointer = 0x80

		cpu.execute(OpCodeAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRegisterInstructionsImpliedMode(t *testing.T) {
	t.Run("DEY - Decrement Y register", func(t *testing.T) {
		t.Run("Decrement Y register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x03

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0x02 {
				t.Errorf("Y register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement Y register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x01

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be decremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement Y register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x00

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0xFF {
				t.Errorf("Y register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TXA - Transfer X to A", func(t *testing.T) {
		t.Run("Transfer X to A", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x03

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer X to A and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x00

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer X to A and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x80

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TYA - Transfer Y to A", func(t *testing.T) {
		t.Run("Transfer Y to A", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x03

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer Y to A and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x00

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer Y to A and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x80

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TAY - Transfer A to Y", func(t *testing.T) {
		t.Run("Transfer A to X", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x03

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x03 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to Y and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x00

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to Y and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x80

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x80 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TAX - Transfer A to X", func(t *testing.T) {
		t.Run("Transfer A to X", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x03

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x03 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to X and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x00

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to X and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x80

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x80 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("INY - Increment Y register", func(t *testing.T) {
		t.Run("Increment Y register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x01

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x02 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment Y register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0xFF

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be incremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment Y register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x7F

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x80 {
				t.Errorf("Y register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("DEX - Decrement X register", func(t *testing.T) {
		t.Run("Decrement X register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x03

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0x02 {
				t.Errorf("X register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement X register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x01

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be decremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement X register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x00

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0xFF {
				t.Errorf("X register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("INX - Increment X register", func(t *testing.T) {
		t.Run("Increment X register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x01

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x02 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment X register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0xFF

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be incremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment X register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x7F

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x80 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})
}
