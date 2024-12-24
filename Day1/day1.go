package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func extract_number(stringLine string) (int, int, error) {
	var values []string = strings.Split(stringLine, "   ")
	leftValue, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println("Error converting string to int: ", err, "")
		return 0, 0, err
	}
	rightValue, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println("Error converting string to int: ", err, "")
		return 0, 0, err
	} else {
		return leftValue, rightValue, nil
	}

}

func read_list(reader *bufio.Reader) ([]int, []int) {
	var left_list []int
	var right_list []int

	for {
		var line, _, exception = reader.ReadLine()
		if len(line) > 0 {
			if exception != nil {
				fmt.Println("Error reading line: ", exception, "")
				break
			}
			var stringLine string = string(line)

			left_number, right_number, err := extract_number(stringLine)
			if err != nil {
				break
			}

			left_list = append(left_list, left_number)
			right_list = append(right_list, right_number)
		}
		if exception != nil {
			fmt.Println("Reached EOF")
			break
		}
	}

	return left_list, right_list
}

func get_difference(left_list []int, right_list []int) ([]int, error) {
	var difference_list []int

	if len(left_list) != len(right_list) {
		fmt.Println("Error: lists are not the same length")
		return difference_list, nil
	}

	for i := 0; i < len(left_list); i++ {
		var difference int = 0
		if left_list[i] > right_list[i] {
			difference = left_list[i] - right_list[i]
		} else {
			difference = right_list[i] - left_list[i]
		}

		difference_list = append(difference_list, difference)
	}

	return difference_list, nil
}

func calc_total_distance(difference_list []int) int {
	var total_distance int = 0

	for i := 0; i < len(difference_list); i++ {
		total_distance += difference_list[i]
	}
	return total_distance
}

func find_value_in_slice(slice []int, value int) int {
	var count int = 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			count++
		}
	}
	return count
}

func calc_similarity_score(similarity_score_list []int) int {
	var similarity_score int
	for i := 0; i < len(similarity_score_list); i++ {
		similarity_score += similarity_score_list[i]
	}
	return similarity_score
}

func find_similarity_score(left_list []int, right_list []int) (int, error) {
	var similarity_score_list []int

	if len(left_list) != len(right_list) {
		fmt.Println("Error: cannot calculate similarity, lists are not the same length")
		return 0, nil
	}

	for i := 0; i < len(left_list); i++ {
		var count int = find_value_in_slice(right_list, left_list[i])
		var similarity_score int = left_list[i] * count
		similarity_score_list = append(similarity_score_list, similarity_score)
	}

	return calc_similarity_score(similarity_score_list), nil
}

func main() {
	fmt.Println("Day 1\n Reading input file...")
	file, exception := os.Open("input.txt")
	if exception != nil {
		fmt.Println("Error reading file: ", exception, "")
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	left_list, right_list := read_list(reader)

	slices.Sort(left_list)
	slices.Sort(right_list)

	difference_list, err := get_difference(left_list, right_list)
	if err != nil {
		fmt.Println("Error getting difference: ", err, "")
		os.Exit(1)
	}
	total_distance := calc_total_distance(difference_list)
	fmt.Println("Total distance: ", total_distance, "")
	fmt.Println("Total distance calculated successfully\nStart finding similarity score")

	similarity_score, err := find_similarity_score(left_list, right_list)
	if err != nil {
		fmt.Println("Error getting similarity score: ", err, "")
		os.Exit(1)
	}

	fmt.Println("Similarity score: ", similarity_score, "")
}
