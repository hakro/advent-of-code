#include <iostream>
#include <fstream>
#include <vector>
#include <cassert>

#define print(x) std::cout << x << std::endl;

using Grid = std::vector<std::string>;
// Return item at specific location
char item_at(const Grid &grid, const int r, const int c);
// Count number of @'s adjacent to a given location
int count_adjacent_rolls(const Grid &grid, const int r, const int c);

// Part2
// Check if there are roll that can be moved in the grid
bool can_remove(const Grid &grid);
struct Roll {
	int r;
	int c;
};

int main() {
	// std::ifstream file("input-example.txt");
	std::ifstream file("input.txt");
	std::string line;
	Grid grid;

	while(std::getline(file, line)) {
		grid.push_back(line);
	}

	int result_1 = 0;
	for (int r = 0; r < grid.size(); r++) {
		for (int c = 0; c < grid[r].size(); c++) {
			if (item_at(grid, r, c) == '@' && count_adjacent_rolls(grid, r, c) < 4) {
				result_1++;
			}
		}
	}

	print("(part1)The answer my friend, is: " << result_1);


	// Part 2
	int result_2 = 0;
	while (can_remove(grid)) {
		std::vector<Roll> to_delete;
		for (int r = 0; r < grid.size(); r++) {
			for (int c = 0; c < grid[r].size(); c++) {
				if (item_at(grid, r, c) == '@' && count_adjacent_rolls(grid, r, c) < 4) {
					to_delete.push_back(Roll{r,c});
				}
			}
		}
		for (const Roll &roll : to_delete) {
			grid[roll.r][roll.c] = '.';
		}
		// print("count " << to_delete.size());
		result_2 += to_delete.size();
		if (to_delete.size() == 0) {
			break;
		}
	}
	print("(part2)The answer my friend, is: " << result_2);
}

char item_at(const Grid &grid, const int r, const int c) {
	assert(r < grid.size() && "fetched row cannot be bigger than grid rows");
	assert(c < grid[0].size() && "fetched column cannot be bigger than grid columns");
	return grid[r][c];
}

int count_adjacent_rolls(const Grid &grid, const int r, const int c) {
	bool first_row = r == 0;
	bool first_col = c == 0;
	bool last_row = r == grid.size() - 1;
	bool last_col = c == grid[0].size() - 1;

	int roll_count = 0;

	// North
	if (!first_row) {
		if (item_at(grid, r - 1, c) == '@') {
			roll_count++;
		}
		// West
		if (!first_col && item_at(grid, r - 1, c - 1) == '@') {
			roll_count++;
		}
		// Est
		if (!last_col && item_at(grid, r - 1, c + 1) == '@') {
			roll_count++;
		}
	}
	// Center
	if (!first_col && item_at(grid, r, c - 1) == '@') {
		roll_count++;
	}
	if (!last_col && item_at(grid, r, c + 1) == '@') {
		roll_count++;
	}
	// South
	if (!last_row) {
		if (item_at(grid, r + 1, c) == '@') {
			roll_count++;
		}
		// West
		if (!first_col && item_at(grid, r + 1, c - 1) == '@') {
			roll_count++;
		}
		// Est
		if (!last_col && item_at(grid, r + 1, c + 1) == '@') {
			roll_count++;
		}
	}

	return roll_count;
}

bool can_remove(const Grid &g) {
	bool res = false;
	for (int r = 0; r < g.size(); r++) {
		for (int c = 0; r < g[0].size(); c++) {
			if (count_adjacent_rolls(g, r, c) < 4) {
				return true;
			}
		}
	}
	return false;
}
