#include <stdio.h>
#include <stdlib.h>

// For by qsort
int compare(const void *i1, const void *i2) {
	return *(int*)i1 - *(int*)i2;
}

int main() {
	FILE *f = fopen("input.txt", "r");
	if (f == NULL) {
		printf("Unable to open input file");
		return 1;
	}

	// Get line count in file
	int lc = 0;
	while (!feof(f)) {
		if (fgetc(f) == '\n') {
			lc++;
		}
	}

	printf("line count: %d\n", lc);

	int arr_left[lc];
	int arr_right[lc];

	char line[20];
	int line_num = 0;

	rewind(f);
	// Read line by line
	while (fgets(line, sizeof(line), f)) {
		int i1, i2;
		sscanf(line, "%d   %d", &i1, &i2);	

		arr_left[line_num] = i1;
		arr_right[line_num] = i2;

		line_num++;
	}
	fclose(f);

	qsort(arr_left, lc, sizeof(int), compare);
	qsort(arr_right, lc, sizeof(int), compare);

	int total_dist = 0;
	for (int i = 0; i < lc; i++) {
		total_dist += abs(arr_left[i] - arr_right[i]);
	}
	
	printf("Part 1 - Solution : %d\n", total_dist);

	// Part 2 - naive approach (loops of death)
	int similarity_score = 0;
	// Left loop
	for (int i = 0; i < lc; i++) {
		// Right loop
		int count = 0;
		for (int j = 0; j < lc; j++) {
			if (arr_left[i] == arr_right[j]) {
				count++;
			}
		}	
		similarity_score += arr_left[i] * count;
	}	
	printf("Part 2 - Solution : %d\n", similarity_score);
	return 0;
}
