#include <iostream>
#include <fstream>

int main() {
	bool part2 = true;

	std::ifstream file("input.txt");
	std::string line;

	int current_dial = 50;
	int solution = 0;

	while(std::getline(file, line)) {
		int next_dial_rot = std::stoi(line.substr(1));
		int old_dial = current_dial;

		if (part2 && next_dial_rot > 99) {
			solution += (next_dial_rot / 100);
			next_dial_rot = next_dial_rot - (next_dial_rot / 100) * 100;
	 	}

		switch (line[0]) {
			case 'L': 
				current_dial = (100 + (current_dial - next_dial_rot)) % 100;
				if (part2) {
					if ((old_dial <= current_dial) && current_dial != 0 && old_dial != 0) {
						solution++;
					}  
				}
				break;  
			case 'R': 
				current_dial = (current_dial + next_dial_rot) % 100;
				if (part2) {
					if ((old_dial >= current_dial) && current_dial != 0 && old_dial != 0) {
						solution++;
					}  
				}
	 			break;  
		}
		// std::cout << current_dial << std::endl;
		if (current_dial == 0) {
			solution++;
		}
	}

	std::cout << "----------------------------------------" << std::endl;
	std::cout << "The answer my friend, is : " << solution << std::endl;
}
