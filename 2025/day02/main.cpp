#include <cassert>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

bool is_valid(const std::string &id);
bool is_valid(long id);

// For part2
bool is_valid2(const std::string &id);
bool is_valid2(long id);
std::vector<std::string> split(const std::string &id, int by);
bool is_repeated(const std::vector<std::string> &v);

int main() {
	std::ifstream file("input.txt");
	std::string line;
	std::vector<std::string> ranges;

	while (std::getline(file, line)) {
		int next_comma_pos = line.find(',');
		while (next_comma_pos != std::string::npos) {
			ranges.push_back(line.substr(0, next_comma_pos));
			line = line.substr(next_comma_pos + 1, line.size());
			next_comma_pos = line.find(',');
		}
		// Add the last one
		ranges.push_back(line);
	}

	long result1 = 0;
	long result2 = 0;
	for(auto v : ranges) {
		// std::cout << v << std::endl;
		long low = std::stol(v.substr(0, v.find('-')));
		long high = std::stol(v.substr(v.find('-') + 1));
		for (long i = low; i <= high; i++) {
			if (!is_valid(i)) {
				result1 += i;
			}
			if (!is_valid2(i)) {
				result2 += i;
			}
		}
	}

	std::cout << "(part1)The answer my friend, is : " << result1 << std::endl;
	std::cout << "(part2)The answer my friend, is : " << result2 << std::endl;
	return 0;
}

// All IDs with repeating sequences are invalid
// Examples: 55, 6464 ... are invalid
bool is_valid(const std::string &id) {
	// odd-sized id's are valid
	if (id.size() % 2) {
		return true;
	}

	if (!id.substr(0, id.size()/2).compare(id.substr(id.size()/2, id.size()))) {
		return false;
	}
	return true;
}
bool is_valid(long id) {
	return is_valid(std::to_string(id));	
}

// For part2
bool is_valid2(const std::string &id) {
	int offset = 1;
	while (offset <= id.size() / 2) {
		for (int i = 0; i < id.size(); i += offset) {
			auto splits = split(id, offset);
			if (is_repeated(splits)) {
				return false;
			}
		}
		offset++;
	}
	return true;
}
bool is_valid2(long id) {
	return is_valid2(std::to_string(id));
}

// Split a string in equal parts
// 222 by 1 becomes {'2', '2', '2'}
// 4444 by 2 becomes {'44', '44'}
std::vector<std::string> split(const std::string &id, int by) {
	assert(by <= id.size() / 2 && "Split should stop at half the size of the input string");
	std::vector<std::string> res;
	for (int i = 0; i < id.size(); i += by) {
		res.push_back(id.substr(i, by));
	}
	return res;
}

// Check if all element in the vector are the same
bool is_repeated(const std::vector<std::string> &vec) {
	assert(vec.size() > 0 && "This vector should not be empty");
	for (auto val : vec) {
		if (val.compare(vec[0]) != 0) {
			return false;
		}
	}
	return true;
}
