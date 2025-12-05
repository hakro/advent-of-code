#include <iostream>
#include <vector>
#include <fstream>

#define print(x) std::cout << x << std::endl

struct Range {
	long min;
	long max;
};

// Part2
// Merges all the ranges that overlap
void merge_ranges(std::vector<Range> &ranges);

int main() {
	// std::ifstream file("input-example.txt");
	std::ifstream file("input.txt");
	std::string line;

	std::vector<Range> ranges;
	// std::vector<long> ingredients;

	// Parse ranges
	while (std::getline(file, line)) {
		// print(line.size());
		if (line.size() == 0) [[unlikely]] {
			break;
		}
		int dash_pos = line.find('-');
		Range range {
			std::stol(line.substr(0, dash_pos)),
			std::stol(line.substr(dash_pos + 1))
		};
		ranges.push_back(range);
	}

	int result_1 = 0;
	// Then parse ingredients
	while (std::getline(file, line)) {
		long ing = std::stol(line);
		// ingredients.push_back(std::stol(line));
		for (const Range &range : ranges) {
			if (ing >= range.min && ing <= range.max) {
				result_1++;
				break;
			}
		}
	}

	print("(part1)The answer my friend, is : " << result_1);

	// Part 2
	long result_2 = 0;
	merge_ranges(ranges);
	for (const Range &range : ranges) {
		// print(range.min << " - " << range.max);
		result_2 += range.max - range.min + 1;
	}
	print("(part2)The answer my friend, is : " << result_2);
}

// Part2
void merge_ranges(std::vector<Range> &ranges) {
	// the vector is dirty if it changes, and we need to iterate again
	bool dirty = true;
	while (dirty) {
		dirty = false;
		for (int i = 0; i < ranges.size(); i++) {
			for (int j = 0; j < ranges.size(); j++) {
				if (i == j) continue;
				// range i :    ------
				// range j :  -----------
				// result  :  -----------
				if (ranges[i].min >= ranges[j].min && ranges[i].max <= ranges[j].max) {
					// enlarge range i
					ranges[i].min = ranges[j].min;
					ranges[i].max = ranges[j].max;
					// erase range j
					ranges.erase(ranges.begin() + j);
					dirty = true;
					break;
				}
				// range i :  -----------
				// range j :    -------
				// result  :  -----------
				if (ranges[i].min <= ranges[j].min && ranges[i].max >= ranges[j].max) {
					// erase range j
					ranges.erase(ranges.begin() + j);
					dirty = true;
					break;
				}
				// range i :  ------
				// range j :    -------
				// result  :  ---------
				if (ranges[i].min <= ranges[j].min && ranges[i].max >= ranges[j].min && ranges[i].max <= ranges[j].max) {
					ranges[i].max = ranges[j].max;
					// erase range j
					ranges.erase(ranges.begin() + j);
					dirty = true;
					break;
				}
				// range i :    -------
				// range j :  ------
				// result  :  ---------
				if (ranges[i].min >= ranges[j].min && ranges[i].min <= ranges[j].max && ranges[i].max >= ranges[j].max) {
					ranges[i].min = ranges[j].min;
					// erase range j
					ranges.erase(ranges.begin() + j);
					dirty = true;
					break;
				}
			}
			if (dirty) break;
		}
	}
}
