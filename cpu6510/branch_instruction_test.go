package cpu6510

import (
	"testing"
)

func TestBPL(t *testing.T) {

	t.Run("Branch if Positive", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BPL"))

		if cpu.programCounter != 0x43 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Negative", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BPL"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})
}

func TestBMI(t *testing.T) {
	t.Run("Branch if Negative", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BMI"))

		if cpu.programCounter != 0x43 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Positive", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BMI"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})
}

func TestBVC(t *testing.T) {
	t.Run("Branch if Overflow Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BVC"))

		if cpu.programCounter != 0x43 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Overflow Set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BVC"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})
}

func TestBVS(t *testing.T) {
	t.Run("Branch if Overflow Set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BVS"))

		if cpu.programCounter != 0x43 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Overflow Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BVS"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})
}

func TestBCC(t *testing.T) {
	t.Run("Branch if Carry Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BCC"))

		if cpu.programCounter != 0x43 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Carry Set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BCC"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})
}
