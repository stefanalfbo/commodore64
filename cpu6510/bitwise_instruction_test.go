package cpu6510

import "testing"

func TestImmediateAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAImmediate", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAImmediate", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDImmediate", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDImmediate", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORImmediate", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORImmediate", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.ram[1] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestAbsoluteAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAAbsolute", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAAbsolute", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDAbsolute", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDAbsolute", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORAbsolute", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORAbsolute", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestAbsoluteXAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAAbsoluteX", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAAbsoluteX", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDAbsoluteX", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDAbsoluteX", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORAbsoluteX", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORAbsoluteX", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestAbsoluteYAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAAbsoluteY", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAAbsoluteY", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDAbsoluteY", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDAbsoluteY", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORAbsoluteY", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORAbsoluteY", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestZeroPageAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAZeroPage", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAZeroPage", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDZeroPage", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDZeroPage", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORZeroPage", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORZeroPage", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestZeroPageXAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAZeroPageX", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAZeroPageX", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDZeroPageX", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDZeroPageX", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORZeroPageX", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORZeroPageX", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestIndexedIndirectAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAIndexedIndirect", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAIndexedIndirect", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDIndexedIndirect", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDIndexedIndirect", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORIndexedIndirect", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORIndexedIndirect", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestIndirectIndexedAddressingModeInstructions(t *testing.T) {
	tests := []struct {
		instruction  string
		accumulator  byte
		value        byte
		expected     byte
		zeroFlag     bool
		negativeFlag bool
	}{
		{"ORAIndirectIndexed", 0b01010101, 0b10101010, 0b11111111, false, true},
		{"ORAIndirectIndexed", 0b00000000, 0b00000000, 0b00000000, true, false},
		{"ANDIndirectIndexed", 0b10101010, 0b10101010, 0b10101010, false, true},
		{"ANDIndirectIndexed", 0b01010101, 0b10101010, 0b00000000, true, false},
		{"EORIndirectIndexed", 0b00110011, 0b11000011, 0b11110000, false, true},
		{"EORIndirectIndexed", 0b00000001, 0b00000001, 0b00000000, true, false},
	}

	for _, test := range tests {
		cpu := NewCPU()
		cpu.accumulator = test.accumulator
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = test.value

		cpu.execute(InstructionAsHex(test.instruction))

		if cpu.accumulator != test.expected {
			t.Errorf("Accumulator should be %08b, got %08b for instruction: %s", test.expected, cpu.accumulator, test.instruction)
		}

		if cpu.statusRegister.zeroFlag != test.zeroFlag {
			t.Errorf("Zero flag should be %t, got %t for instruction: %s", test.zeroFlag, cpu.statusRegister.zeroFlag, test.instruction)
		}

		if cpu.statusRegister.negativeFlag != test.negativeFlag {
			t.Errorf("Negative flag should be %t, got %t for instruction: %s", test.negativeFlag, cpu.statusRegister.negativeFlag, test.instruction)
		}
	}
}

