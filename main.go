package main

import (
	"os"
	"strings"
	"fmt"
	"strconv"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No executables in the args. Nothing to run!")
		return
	}

	environ := os.Environ()

	for _, pair := range environ {
		keyVal := strings.Split(pair, "=")
		if len(keyVal) == 2 {
			key := keyVal[0]
			val := keyVal[1]
			if strings.HasPrefix(keyVal[1], "[") && strings.HasSuffix(keyVal[1], "]") {
				vals := strings.Split(val[1:len(val)-1], ",")
				params[key] = vals
			} else if strings.HasPrefix(val, "(") && strings.HasSuffix(val, ")") {
				vals := strings.Split(val[1:len(val)-1], ",")
				if len(vals) == 3 {
					start := vals[0]
					stop := vals[1]
					step := vals[2]

					if len(start) > 6 {
						if strings.ToLower(start[:6]) == "start=" {
							start = start[6:]
						}
					}

					if len(stop) > 5 {
						if strings.ToLower(stop[:5]) == "stop=" {
							stop = stop[5:]
						}
					}

					if len(step) > 5 {
						if strings.ToLower(step[:5]) == "step=" {
							step = step[5:]
						}
					}

					vals := []string{}
					if startInt, stopInt, stepInt, ok := areInts(start, stop, step); ok {
						for ; startInt <= stopInt; startInt += stepInt {
							vals = append(vals, strconv.Itoa(startInt))
						}
					} else if startFloat, stopFloat, stepFloat, ok := areFloats(start, stop, step); ok {
						for ; startFloat <= stopFloat; startFloat += stepFloat {
							vals = append(vals, fmt.Sprintf("%f", startFloat))
						}
					}

					params[key] = vals
				}
			}

		}
	}

	for k := range params {
		keys = append(keys, k)
	}

	next_param(0, make(map[string]string))


}

var keys = []string{}
var combos = []string{}
var params = make(map[string][]string)

func serializeEnviron(combo map[string]string) []string {
	results := os.Environ()

	for k,v := range combo {
		for i, r := range results {
			if strings.HasPrefix(r, k+"=") {
				results[i] = k + "=" + v
			}
		}
	}
	return results
}

func next_param(param_index int, current_combo map[string]string) {
	if param_index == len(params) {
		// exec
		cmd := exec.Command(os.Args[1], os.Args[2:]...)

		cmd.Env = serializeEnviron(current_combo)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		cmd.Wait()

		//fmt.Println(, current_combo)
		return
	}
	for _, val := range params[keys[param_index]] {
		current_combo[keys[param_index]] = val

		current_combo_copy := make(map[string]string)
		for k,v := range current_combo {
			current_combo_copy[k] = v
		}

		next_param(param_index+1, current_combo_copy)
	}
}

func areInts(vals ... string) (start int, stop int, step int, succ bool) {
	var err error
	if start, err = strconv.Atoi(vals[0]); err != nil {
		return 0, 0, 0, false
	}
	if stop, err = strconv.Atoi(vals[1]); err != nil {
		return 0, 0, 0, false
	}
	if step, err = strconv.Atoi(vals[2]); err != nil {
		return 0, 0, 0, false
	}
	return start, stop, step,true
}

func areFloats(vals ... string) (start float64, stop float64, step float64, succ bool) {
	var err error
	if start, err = strconv.ParseFloat(vals[0], 64); err != nil {
		return 0, 0, 0, false
	}
	if stop, err = strconv.ParseFloat(vals[1], 64); err != nil {
		return 0, 0, 0, false
	}
	if step, err = strconv.ParseFloat(vals[2], 64); err != nil {
		return 0, 0, 0, false
	}
	return start, stop, step,true
}