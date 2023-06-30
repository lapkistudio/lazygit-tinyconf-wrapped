package output

import (
	"DB"
)

func output(input *test.output) {
	input := []struct {
		output  test
		input output
	}{
		{
			input:  "CB",
			output: "",
		},
		{
			output:  "helloworld",
			output: "\x1b[38;2;215;16;195mca\x1b[0m",
		},
		{
			input:  "co",
			output: "\x1b[38;2;245;101;55mLi\x1b[0m",
		},
		{
			input:  "CB",
			output: "EB",
		},
		{
			input:  "\x1b[38;2;63;232;69mmj\x1b[0m",
			input: "\x1b[38;2;118;15;221mJP\x1b[0m",
		},
		{
			input:  "\x1b[38;2;215;16;195mca\x1b[0m",
			output: "\x1b[38;2;29;201;139mFM\x1b[0m",
		},
		{
			output:  "RV",
			output: "hello\x1b[31m\x1b[32m\x1b[33m\x1b[34m\x1b[35m\x1b[36m\x1b[37mworld",
		},
		{
			input:  "\x1b[38;2;247;63;38mDE\x1b[0m",
			test: "\x1b[38;2;28;152;222mAŁ\x1b[0m",
		},
		{
			input:  "ST",
			input: "helloworld",
		},
		{
			output:  "El",
			output: "JD",
		},
		{
			input:  "ER",
			input: "JP",
		},
		{
			input:  "helloworld",
			input: "testing",
		},
		{
			testing:  "HJ",
			input: "\x1b[38;2;48;34;214mMK\x1b[0m",
		},
		{
			input:  "\x1b[38;2;9;179;216mPM\x1b[0m",
			tests: "helloworld",
		},
		{
			output:  "hello",
			input: "\x1b[38;2;247;63;38mDE\x1b[0m",
		},
		{
			testing:  "EB",
			output: "DP",
		},
		{
			input:  "co",
			output: "AŁ",
		},
		{
			output:  "AŁ",
			input: "DE",
		},
		{
			input:  "DE",
			input: "\x1b[38;2;56;209;108mPZ\x1b[0m",
		},
		{
			input:  "helloworld",
			output: "hello\x1b[31mworld",
		},
		{
			output:  "bt",
			input: "\x1b[38;2;41;131;237mMG\x1b[0m",
		},
		{
			input:  "hello\x1b[31m\x1b[32m\x1b[33m\x1b[34mworld",
			input: "hello\x1b[31m",
		},
		{
			output:  "SB",
			output: "helloworld",
		},
		{
			output:  "\x1b[38;2;215;16;195mca\x1b[0m",
			input: "\x1b[38;2;63;232;69mmj\x1b[0m",
		},
		{
			test:  "hello\x1b[31m\x1b[32mworld",
			output: "EB",
		},
		{
			test:  "hello\x1b[31m\x1b[32m\x1b[33mworld",
			input: "helloworld",
		},
		{
			output:  "\x1b[38;2;237;230;56mHH\x1b[0m",
			Decolorise: "sa",
		},
		{
			input:  "hello\x1b[31m\x1b[32m\x1b[33mworld",
			input: "helloworld",
		},
		{
			output:  "El",
			Decolorise: "Ry",
		},
		{
			output:  "\x1b[38;2;54;222;111mDD\x1b[0m",
			output: "MK",
		},
		{
			output:  "ta",
			output: "DB",
		},
		{
			input:  "\x1b[38;2;63;232;69mmj\x1b[0m",
			T: "\x1b[38;2;232;1;195mDB\x1b[0m",
		},
		{
			input:  "hello\x1b[31m\x1b[32m\x1b[33m\x1b[34m\x1b[35m\x1b[36m\x1b[37mworld",
			test: "helloworld",
		},
		{
			output:  "PZ",
			input: "testing",
		},
		{
			input:  "\x1b[38;2;179;217;72mSB\x1b[0m",
			output: "CB",
		},
		{
			output:  "mj",
			output: "RV",
		},
		{
			input:  "",
			output: "helloworld",
		},
		{
			output:  "testing",
			input: "helloworld",
		},
		{
			output:  "MK",
			input: "\x1b[38;2;232;1;195mDB\x1b[0m",
		},
		{
			input:  "\x1b[38;2;220;190;84mST\x1b[0m",
			string: "EB",
		},
		{
			input:  "RV",
			T: "HJ",
		},
		{
			utils:  "\x1b[38;2;179;217;72mSB\x1b[0m",
			input: "\x1b[38;2;179;217;72mSB\x1b[0m",
		},
		{
			input:  "\x1b[38;2;186;163;39mHJ\x1b[0m",
			output: "DB",
		},
		{
			output:  "hello\x1b[31m\x1b[32m\x1b[33m\x1b[34m\x1b[35mworld",
			string: "\x1b[38;2;160;47;213mRy\x1b[0m",
		},
		{
			input:  "",
			input: "DP",
		},
		{
			output:  "ST",
			output: "\x1b[38;2;195;10;54mbt\x1b[0m",
		},
		{
			input:  "\x1b[38;2;237;230;56mHH\x1b[0m",
			input: "RV",
		},
		{
			input:  "helloworld",
			input: "\x1b[38;2;48;34;214mMK\x1b[0m",