func TestASLAccumulator(t *testing.T) {
	t.Run("Shift all bits in the accumulator", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0x03

		cpu.execute(InstructionAsHex("ASLAccumulator"))

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

		cpu.execute(InstructionAsHex("ASLAccumulator"))

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

		cpu.execute(InstructionAsHex("ASLAccumulator"))

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

		cpu.execute(InstructionAsHex("ASLAbsolute"))
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

		cpu.execute(InstructionAsHex("ASLAbsolute"))
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

		cpu.execute(InstructionAsHex("ASLAbsolute"))
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

func TestASLZeroPage(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the single byte address", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x03
		cpu.ram[0x03] = 0x03

		cpu.execute(InstructionAsHex("ASLZeroPage"))

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

		cpu.execute(InstructionAsHex("ASLZeroPage"))

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

		cpu.execute(InstructionAsHex("ASLZeroPage"))

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

func TestASLZeroPageX(t *testing.T) {
	t.Run("Shift all bits in the memory location specified by the single byte address and the X register", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[cpu.programCounter+1] = 0x01
		cpu.xRegister = 0x01
		cpu.ram[0x02] = 0x03

		cpu.execute(InstructionAsHex("ASLZeroPageX"))

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

		cpu.execute(InstructionAsHex("ASLZeroPageX"))

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

		cpu.execute(InstructionAsHex("ASLZeroPageX"))

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

		cpu.execute(InstructionAsHex("ASLAbsoluteX"))
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

		cpu.execute(InstructionAsHex("ASLAbsoluteX"))
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

		cpu.execute(InstructionAsHex("ASLAbsoluteX"))

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

func TestLSRAccumulator(t *testing.T) {
	t.Run("Shift all bits right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0b11000000

		cpu.execute(InstructionAsHex("LSRAccumulator"))

		if cpu.accumulator != 0b01100000 {
			t.Errorf("Accumulator should be shifted right, expected 0x60, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits right and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b10000000

		cpu.execute(InstructionAsHex("LSRAccumulator"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Shift all bits right and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01

		cpu.execute(InstructionAsHex("LSRAccumulator"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestLSRAbsolute(t *testing.T) {
	t.Run("Shift all bits right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b11000000

		cpu.execute(InstructionAsHex("LSRAbsolute"))

		if cpu.ram[0x1337] != 0b01100000 {
			t.Errorf("Value should be shifted right, expected 0x60, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits right and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b10000000

		cpu.execute(InstructionAsHex("LSRAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Shift all bits right and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01

		cpu.execute(InstructionAsHex("LSRAbsolute"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestLSRAbsoluteX(t *testing.T) {
	t.Run("Shift all bits right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b11000000

		cpu.execute(InstructionAsHex("LSRAbsoluteX"))

		if cpu.ram[0x1338] != 0b01100000 {
			t.Errorf("Value should be shifted right, expected 0x60, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits right and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10000000

		cpu.execute(InstructionAsHex("LSRAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Shift all bits right and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01

		cpu.execute(InstructionAsHex("LSRAbsoluteX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestLSRZeroPage(t *testing.T) {
	t.Run("Shift all bits right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b11000000

		cpu.execute(InstructionAsHex("LSRZeroPage"))

		if cpu.ram[0x13] != 0b01100000 {
			t.Errorf("Value should be shifted right, expected 0x60, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits right and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b10000000

		cpu.execute(InstructionAsHex("LSRZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Shift all bits right and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x01

		cpu.execute(InstructionAsHex("LSRZeroPage"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestLSRZeroPageX(t *testing.T) {
	t.Run("Shift all bits right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b11000000

		cpu.execute(InstructionAsHex("LSRZeroPageX"))

		if cpu.ram[0x14] != 0b01100000 {
			t.Errorf("Value should be shifted right, expected 0x60, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Shift all bits right and set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b10000000

		cpu.execute(InstructionAsHex("LSRZeroPageX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Shift all bits right and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x01

		cpu.execute(InstructionAsHex("LSRZeroPageX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRORAccumulator(t *testing.T) {
	t.Run("Rotate right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.accumulator = 0b10010001

		cpu.execute(InstructionAsHex("RORAccumulator"))

		if cpu.accumulator != 0b01001000 {
			t.Errorf("Accumulator should be shifted right, expected 0x48, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Rotate right set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001

		cpu.execute(InstructionAsHex("RORAccumulator"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Rotate and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01

		cpu.execute(InstructionAsHex("RORAccumulator"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRORAbsolute(t *testing.T) {
	t.Run("Rotate right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b10010001

		cpu.execute(InstructionAsHex("RORAbsolute"))

		if cpu.ram[0x1337] != 0b01001000 {
			t.Errorf("Value should be shifted right, expected 0x48, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Rotate right set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b00000001

		cpu.execute(InstructionAsHex("RORAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Rotate and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01

		cpu.execute(InstructionAsHex("RORAbsolute"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRORAbsoluteX(t *testing.T) {
	t.Run("Rotate right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 3
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10010001

		cpu.execute(InstructionAsHex("RORAbsoluteX"))

		if cpu.ram[0x1338] != 0b01001000 {
			t.Errorf("Value should be shifted right, expected 0x48, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Rotate right set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b00000001

		cpu.execute(InstructionAsHex("RORAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Rotate and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01

		cpu.execute(InstructionAsHex("RORAbsoluteX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRORZeroPage(t *testing.T) {
	t.Run("Rotate right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b10010001

		cpu.execute(InstructionAsHex("RORZeroPage"))

		if cpu.ram[0x13] != 0b01001000 {
			t.Errorf("Value should be shifted right, expected 0x48, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Rotate right set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b00000001

		cpu.execute(InstructionAsHex("RORZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Rotate and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x01

		cpu.execute(InstructionAsHex("RORZeroPage"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestRORZeroPageX(t *testing.T) {
	t.Run("Rotate right", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 2
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b10010001

		cpu.execute(InstructionAsHex("RORZeroPageX"))

		if cpu.ram[0x14] != 0b01001000 {
			t.Errorf("Value should be shifted right, expected 0x48, got 0x%02x", cpu.accumulator)
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Rotate right set carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b00000001

		cpu.execute(InstructionAsHex("RORZeroPageX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}
	})

	t.Run("Rotate and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x01

		cpu.execute(InstructionAsHex("RORZeroPageX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}
