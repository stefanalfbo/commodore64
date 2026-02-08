package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	original := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	os.Stdout = writer

	fn()

	writer.Close()
	os.Stdout = original

	out, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	return string(out)
}

func TestGetByteAddress(t *testing.T) {
	buffer := bytes.NewReader([]byte{0xAB})
	addr, err := getByteAddress(buffer)
	if err != nil {
		t.Fatalf("getByteAddress error: %v", err)
	}
	if addr != 0xAB {
		t.Fatalf("expected 0xAB, got 0x%02X", addr)
	}
}

func TestGetWordAddress(t *testing.T) {
	buffer := bytes.NewReader([]byte{0x34, 0x12})
	addr, err := getWordAddress(buffer)
	if err != nil {
		t.Fatalf("getWordAddress error: %v", err)
	}
	if addr != 0x1234 {
		t.Fatalf("expected 0x1234, got 0x%04X", addr)
	}
}

func TestRunDisassemblesSequence(t *testing.T) {
	// LDA #$10, ORA ($20), X, LDA $44, X, STA $5678, X, ASL A, BRK
	buffer := bytes.NewReader([]byte{
		0xA9, 0x10,
		0x01, 0x20,
		0xB5, 0x44,
		0x9D, 0x78, 0x56,
		0x0A,
		0x00,
	})

	output := captureStdout(t, func() {
		if err := run(buffer); err != nil {
			t.Fatalf("run error: %v", err)
		}
	})

	expected := "" +
		"LDA #$10\n" +
		"ORA ($20), X\n" +
		"LDA $44, X\n" +
		"STA $5678, X\n" +
		"ASL A\n" +
		"BRK\n"

	if output != expected {
		t.Fatalf("unexpected output:\nexpected:\n%q\ngot:\n%q", expected, output)
	}
}

func TestDisassembleUnknownInstructionPanics(t *testing.T) {
	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("expected panic for unknown instruction")
		}
	}()

	disassemble(bytes.NewReader(nil), 0x02)
}
