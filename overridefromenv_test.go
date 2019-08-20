// Copyright 2019 Carleton University Library
// All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package overridefromenv

import (
	"flag"
	"os"
	"testing"
	"time"
)

func TestOverrideIgnoreSetFlags(t *testing.T) {

	prefix := "OVERRIDEFROMENVTEST_"
	target := prefix + "TEST"

	// Find the old value of a test environment variable and save it.
	old := os.Getenv(target)
	defer os.Setenv(target, old)

	// Set the new value
	os.Setenv(target, "override")

	// Setup the test flag.
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	s := fs.String("test", "default", "")
	fs.Set("test", "newvalue")

	Override(fs, prefix)

	if *s != "newvalue" {
		t.Error("An already set flag was overridden.")
	}
}

func TestOverrideError(t *testing.T) {

	prefix := "OVERRIDEFROMENVTEST_"
	target := prefix + "TEST"

	// Find the old value of a test environment variable and save it.
	old := os.Getenv(target)
	defer os.Setenv(target, old)

	// Set the new value
	os.Setenv(target, "override")

	// Setup the test flag.
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Float64("test", 0.1, "")

	err := Override(fs, prefix)

	if err == nil {
		t.Error("Overriding a float flag with a string didn't cause an error.")
	}
}

func TestOverrideUnsetFlags(t *testing.T) {

	prefix := "OVERRIDEFROMENVTEST_"

	fs := flag.NewFlagSet("test", flag.ExitOnError)

	b := fs.Bool("booltest", true, "")
	booltestold := os.Getenv(prefix + "BOOLTEST")
	defer os.Setenv(prefix+"BOOLTEST", booltestold)
	os.Setenv(prefix+"BOOLTEST", "false")

	defaultduration, _ := time.ParseDuration("1h")
	d := fs.Duration("durationtest", defaultduration, "")
	durationtestold := os.Getenv(prefix + "DURATIONTEST")
	defer os.Setenv(prefix+"DURATIONTEST", durationtestold)
	nd, _ := time.ParseDuration("2h")
	os.Setenv(prefix+"DURATIONTEST", "2h")

	fl := fs.Float64("floattest", 0.1, "")
	floattestold := os.Getenv(prefix + "FLOATTEST")
	defer os.Setenv(prefix+"FLOATTEST", floattestold)
	os.Setenv(prefix+"FLOATTEST", "0.2")

	i := fs.Int("inttest", 1, "")
	inttestold := os.Getenv(prefix + "INTTEST")
	defer os.Setenv(prefix+"INTTEST", inttestold)
	os.Setenv(prefix+"INTTEST", "2")

	i64 := fs.Int64("int64test", 1, "")
	int64testold := os.Getenv(prefix + "INT64TEST")
	defer os.Setenv(prefix+"INT64TEST", int64testold)
	os.Setenv(prefix+"INT64TEST", "2")

	s := fs.String("stringtest", "default", "")
	stringtestold := os.Getenv(prefix + "STRINGTEST")
	defer os.Setenv(prefix+"STRINGTEST", stringtestold)
	os.Setenv(prefix+"STRINGTEST", "newvalue")

	u := fs.Uint64("uinttest", 1, "")
	uinttestold := os.Getenv(prefix + "UINTTEST")
	defer os.Setenv(prefix+"UINTTEST", uinttestold)
	os.Setenv(prefix+"UINTTEST", "2")

	u64 := fs.Uint64("uint64test", 1, "")
	uint64testold := os.Getenv(prefix + "UINT64TEST")
	defer os.Setenv(prefix+"UINT64TEST", uint64testold)
	os.Setenv(prefix+"UINT64TEST", "2")

	Override(fs, prefix)

	if *b != false {
		t.Error("bool flag was not overwritten.")
	}
	if *d != nd {
		t.Error("duration flag was not overwritten.")
	}
	if *fl != 0.2 {
		t.Error("float flag was not overwritten.")
	}
	if *i != 2 {
		t.Error("int flag was not overwritten.")
	}
	if *i64 != 2 {
		t.Error("int64 flag was not overwritten.")
	}
	if *s != "newvalue" {
		t.Error("string flag was not overwritten.")
	}
	if *u != 2 {
		t.Error("uint flag was not overwritten.")
	}
	if *u64 != 2 {
		t.Error("uint64 flag was not overwritten.")
	}
}
