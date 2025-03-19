package cpu6510

import "testing"

func TestCMPImmediate(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(OpCodeAsHex("CMPImmediate"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsolute(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(OpCodeAsHex("CMPAbsolute"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsoluteY(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsoluteY"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(OpCodeAsHex("CMPAbsoluteY"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsoluteY"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsoluteX(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(OpCodeAsHex("CMPAbsoluteX"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(OpCodeAsHex("CMPAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}
