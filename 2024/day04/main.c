#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define H 150
#define L 150
#define XMAS_LEN 4

// Check the table at position (l, c).
// Use max to check you don't get out of array bounds
int check_hor_right(char **table, size_t l, size_t c, size_t max);
int check_hor_left(char **table, size_t l, size_t c);
int check_down(char **table, size_t l, size_t c, size_t max);
int check_up(char **table, size_t l, size_t c);

int main() {

	FILE *f = fopen("input-example.txt", "r");
	if (f == NULL) {
		printf("err: cant read input file");
		return 1;
	}
	size_t nb_lines = 0;
	size_t nb_columns = 0;

	// Will hold all the input
	char **table = (char**)malloc(H * sizeof(char*));
	if (table == NULL) {
		printf("err: can't allocate memory for table");
		return 1;
	}

	char line[L];
	// For every line
	while (fgets(line, L, f)) {
		table[nb_lines] = (char*)malloc(L * sizeof(char));
		strcpy(table[nb_lines], line);
		nb_lines++;
	}

	nb_columns = strlen(table[0]) - 1; // -1 to ignore the \n
	printf("nb_lines: %ld / nb_col: %ld\n", nb_lines, nb_columns);

	for (int l = 0; l < nb_lines; l++) {
		for (int c = 0; c < nb_columns; c++) {
			if (table[l][c] != 'X') {
				continue;
			}
			/* printf("--> %d\n", check_hor_right(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_hor_left(table, l, c)); */
			/* printf("--> %d\n", check_down(table, l, c, nb_columns)); */
			printf("--> %d\n", check_up(table, l, c));
		}
	}



	// -------- Clean up ------------ //
	for (int i = 0; i < nb_lines; i++) {
		free(table[i]);
	}
	free(table);
	// -------- Close --------- //
	fclose(f);
}


int check_hor_right(char **table, size_t l, size_t c, size_t max) {
	if (c <= max - XMAS_LEN) {
		char dst[XMAS_LEN];
		strncpy(dst, table[l] + c, XMAS_LEN);
		if (strcmp(dst, "XMAS") == 0) {
			printf("-- found : %ld / %ld\n", l, c);
			return 1;
		}
	} 
	return 0;
}

int check_hor_left(char **table, size_t l, size_t c) {
	if (c >= XMAS_LEN - 1) {
		char dst[XMAS_LEN];
		strncpy(dst, table[l] + c - XMAS_LEN + 1, XMAS_LEN);
		if (strcmp(dst, "SAMX") == 0) {
			printf("-- found : %ld / %ld\n", l, c);
			return 1;
		}
	} 
	return 0;
}

int check_down(char **table, size_t l, size_t c, size_t max) {
	if (l <= max - XMAS_LEN) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l+1][c];
		dst[2] = table[l+2][c];
		dst[3] = table[l+3][c];

		if (strcmp(dst, "XMAS") == 0) {
			printf("-- found : %ld / %ld\n", l, c);
			return 1;
		}
	}	
	return 0;
}

int check_up(char **table, size_t l, size_t c) {
	if (l >= XMAS_LEN - 1) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l-1][c];
		dst[2] = table[l-2][c];
		dst[3] = table[l-3][c];

		if (strcmp(dst, "XMAS") == 0) {
			printf("-- found : %ld / %ld\n", l, c);
			return 1;
		}
	}	
	return 0;
}
