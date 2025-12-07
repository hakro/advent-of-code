#include <iostream>
#include <sstream>
#include <fstream>
#include <vector>

#define print(x) std::cout << x << std::endl;

int main() {
	// std::ifstream file("input-example.txt");
	std::ifstream file("input.txt");
	std::string line;

	std::vector<std::vector<int>> numbers;
	// Parse all numbers
	while(std::getline(file, line)) {
		if (line[0] == '*') {
			break;
		}
		std::stringstream ss(line);
		std::vector<int> line_numbers;

		int number;
		while (ss >> number) {
			line_numbers.push_back(number);
		}
		numbers.push_back(line_numbers);
	}
	// Parse operations *, +, - ...
	std::vector<char> ops;
	std::stringstream ss(line);
	char op;
	while (ss >> op) {
		ops.push_back(op);
	}

	long result_1 = 0;
	for (int c = 0; c < ops.size(); c++) {
		switch (ops[c]) {
			case '+': {
				long column_result = 0;
				for (int r = 0; r < numbers.size(); r++) {
					column_result += numbers[r][c];
				}
				result_1 += column_result;
			} break;
			case '*': {
				long column_result = 1;
				for (int r = 0; r < numbers.size(); r++) {
					column_result *= numbers[r][c];
				}
				result_1 += column_result;
			} break;
		}
	}
	print("The answer my friend, is: " << result_1);
}
