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
