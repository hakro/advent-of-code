#include <iostream>
#include <fstream>
#include <vector>

#define print(x) std::cout << x << std::endl;

int main() {
	// std::ifstream file("input-example.txt");
	std::ifstream file("input.txt");
	std::string line;
	std::vector<std::string> lines;
	while (std::getline(file, line)) {
		lines.push_back(line);	
	}	
	int result_1 = 0;
	for (int l = 0; l < lines.size() - 1; l++) {
		for (int c = 0; c < lines[l].size(); c++) {
			if (lines[l][c] == 'S' || lines[l][c] == '|') {
				if (lines[l + 1][c] == '^') {
					result_1++;
					lines[l + 1][c - 1] = '|';
					lines[l + 1][c + 1] = '|';
				} else {
					lines[l + 1][c] = '|';
				}
			}
		}	
	}

	print("The answer my friend, is : " << result_1);
}
