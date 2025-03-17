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
		address := ConvertTwoBytesToAddress(test.high, test.low)
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

func TestNOP(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("NOP"))

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}
