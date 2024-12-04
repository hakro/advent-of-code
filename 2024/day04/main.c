#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define H 150
#define L 150
#define XMAS_LEN 4

// Check the table at position (l, c).
// Use max to check you don't get out of array bounds
int check_right(char **table, size_t l, size_t c, size_t max);
int check_left(char **table, size_t l, size_t c);
int check_down(char **table, size_t l, size_t c, size_t max);
int check_up(char **table, size_t l, size_t c);
int check_diag_up_left(char **table, size_t l, size_t c, size_t max);
int check_diag_up_right(char **table, size_t l, size_t c, size_t max);
int check_diag_down_left(char **table, size_t l, size_t c, size_t max);
int check_diag_down_right(char **table, size_t l, size_t c, size_t max);
int get_all_xmases_for_pos(char **table, size_t l, size_t c, size_t max); //Returns the total xmas words

int main() {

	FILE *f = fopen("input.txt", "r");
	if (f == NULL) {
		printf("err: cant read input file");
		return 1;
	}
	int solution = 0;
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
			/* printf("--> %d\n", check_right(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_left(table, l, c)); */
			/* printf("--> %d\n", check_down(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_up(table, l, c)); */
			/* printf("--> %d\n", check_diag_up_left(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_diag_up_right(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_diag_down_left(table, l, c, nb_columns)); */
			/* printf("--> %d\n", check_diag_down_right(table, l, c, nb_columns)); */
			solution += get_all_xmases_for_pos(table, l, c, nb_columns);
		}
	}
	printf("Part1 solution : %d\n", solution);
	

	// -------- Clean up ------------ //
	for (int i = 0; i < nb_lines; i++) {
		free(table[i]);
	}
	free(table);
	// -------- Close --------- //
	fclose(f);
}


int check_right(char **table, size_t l, size_t c, size_t max) {
	if (c <= max - XMAS_LEN) {
		char dst[XMAS_LEN];
		strncpy(dst, table[l] + c, XMAS_LEN);
		if (strcmp(dst, "XMAS") == 0) {
			return 1;
		}
	} 
	return 0;
}

int check_left(char **table, size_t l, size_t c) {
	if (c >= XMAS_LEN - 1) {
		char dst[XMAS_LEN];
		strncpy(dst, table[l] + c - XMAS_LEN + 1, XMAS_LEN);
		if (strcmp(dst, "SAMX") == 0) {
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
			return 1;
		}
	}	
	return 0;
}

int check_diag_up_left(char **table, size_t l, size_t c, size_t max) {
	if (l >= XMAS_LEN - 1 && c >= XMAS_LEN - 1) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l-1][c-1];
		dst[2] = table[l-2][c-2];
		dst[3] = table[l-3][c-3];

		if (strcmp(dst, "XMAS") == 0) {
			return 1;
		}
	}
	return 0;
}

int check_diag_up_right(char **table, size_t l, size_t c, size_t max) {
	if (l >= XMAS_LEN - 1 && c <= max - XMAS_LEN) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l-1][c+1];
		dst[2] = table[l-2][c+2];
		dst[3] = table[l-3][c+3];

		if (strcmp(dst, "XMAS") == 0) {
			return 1;
		}
	}
	return 0;
}
int check_diag_down_left(char **table, size_t l, size_t c, size_t max) {
	if (l <= max - XMAS_LEN && c >= XMAS_LEN - 1) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l+1][c-1];
		dst[2] = table[l+2][c-2];
		dst[3] = table[l+3][c-3];

		if (strcmp(dst, "XMAS") == 0) {
			return 1;
		}
	}
	return 0;
}
int check_diag_down_right(char **table, size_t l, size_t c, size_t max) {
	if (l <= max - XMAS_LEN && c <= max - XMAS_LEN) {
		char dst[XMAS_LEN];
		dst[0] = table[l][c];
		dst[1] = table[l+1][c+1];
		dst[2] = table[l+2][c+2];
		dst[3] = table[l+3][c+3];

		if (strcmp(dst, "XMAS") == 0) {
			return 1;
		}
	}
	return 0;
}

int get_all_xmases_for_pos(char **table, size_t l, size_t c, size_t max) {

	int total = 0;

	total += check_right(table, l, c, max);
	total += check_left(table, l, c);
	total += check_down(table, l, c, max);
	total += check_up(table, l, c);
	total += check_diag_up_left(table, l, c, max);
	total += check_diag_up_right(table, l, c, max);
	total += check_diag_down_left(table, l, c, max);
	total += check_diag_down_right(table, l, c, max);

	return total;
}
