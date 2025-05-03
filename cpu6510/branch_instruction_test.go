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

		if cpu.programCounter != 0x44 {
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

	t.Run("Branch if Positive and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = false
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BPL"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBMI(t *testing.T) {
	t.Run("Branch if Negative", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BMI"))

		if cpu.programCounter != 0x44 {
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

	t.Run("Branch if Negative and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.negativeFlag = true
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BMI"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBVC(t *testing.T) {
	t.Run("Branch if Overflow Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = false
		cpu.ram[cpu.programCounter+1] = 0x44

		cpu.execute(InstructionAsHex("BVC"))

		if cpu.programCounter != 0x46 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x46, cpu.programCounter)
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

	t.Run("Branch if Overflow Clear and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = false
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BVC"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBVS(t *testing.T) {
	t.Run("Branch if Overflow Set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BVS"))

		if cpu.programCounter != 0x44 {
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

	t.Run("Branch if Overflow Set and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.overflowFlag = true
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BVS"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBCC(t *testing.T) {
	t.Run("Branch if Carry Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BCC"))

		if cpu.programCounter != 0x44 {
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

	t.Run("Branch if Carry Clear and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = false
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BCC"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBCS(t *testing.T) {
	t.Run("Branch if Carry Set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BCS"))

		if cpu.programCounter != 0x44 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Carry Clear", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BCS"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})

	t.Run("Branch if Carry Set and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carryFlag = true
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BCS"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBNE(t *testing.T) {
	t.Run("Branch if Not Equal", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BNE"))

		if cpu.programCounter != 0x44 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Equal", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BNE"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})

	t.Run("Branch if Not Equal and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = false
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BNE"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}

func TestBEQ(t *testing.T) {
	t.Run("Branch if Equal", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = true
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BEQ"))

		if cpu.programCounter != 0x44 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x44, cpu.programCounter)
		}
	})

	t.Run("Do Not Branch if Not Equal", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = false
		cpu.ram[cpu.programCounter+1] = 0x42

		cpu.execute(InstructionAsHex("BEQ"))
		if cpu.programCounter != 0x01 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0x01, cpu.programCounter)
		}
	})

	t.Run("Branch if Equal and jump to negative address", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.zeroFlag = true
		cpu.programCounter = 0xC002
		cpu.ram[cpu.programCounter+1] = 0xFC

		cpu.execute(InstructionAsHex("BEQ"))

		if cpu.programCounter != 0xC000 {
			t.Errorf("Expected PC to be 0x%02X, got 0x%02X", 0xC000, cpu.programCounter)
		}
	})
}
