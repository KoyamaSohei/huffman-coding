package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	c := &cobra.Command{
		Use:   "huffman-coding",
		Short: "huffman encoding / decoding tool",
		Long: `huffman encoding / decoding tool
refer: https://en.wikipedia.org/wiki/Huffman_coding`,
	}
	e := &cobra.Command{
		Use:   "encode [input file name] [output file name]",
		Short: "huffman encoding",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			i, err := ioutil.ReadFile(args[0])
			if err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error: cannot read %s\n", args[0])))
				os.Exit(1)
			}
			name := args[1]
			k, b, err := encode(string(i))
			if err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error occured on encoding\n")))
				os.Exit(1)
			}
			if err := ioutil.WriteFile(name+".json", k, 0644); err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error occured on save key\n")))
				os.Exit(1)
			}
			if err := ioutil.WriteFile(name, b, 0644); err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error occured on save text\n")))
				os.Exit(1)
			}
		},
	}
	d := &cobra.Command{
		Use:   "decode [encoded name]",
		Short: "huffman encoding",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t, err := ioutil.ReadFile(args[0])
			if err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error: cannot read %s\n", args[0])))
				os.Exit(1)
			}
			k, err := ioutil.ReadFile(args[0] + ".json")
			if err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error: cannot read %s.json\n", args[0])))
				os.Exit(1)
			}
			res, err := decode(k, t)
			if err != nil {
				w := cmd.ErrOrStderr()
				w.Write([]byte(fmt.Sprintf("Error occured on decoding\n")))
				os.Exit(1)
			}
			w := cmd.OutOrStdout()
			w.Write([]byte(res))
		},
	}
	c.AddCommand(e, d)
	c.Execute()
}
