#include <fstream>
#include <iostream>

#define print(x) std::cout << x << std::endl

long max(const std::string &s);
char max_before(const std::string &s, int b);

int main() {
	std::ifstream file("input-example.txt");
	std::string bank;
	long result = 0;
	while (std::getline(file, bank)) {
		result += max(bank);
	}
	print("The answer my friend, is: " << result);
}

long max(const std::string &s) {
	char first_max = '0';
	int first_max_index = 0;
	char second_max = '0';
	// Find first max digit in the string
	// The last digit in the string needs to be ignored
	for (int i = 0; i < s.size() - 1; i++) {
		if (s[i] > first_max) {
			first_max = s[i];
			first_max_index = i;
		}
	}
	// print(s.substr(first_max_index + 1));
	// Find the second highest max in the remaining string
	for (char c : s.substr(first_max_index + 1)) {
		if (c > second_max) {
			second_max = c;
		}
	}
	return std::stol(std::string{first_max, second_max});
}
