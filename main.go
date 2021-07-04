package main

import "github.com/jnst/x-go/cmd"

func main() {
	cmd.Execute()

	//flag.Parse()
	//args := flag.Args()
	//if len(args) == 0 {
	//	_, _ = fmt.Fprintln(os.Stderr, "specify the command name as argument")
	//	return
	//}
	//
	//switch command := strings.ToLower(args[0]); command {
	//case "factorial":
	//	code.Factorial(10)
	//case "slack":
	//	sample.Post()
	//case "solana-keygen":
	//	solana.GenerateVanityAddress(args[1])
	//case "sort":
	//	std.PrintSortedHeroes()
	//	std.PrintSortedYearMonths()
	//case "uid":
	//	sample.PrintXID()
	//	sample.PrintULID()
	//default:
	//	flag.Usage()
	//}
}
