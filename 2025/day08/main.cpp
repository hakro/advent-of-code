#include <algorithm>
#include <cmath>
#include <fstream>
#include <iostream>
#include <set>
#include <sstream>
#include <vector>

#define print(x) std::cout << x << std::endl;

// For the example input
int MAX_PAIRS = 10;
// For the actual real input
// int MAX_PAIRS = 1000;

struct Junction {
	long x;
	long y;
	long z;

	long distance_squared(const Junction &other) {
		return std::pow(other.x - x, 2) + std::pow(other.y - y, 2) + std::pow(other.z - z, 2);
	}
};

struct Pair {
	int a; // Index of the junction in the junctions vector
	int b; // same as above
	long distance_squared;
	Pair(int a, int b, long d) : a(a), b(b), distance_squared(d) {};
	// for sorting
	bool operator<(const Pair &other) {
		return distance_squared < other.distance_squared;
	}
};

int main() {
	std::vector<Junction> junctions;
	std::vector<Pair> pairs;
	std::vector<std::set<int>> circuits; // each circuit is a vector of indices in the junctions vector
	std::ifstream file("input-example.txt");
	// std::ifstream file("input.txt");
	std::string line;

	while (std::getline(file, line)) {
		Junction junction;
		std::stringstream ss(line);
		char discard;
		ss >> junction.x;
		ss >> discard;
		ss >> junction.y;
		ss >> discard;
		ss >> junction.z;

		junctions.push_back(junction);
	}

	for (int i = 0; i < junctions.size(); i++) {
		for (int j = i + 1; j < junctions.size(); j++) {
			pairs.emplace_back(i, j, junctions[i].distance_squared(junctions[j]));
		}	
	}

	std::sort(pairs.begin(), pairs.end());
	int result = 1;
	int conn_counter = 0;
	for (const Pair &p : pairs) {
		if (conn_counter > MAX_PAIRS) {
			break;
		}
	// for (int i = 0; i < MAX_PAIRS; i++) {
		// const Pair &p = pairs[i];		
		bool added = false;
		// print(p.distance_squared << " " << p.a << " " << p.b);
		for (std::set<int> &c: circuits) {
			// if (std::find(c.begin(), c.end(), p.a) != c.end() || std::find(c.begin(), c.end(), p.b) != c.end()) {
			if (c.contains(p.a) && c.contains(p.b)) {
				break;
			}
			if (c.contains(p.a) || c.contains(p.b)) {
				// Add both, the std::set will make it so that there won't be any dups
				c.insert(p.a);
				c.insert(p.b);
				added = true;
				conn_counter++;
				break;
			}
		}
		if (added) {
			continue;
		}
		// The junctions are not in a circuit already, so add them as a new circuit (set)
		circuits.push_back({p.a, p.b});
		conn_counter++;
	}

	for (int i = 0; i < 3; i++) {
		// print(i << ": " << circuits[i].size());
		result *= circuits[i].size();		
	}
	print("The answer my friend, is: " << result);	
}
