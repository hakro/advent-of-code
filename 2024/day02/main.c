#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main() {
	FILE *f = fopen("input.txt", "r");
	int part = 2; // Part 1 or 2 
	if (f == NULL) {
		printf("Unable to open input file\n");
		return 1;
	}

	// Read line by line (reports)
	int count_safe_lines = 0; // Part1 result
	char line[30];
	while (fgets(line, sizeof(line), f)) {
		int safe = 1; // Line is considered safe until proven otherwise
		int count_up = 0;
		int count_down = 0;
		int i = -1; //index of number being read in a line
		//read each int
		int num;
		int last_num;
		int line_dumpeners = 0;
		char *token;
		// Iterate on each number in a line
		token = strtok(line, " ");
		while (token != NULL) {
			// Nothing to do on the first one
			if (i == -1) {
				num = atoi(token);
				token = strtok(NULL, " ");
				i++;
				continue;
			}

			last_num = num;
			num = atoi(token);

			// Check two adjacent levels differ by at least one and at most three
			if (!(abs(num - last_num) >= 1 && abs(num - last_num) <= 3 )) {
				if (line_dumpeners == 0) {
					line_dumpeners++;
				} else {
					safe = 0;
					break;
				}
			}

			if (num > last_num) {
				count_up++;
			} else {
				count_down++;
			}

			// Next Token
			token = strtok(NULL, " ");
			i++;
		}
		// Check levels are either all increasing or all decreasing
		if (count_up > 0 && count_down > 0) {
			safe = 0;
		}
		if (safe == 1) {
			count_safe_lines++;
		}
		/* printf("up: %d -- down: %d\n", count_up, count_down); */
	}

	printf("Part1 result - Safe reports count : %d\n", count_safe_lines);
	
	fclose(f);
	return 0;
}
