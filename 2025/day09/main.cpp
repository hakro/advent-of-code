#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>
#include <sstream>
#include <cmath>

#define print(x) std::cout << x << std::endl

struct Tile {
	enum class Color {
		RED,
		GREEN
	};

	long x;
	long y;
	Color c; // for part 2

	Tile(long x, long y, Color c) : x(x), y(y), c(c) {}
	long area(const Tile &other) {
		return (labs(other.x - x) + 1) * (labs(other.y - y) + 1);
	}
};

int main() {
	std::ifstream file("input-example.txt");
	// std::ifstream file("input.txt");
	std::string line;
	std::vector<Tile> tiles;

	while (std::getline(file, line)) {
		std::stringstream ss(line);	
		long x, y;
		char discard;
		ss >> x >> discard >> y;
		tiles.emplace_back(x, y, Tile::Color::RED);
	}
	// for (auto t: tiles) {
	// 	print(t.x << ", " << t.y);
	// }
	//
	long max_area = 0;
	for (int i = 0; i < tiles.size(); i++) {
		for (int j = i + 1; j < tiles.size(); j++) {
			long area = tiles[i].area(tiles[j]);
			// print(area << " area: " << " for (" << tiles[i].x << ", " << tiles[i].y << ") and (" << tiles[j].x << ", " << tiles[j].y << ")");
			if (area > max_area) {
				max_area = area;
			}
		}
	}
	print("(part1)The answer my friend, is: " << max_area);

	// Part 2
	// Fill up the green line between the consecutive red ones
	for (int i = 0; i < tiles.size() - 1; i++) {
		if (tiles[i].x == tiles[i + 1].x) {
			long min_y = std::min(tiles[i].y, tiles[i + 1].y);
			long max_y = std::max(tiles[i].y, tiles[i + 1].y);	
			for (int y = min_y; y < max_y; y++)	{
				tiles.emplace_back(tiles[i].x, y, Tile::Color::GREEN);
			}
		}
		if (tiles[i].y == tiles[i + 1].y) {
			long min_x = std::min(tiles[i].x, tiles[i + 1].x);
			long max_x = std::max(tiles[i].x, tiles[i + 1].x);
			for (int x = min_x; x < max_x; x++)	{
				tiles.emplace_back(x, tiles[i].y, Tile::Color::GREEN);
			}
		}
	}
	// for (int i = 0; i < tiles.size(); i++) {
	// 	print("(" << tiles[i].x << ", " << tiles[i].y << ")");
	// }

	// // Part 2 example visualization
	// std::vector<char> chars{'.'};
	// for (int i = 0; i < 9 * 14; i++) {
	// 	chars.push_back('.');
	// 	if (i % 14 == 0) {
	// 		std::cout << std::endl;
	// 	}
	// 	// std::cout << chars[i];
	// }
	// std::cout << std::endl;
	

}
